package initialize

import (
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

func Gorm() *gorm.DB {
	return GormSqlite()
}

func RegisterTables() {
	db := global.DB
	err := db.AutoMigrate(
		system.StableDiffusion{},
		system.StableDiffusionLoras{},
	)
	if err != nil {
		global.Log.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}

	err = bizModel(db)

	if err != nil {
		global.Log.Error("register biz_table failed", zap.Error(err))
		os.Exit(0)
	}
	global.Log.Info("初始化数据库表成功")
}
