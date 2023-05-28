package main

import (
	"github.com/NHollmann/yotei/controller"
)

func main() {
	var server controller.YoteiServer
	server.Start("127.0.0.1:9000", controller.DatabaseConfig{
		Driver:   "sqlite",
		Database: "yotei.db",
	})
}
