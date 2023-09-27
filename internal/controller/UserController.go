package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gomodule/redigo/redis"
	logger "github.com/sirupsen/logrus"
	"manatee-publish/internal/constant"
	"manatee-publish/internal/domain"
	"manatee-publish/internal/helper"
	"manatee-publish/internal/middleware"
	"manatee-publish/internal/model"
	"manatee-publish/internal/service"
	"manatee-publish/pkg/cache"
	"manatee-publish/pkg/common"
	"manatee-publish/pkg/config"
	"manatee-publish/pkg/crypto"
	"manatee-publish/pkg/server"
	"manatee-publish/pkg/util"
	"strconv"
)

type userController struct {
	common.BaseController[*model.User, *domain.UserDomain]
}

func (c *userController) InitManage() {
	r := server.Group("/api/v1/user")
	// 登录
	r.Post("/login", c.Login)

	r.Post("/add", middleware.NeedAuthorization(constant.ResourceUserAdd), c.Add)
	r.Put("/update", middleware.NeedAuthorization(constant.ResourceUserUpdate), c.Update)
	r.Delete("/delete", middleware.NeedAuthorization(constant.ResourceUserDelete), c.Delete)
	r.Get("/detail", middleware.NeedAuthorization(constant.ResourceUserDetail), c.Detail)
	r.Get("/list", middleware.NeedAuthorization(constant.ResourceUserList), c.List)

	r.Put("/updatePassword", middleware.NeedAuthorization(constant.NeedLogin), c.UpdatePassword)
	r.Put("/resetPassword", middleware.NeedAuthorization(constant.ResourceUserResetPassword), c.ResetPassword)
	r.Get("/roles", middleware.NeedAuthorization(constant.ResourceUserList), c.UserRoles)
	r.Post("/assignRoles", middleware.NeedAuthorization(constant.ResourceUserAssignRole), c.AssignRoles)
}

func (*userController) GetService() common.Service[*model.User] {
	return service.UserService
}

func (*userController) NewModel() *model.User {
	return new(model.User)
}

func (*userController) NewDomain() *domain.UserDomain {
	return new(domain.UserDomain)
}

func (c *userController) Add(ctx *fiber.Ctx) error {
	instance := new(model.User)
	if err := ctx.BodyParser(instance); err != nil {
		logger.Errorln(err)
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamParseError,
			Msg:  server.ResponseMsgParamParseError,
		})
	}

	// 处理必填
	if instance.Username == "" {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamNotEnough,
			Msg:  server.ResponseMsgParamNotEnough + " username",
		})
	}

	// 解析密码
	if pwd, err := crypto.Sm2Decrypt(instance.Password); err != nil {
		logger.Errorln(err)
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamParseError,
			Msg:  server.ResponseMsgParamParseError + "密码不准确",
		})
	} else {
		instance.Password = pwd
	}

	if instance.Password != "" {
		isStrong := util.PasswordStrengthCheck(8, 50, 4, instance.Password)
		if !isStrong {
			return ctx.JSON(&server.CommonResponse{
				Code: server.ResponseCodePasswordStrengthInvalid,
				Msg:  server.ResponseMsgPasswordStrengthInvalid,
			})
		}
	}

	duplicated, err := service.UserService.Add(instance)
	if err != nil {
		logger.Errorln(err)
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeDatabase,
			Msg:  server.ResponseMsgDatabase + err.Error(),
		})
	}
	if duplicated {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeDuplicated,
			Msg:  server.ResponseMsgDuplicated,
		})
	}
	return ctx.JSON(&server.CommonResponse{
		Data: instance,
	})
}

func (c *userController) UpdatePassword(ctx *fiber.Ctx) (err error) {
	instance := new(model.User)
	if err = ctx.BodyParser(instance); err != nil {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamParseError,
			Msg:  server.ResponseMsgParamParseError,
		})
	}
	if instance.Password == "" {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamNotEnough,
			Msg:  server.ResponseMsgParamNotEnough + ": password",
		})
	}

	// 解析密码
	var pwd string
	if pwd, err = crypto.Sm2Decrypt(instance.Password); err != nil {
		logger.Errorln(err)
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamParseError,
			Msg:  server.ResponseMsgParamParseError + "密码不准确",
		})
	}
	instance.Password = pwd

	instance.ID, err = helper.GetCurrentUserID(ctx)
	var success bool
	success, err = service.UserService.UpdatePassword(instance)
	if err != nil {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeDatabase,
			Msg:  server.ResponseMsgDatabase + err.Error(),
		})
	}
	return ctx.JSON(&server.CommonResponse{
		Data: success,
	})
}

func (c *userController) ResetPassword(ctx *fiber.Ctx) error {
	instance := new(model.User)
	if err := ctx.BodyParser(instance); err != nil {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamParseError,
			Msg:  server.ResponseMsgParamParseError,
		})
	}
	if instance.ID == 0 || instance.Password == "" {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamNotEnough,
			Msg:  server.ResponseMsgParamNotEnough + " id/password",
		})
	}

	// 解析密码
	if pwd, err := crypto.Sm2Decrypt(instance.Password); err != nil {
		logger.Errorln(err)
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamParseError,
			Msg:  server.ResponseMsgParamParseError + "密码不准确",
		})
	} else {
		instance.Password = pwd
	}

	success, err := service.UserService.UpdatePassword(instance)
	if err != nil {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeDatabase,
			Msg:  server.ResponseMsgDatabase + err.Error(),
		})
	}
	return ctx.JSON(&server.CommonResponse{
		Data: success,
	})
}

