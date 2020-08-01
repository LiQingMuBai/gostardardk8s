package initialize

import (
	"gostardardk8s/pkg/global"
	"gostardardk8s/pkg/model"
)

// 注册数据库表专用
func DBTables() {
	db := global.GVA_DB
	db.AutoMigrate(model.SysUser{},
		model.SysAuthority{},
		model.SysApi{},
		model.SysBaseMenu{},
		model.JwtBlacklist{},
		model.ExaFileUploadAndDownload{},
		model.ExaFile{},
		model.ExaFileChunk{},

	)
	global.GVA_LOG.Debug("register table success")
}
