package controller

import (
	"github.com/gofiber/fiber/v2"
	"manatee-publish/pkg/common"
	"manatee-publish/pkg/crypto"
	"manatee-publish/pkg/server"
)

var Controllers []common.RouterController

func InitRouter() {
	server.Get("/api/v1/pk", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&server.CommonResponse{
			Data: crypto.PublicKeyString(),
		})
	})
}

func InitManage() {
	for _, c := range Controllers {
		c.InitManage()
	}
}
