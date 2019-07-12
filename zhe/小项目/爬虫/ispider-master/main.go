package main

import (
	_ "ispider/routers"
	"ispider/spider"

	"github.com/Chain-Zhang/igo/ilog"
	"github.com/astaxie/beego"
)

func main() {
	ilog.Info("start")
	go spider.Start()
	beego.Run(":8089")
}
