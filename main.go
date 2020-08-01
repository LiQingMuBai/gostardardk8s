package main

import (
	"gostardardk8s/pkg/core"
	"gostardardk8s/pkg/global"
	"gostardardk8s/initialize"
)


func main() {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		initialize.Mysql()
	default:
		initialize.Mysql()
	}
	initialize.DBTables()

	defer global.GVA_DB.Close()
	core.RunWindowsServer()
}
