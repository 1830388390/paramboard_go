package main

import (
	"paramboard_go/api/v1/routes"
	"paramboard_go/conf"
)

func main() {
	conf.Init()
	r := routes.Router()
	_ = r.Run(conf.HttpPort)

}
