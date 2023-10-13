package main

import (
	"github.com/panjf2000/ants/v2"
	logger "github.com/sirupsen/logrus"
	"manatee-publish/internal/controller"
	"manatee-publish/internal/data"
	"manatee-publish/pkg/config"
	"manatee-publish/pkg/database"
	"manatee-publish/pkg/server"
	"manatee-publish/pkg/task"
	"manatee-publish/pkg/util"
)

var VERSION = "unknown"

func main() {
	defer ants.Release()

	config.InitialLogger()

	database.Initial()
	defer database.Close()

	_ = util.InitNode(0)

	// 初始化数据
	data.InitData()

	task.Start()
	defer task.Stop()

	controller.InitRouter()
	controller.InitManage()

	for {
		err := server.Start()
		if err != nil {
			logger.Errorln(err)
		}
	}
}
