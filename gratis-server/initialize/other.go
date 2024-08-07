package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github/stable-diffusion-go/server/global"
	"github/stable-diffusion-go/server/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.Config.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.Config.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
