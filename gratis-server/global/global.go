package global

import (
	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
	"github.com/spf13/viper"
	"github/stable-diffusion-go/server/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Viper              *viper.Viper  // Viper实例
	DB                 *gorm.DB      // 数据库实例
	Log                *zap.Logger   //日志
	Config             config.Config // yaml配置文件
	Seg                gse.Segmenter
	PosSeg             pos.Segmenter
	ParticipleBookName = "participleBook.txt" //分词后的文件
)

// Animations 动画列表
var Animations = []string{"shrink", "magnify", "left_move", "right_move", "up_move", "down_move"}