func (c *userController) UserRoles(ctx *fiber.Ctx) error {
	userIdStr := ctx.Query("id")
	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		logger.Errorln(err)
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamParseError,
			Msg:  server.ResponseMsgParamParseError + " id",
		})
	}
	if userId == 0 {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamNotEnough,
			Msg:  server.ResponseMsgParamNotEnough + " id",
		})
	}
	roleList, err := service.UserService.Roles(userId)
	if err != nil {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeDatabase,
			Msg:  server.ResponseMsgDatabase + err.Error(),
		})
	}
	var roleIdList []string
	for _, role := range roleList {
		roleIdList = append(roleIdList, strconv.FormatUint(role.ID, 10))
	}
	return ctx.JSON(&server.CommonResponse{
		Data: roleIdList,
	})
}

func (c *userController) AssignRoles(ctx *fiber.Ctx) error {
	instance := new(domain.UserDispatchRolesRequest)
	if err := ctx.BodyParser(instance); err != nil {
		logger.Errorln(err)
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamParseError,
			Msg:  server.ResponseMsgParamParseError,
		})
	}
	// 处理必填
	if instance.UserID == 0 {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamNotEnough,
			Msg:  server.ResponseMsgParamNotEnough + " user id",
		})
	}

	success, err := service.UserService.DispatchRoles(instance.UserID, instance.RoleIDList)
	if err != nil {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeDatabase,
			Msg:  server.ResponseMsgDatabase + err.Error(),
		})
	}

	if success {
		// 成功处理后，清除用户的权限缓存
		conn := cache.Get()
		defer func(conn redis.Conn) {
			_ = conn.Close()
		}(conn)
		key := fmt.Sprintf("%s:%d", constant.RedisKeyUserRoles, instance.UserID)
		_, err = conn.Do("DEL", key)
		if err != nil {
			logger.Errorln(err)
		}
		// 删除即可，等待中间件重新加载
	}

	return ctx.JSON(&server.CommonResponse{
		Data: success,
	})
}

func (c *userController) Login(ctx *fiber.Ctx) error {
	instance := new(model.User)
	if err := ctx.BodyParser(instance); err != nil {
		logger.Errorln(err)
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamParseError,
			Msg:  server.ResponseMsgParamParseError,
		})
	}

	// 处理必填
	if instance.Username == "" || instance.Password == "" {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamNotEnough,
			Msg:  server.ResponseMsgParamNotEnough + ": 用户名及密码",
		})
	}

	// 解析密码
	if pwd, err := crypto.Sm2Decrypt(instance.Password); err != nil {
		logger.Errorln(err)
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamParseError,
			Msg:  server.ResponseMsgParamParseError + "密码不准确",
		})
	} else {
		instance.Password = pwd
	}

	isStrong := util.PasswordStrengthCheck(8, 50, 4, instance.Password)
	if !isStrong {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodePasswordStrengthInvalid,
			Msg:  server.ResponseMsgPasswordStrengthInvalid,
		})
	}

	user, notMatch, err := service.UserService.LoginWithUsernameAndPassword(instance.Username, instance.Password)
	if err != nil {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeDatabase,
			Msg:  server.ResponseMsgDatabase + err.Error(),
		})
	}
	if notMatch {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeDataNotMatch,
			Msg:  "用户名与密码" + server.ResponseMsgDataNotMatch,
		})
	}
	return c.generateLoginResponse(user, ctx)
}

func (c *userController) generateLoginResponse(user *model.User, ctx *fiber.Ctx) error {
	jwtToken, err := generateJwtToken(strconv.FormatUint(user.ID, 10), "")
	if err != nil {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeGeneration,
			Msg:  server.ResponseMsgGeneration + err.Error(),
		})
	}
	user.Password = ""
	userDomain := &domain.UserDomain{
		User: *user,
	}

	return ctx.JSON(&server.CommonResponse{
		Data: map[string]interface{}{
			"token": jwtToken,
			"user":  userDomain,
		},
	})
}

func (c *userController) UpdateMyInfo(ctx *fiber.Ctx) error {
	instance := new(domain.UserDomain)
	if err := ctx.BodyParser(instance); err != nil {
		logger.Errorln(err)
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeParamParseError,
			Msg:  server.ResponseMsgParamParseError,
		})
	}
	userId, err := helper.GetCurrentUserID(ctx)
	if err != nil {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeUnknownError,
			Msg:  "未知错误",
		})
	}
	instance.ID = userId
	success, err := service.UserService.UpdateDomain(&domain.UserDomain{
		User: model.User{
			BaseModel: common.BaseModel{ID: instance.ID},
			NickName:  instance.NickName,
		},
	})
	if err != nil {
		return ctx.JSON(&server.CommonResponse{
			Code: server.ResponseCodeDatabase,
			Msg:  server.ResponseMsgDatabase + err.Error(),
		})
	}
	return ctx.JSON(&server.CommonResponse{
		Data: success,
	})
}

func generateJwtToken(userId, tenantId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	sid := util.GenerateXid()

	conn := cache.Get()
	defer func(conn redis.Conn) {
		_ = conn.Close()
	}(conn)
	sessionKey := fmt.Sprintf("%s:%s", constant.RedisSessionIdKey, sid)

	_, err := conn.Do("SETEX", sessionKey, config.GetInt("userTokenExpire"), userId)
	if err != nil {
		logger.Errorln(err)
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	claims[constant.JwtClaimUserId] = userId
	claims[constant.JwtClaimTenantId] = tenantId
	claims[constant.JwtClaimSessionId] = sid

	t, err := token.SignedString([]byte(constant.JwtSecret))
	if err != nil {
		logger.Errorln(err)
		return "", err
	}
	return t, nil
}

func init() {
	c := new(userController)
	c.BaseController = common.BaseController[*model.User, *domain.UserDomain]{
		Controller: c,
	}

	Controllers = append(Controllers, c)
}
