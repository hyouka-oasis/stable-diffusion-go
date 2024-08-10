package config

import (
	"fmt"
	"path/filepath"
)

type Sqlite struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (s *Sqlite) Dsn() string {
	//return filepath.Join(s.Path, s.Dbname+".db")
	fmt.Println(filepath.Join(ExecutePath, "stable-diffusion.db"), "数据库路径")
	return filepath.Join(ExecutePath, "stable-diffusion.db")
}
