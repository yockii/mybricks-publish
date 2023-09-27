package controller

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	logger "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"manatee-publish/internal/domain"
	"manatee-publish/internal/model"
	"manatee-publish/internal/service"
	"manatee-publish/pkg/common"
	"manatee-publish/pkg/config"
	"manatee-publish/pkg/database"
	"manatee-publish/pkg/server"
	"strconv"
	"strings"
)

type publishController struct {
	common.BaseController[*model.PublishInfo, *domain.PublishInfoDomain]
}

func (c *publishController) GetService() common.Service[*model.PublishInfo] {
	return service.PublishService
}

func (c *publishController) InitManage() {
	r := server.Group("/openapi/v1/publish")

	r.Post("/manatee", c.ManateePublish)

	server.Get("/asset/+", c.CheckOriginAndCache, c.GetAsset)
}

func (c *publishController) ManateePublish(ctx *fiber.Ctx) error {
	body := string(ctx.Body())
	in := gjson.Parse(body)

	page := new(model.Page)
	page.ProductID = in.Get("productId").Int()
	page.Name = in.Get("productName").String()
	page.Type = in.Get("type").String()
	//page.PageVersionID

	pageVersion := new(model.PageVersion)
	//pageVersion.PageID ✔
	//pageVersion.AssetVersionID ✔
	pageVersion.Env = in.Get("env").String()
	pageVersion.PublisherEmail = in.Get("publisherEmail").String()
	pageVersion.PublisherName = in.Get("publisherName").String()
	pageVersion.Version = in.Get("version").String()
	pageVersion.CommitInfo = in.Get("commitInfo").String()
	pageVersion.SchemaJson = in.Get("content.json").String()

	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		d, e := service.PageService.Add(page, tx)
		if e != nil {
			return e
		}
		if d {
			// 已有，则获取
			page, e = service.PageService.Instance(&model.Page{ProductID: page.ProductID})
			if e != nil {
				return e
			}
		}
		// page存储完成
		pageVersion.PageID = page.ID

		// 存储文件内容
		{
			// 存储html
			htmlUrl := fmt.Sprintf("/asset/%d/index.html", page.ID)
			pageVersion.AssetVersionID, e = service.AssetService.AddAsset(&model.Asset{
				Path: htmlUrl,
			}, strings.NewReader(in.Get("content.html").String()))
			if e != nil {
				return e
			}

			// 存储对应的js
			jsList := in.Get("content.js").Array()
			for _, js := range jsList {
				jsUrl := fmt.Sprintf("/asset/%d/%s", page.ID, js.Get("name").String())
				_, e = service.AssetService.AddAsset(&model.Asset{
					Path: jsUrl,
				}, strings.NewReader(js.Get("content").String()))
				if e != nil {
					return e
				}
			}
			// 存储全局通用文件
			globalDeps := in.Get("content.globalDeps").Array()
			for _, f := range globalDeps {
				filePath := f.Get("path").String()
				// 先查找是否存在
				if service.AssetService.Count(filePath) > 0 {
					continue
				}

				// 不存在则添加
				_, e = service.AssetService.AddAsset(&model.Asset{
					Path: filePath,
				}, strings.NewReader(f.Get("content").String()))
				if e != nil {
					return e
				}
			}
		}

		// 存储page版本
		d, e = service.PageVersionService.Add(pageVersion, tx)
		if e != nil {
			return e
		}
		if d {
			return errors.New("页面版本数据重复")
		}

		page.PageVersionID = pageVersion.ID
		// 更新page
		if _, e = service.PageService.Update(page, tx); e != nil {
			return e
		}

		return nil
	}); err != nil {
		return ctx.JSON(&domain.ManateePublishResponse{
			Code:    0,
			Message: "发布失败: " + err.Error(),
		})
	}

	publishInfo := new(domain.PublishInfoDomain)
	if err := ctx.BodyParser(publishInfo); err != nil {
		logger.Errorln(err)
		return ctx.JSON(&domain.ManateePublishResponse{
			Code:    0,
			Data:    domain.PublishData{},
			Message: "发布集成参数解析失败, " + err.Error(),
		})
	}
	publishContent := publishInfo.Content
	saveInfo := publishInfo.PublishInfo
	saveInfo.SchemaJson = publishContent.Json

	//// 构建一个html的url
	//urlPrefix := "/asset/"
	//
	//if err := database.DB.Transaction(func(tx *gorm.DB) error {
	//	d, e := c.GetService().Add(&saveInfo, tx)
	//	if e != nil {
	//		return e
	//	}
	//	if d {
	//		return errors.New("数据重复")
	//	}
	//
	//	urlPrefix = fmt.Sprintf("/asset/%d/", saveInfo.ID)
	//
	//	// 保存html
	//	htmlUrl := urlPrefix + "index.html"
	//	d, e = service.AssetService.Add(&model.Asset{
	//		Path: htmlUrl,
	//	}, strings.NewReader(publishContent.Html))
	//	if e != nil {
	//		return e
	//	}
	//	if d {
	//		return errors.New("数据重复")
	//	}
	//	// 保存html对应的js内容
	//	for _, js := range publishContent.Js {
	//		jsUrl := urlPrefix + js.Name
	//		d, e = service.AssetService.Add(&model.Asset{
	//			Path: jsUrl,
	//		}, strings.NewReader(js.Content))
	//		if e != nil {
	//			return e
	//		}
	//		if d {
	//			return errors.New("数据重复")
	//		}
	//	}
	//	// 处理全局通用文件
	//	for _, f := range publishContent.GlobalDeps {
	//		filePath := f.Path
	//		// 先查找是否存在
	//		if service.AssetService.Count(filePath) > 0 {
	//			continue
	//		}
	//
	//		// 不存在则添加
	//		d, e = service.AssetService.Add(&model.Asset{
	//			Path: filePath,
	//		}, strings.NewReader(f.Content))
	//		if e != nil {
	//			return e
	//		}
	//		if d {
	//			return errors.New("数据重复")
	//		}
	//	}
	//
	//	return nil
	//}); err != nil {
	//	return ctx.JSON(&domain.ManateePublishResponse{
	//		Code:    0,
	//		Message: "发布失败: " + err.Error(),
	//	})
	//}

	return ctx.JSON(&domain.ManateePublishResponse{
		Code: 1,
		Data: domain.PublishData{
			//config.GetString("server.prefix") + "/asset/" + page.ID + "index.html",
			Url: fmt.Sprintf("%s/asset/%d/index.html", config.GetString("server.prefix"), page.ID),
		},
		Message: "发布成功",
	})
}

