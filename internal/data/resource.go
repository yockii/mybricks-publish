package data

import (
	logger "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"manatee-publish/internal/constant"
	"manatee-publish/internal/model"
	"manatee-publish/pkg/common"
	"manatee-publish/pkg/config"
	"manatee-publish/pkg/database"
	"manatee-publish/pkg/util"
)

func InitData() {
	// 自动建表
	//_ = database.AutoMigrate(constant.Models...) // 不采用直接自动建表的方式，而是主动循环，才能创建表注释
	if config.GetString("database.driver") == "mysql" {
		migrator := database.DB.Migrator()
		for _, m := range common.Models {
			if !migrator.HasTable(m) {
				if err := database.DB.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='"+m.TableComment()+"';").AutoMigrate(m); err != nil {
					logger.Errorln(err)
				}
			} else {
				_ = migrator.AutoMigrate(m)
			}
		}
	} else {
		migrator := database.DB.Migrator()
		var mList []interface{}
		for _, m := range common.Models {
			mList = append(mList, m)
		}
		_ = migrator.AutoMigrate(mList...)
	}

	// 初始化一些数据
	// 初始化一个admin用户
	adminUser := &model.User{
		Username: "admin",
	}
	{
		pwd, _ := bcrypt.GenerateFromPassword([]byte(constant.AdminDefaultPassword), bcrypt.DefaultCost)
		attrU := &model.User{
			Status:   model.UserStatusNormal,
			Password: string(pwd),
		}
		attrU.ID = util.SnowflakeId()
		if err := database.DB.Where(adminUser).Attrs(attrU).FirstOrCreate(adminUser).Error; err != nil {
			logger.Errorln(err)
		}
	}

	// 初始化一个超级管理员角色
	superAdminRole := &model.Role{
		Type:           model.RoleTypeSuperAdmin,
		DataPermission: model.RoleDataPermissionAll,
		Status:         model.RoleStatusNormal,
	}
	superAdminRole.ID = constant.SuperAdminRoleId
	{
		if err := database.DB.Where(superAdminRole).Attrs(&model.Role{
			Name: "超级管理员",
		}).FirstOrCreate(superAdminRole).Error; err != nil {
			logger.Errorln(err)
		}
	}

	// 关联admin和超级管理员角色
	{
		userRole := &model.UserRole{
			UserID: adminUser.ID,
			RoleID: superAdminRole.ID,
		}
		attrsUR := new(model.UserRole)
		attrsUR.ID = util.SnowflakeId()
		if err := database.DB.Where(userRole).Attrs(attrsUR).FirstOrCreate(userRole).Error; err != nil {
			logger.Errorln(err)
		}
	}

	// 初始化用户中心的资源
	var resources []*model.Resource

	// 后台管理
	{
		// 首页
		{
			resources = append(resources,
				&model.Resource{
					ResourceName: "首页",
					ResourceCode: constant.ResourceHome,
					Type:         1,
				})

			// 仪表盘
			{
				resources = append(resources, &model.Resource{
					ResourceName: "仪表盘",
					ResourceCode: constant.ResourceDashboard,
					Type:         1,
				})
			}
		}
		// 目录管理
		{
			resources = append(resources, &model.Resource{
				ResourceName: "目录管理",
				ResourceCode: constant.ResourceCategory,
				Type:         1,
			})
			resources = append(resources, &model.Resource{
				ResourceName: "增加目录",
				ResourceCode: constant.ResourceCategoryAdd,
				Type:         1,
			})
			resources = append(resources, &model.Resource{
				ResourceName: "修改目录",
				ResourceCode: constant.ResourceCategoryUpdate,
				Type:         1,
			})
			resources = append(resources, &model.Resource{
				ResourceName: "删除目录",
				ResourceCode: constant.ResourceCategoryDelete,
				Type:         1,
			})
			resources = append(resources, &model.Resource{
				ResourceName: "目录列表",
				ResourceCode: constant.ResourceCategoryList,
				Type:         1,
			})
			resources = append(resources, &model.Resource{
				ResourceName: "目录详情",
				ResourceCode: constant.ResourceCategoryDetail,
				Type:         1,
			})
		}
		// 用户管理
		{
			resources = append(resources, &model.Resource{
				ResourceName: "用户管理",
				ResourceCode: constant.ResourceUser,
				Type:         1,
			})
			resources = append(resources, &model.Resource{
				ResourceName: "增加用户",
				ResourceCode: constant.ResourceUserAdd,
				Type:         1,
			})
			resources = append(resources, &model.Resource{
				ResourceName: "修改用户",
				ResourceCode: constant.ResourceUserUpdate,
				Type:         1,
			})
			resources = append(resources, &model.Resource{
				ResourceName: "删除用户",
				ResourceCode: constant.ResourceUserDelete,
				Type:         1,
			})
			resources = append(resources, &model.Resource{
				ResourceName: "用户列表",
				ResourceCode: constant.ResourceUserList,
				Type:         1,
			})
			resources = append(resources, &model.Resource{
				ResourceName: "用户详情",
				ResourceCode: constant.ResourceUserDetail,
				Type:         1,
			})
			resources = append(resources, &model.Resource{
				ResourceName: "重置密码",
				ResourceCode: constant.ResourceUserResetPassword,
				Type:         1,
			})
			resources = append(resources, &model.Resource{
				ResourceName: "分配角色",
				ResourceCode: constant.ResourceUserAssignRole,
				Type:         1,
			})
		}
	}

	for _, resource := range resources {
		//没有就添加资源
		attrR := new(model.Resource)
		attrR.ID = util.SnowflakeId()
		if err := database.DB.Where(resource).Attrs(attrR).FirstOrCreate(resource).Error; err != nil {
			logger.Errorln(err)
		}
	}
}
