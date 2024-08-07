package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github/stable-diffusion-go/server/global"
	"os"
	"path"
)

func InitViper() *viper.Viper {
	pwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v := viper.New()
	v.SetConfigFile(path.Join(pwd, "config.yaml")) // name of config file (without extension)
	v.SetConfigType("yaml")                        // REQUIRED if the config file does not have the extension in the name
	err = v.ReadInConfig()                         // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.Config); err != nil {
		panic(err)
	}
	return v
}