func (c *publishController) GetAsset(ctx *fiber.Ctx) error {
	path := ctx.Path()
	// 查找是否存在该资产记录
	if service.AssetService.Count(path) == 0 {
		// 如果有 public/ 但是 /asset/ 与 public/ 之间还有其他字符串，则重定向到 /asset/public/xxxxx
		if strings.Contains(path, "public/") {
			if strings.Index(path, "/asset/")+7 != strings.Index(path, "public/") {
				return ctx.Redirect("/asset/" + path[strings.Index(path, "public/"):])
			}
			// 如果已经重定向过，那么去找到该文件
			path = path[strings.Index(path, "public/"):]
			if service.AssetService.Count(path) == 0 {
				return ctx.SendStatus(fiber.StatusNotFound)
			}
		}
	}
	// 从oss中读取文件内容并返回
	content, err := service.AssetService.Download(path)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	// 设置content-type?
	if strings.HasSuffix(path, ".html") {
		ctx.Response().Header.SetContentType("text/html")
	} else if strings.HasSuffix(path, ".js") {
		ctx.Response().Header.SetContentType("application/javascript")
	} else if strings.HasSuffix(path, ".css") {
		ctx.Response().Header.SetContentType("text/css")
	}

	return ctx.SendStream(content)
}

func (c *publishController) CheckOriginAndCache(ctx *fiber.Ctx) error {
	origin := ctx.GetReqHeaders()[fiber.HeaderOrigin]
	path := ctx.Path()
	// path如果是index.html，则检查origin对应的应用是否允许访问path
	if strings.HasSuffix(path, "index.html") {
		if origin != "" {
			// 从path中解析出pageId
			pageIdStr := path[strings.Index(path, "/asset/")+7 : strings.Index(path, "/index.html")]
			pageId, err := strconv.ParseUint(pageIdStr, 10, 64)
			if err != nil {
				logger.Errorln(err)
				return ctx.SendStatus(fiber.StatusNotFound)
			}
			// 检查origin对应的应用是否允许访问path
			if !service.ApplicationService.CheckOrigin(pageId, origin) {
				// 禁止跨域访问
				return ctx.SendStatus(fiber.StatusForbidden)
			}

		}
	}
	ctx.Response().Header.Set(fiber.HeaderAccessControlAllowOrigin, origin)
	ctx.Response().Header.Set(fiber.HeaderAccessControlAllowCredentials, "true")
	err := ctx.Next()
	if err == nil {
		// 设置缓存
		ctx.Response().Header.Set(fiber.HeaderCacheControl, "max-age=86400")
		//TODO 服务器自己也缓存一下，放到redis中

	}
	return err
}

func init() {
	c := new(publishController)
	c.BaseController = common.BaseController[*model.PublishInfo, *domain.PublishInfoDomain]{
		Controller: c,
	}

	Controllers = append(Controllers, c)

}
