package define

import (
	"github.com/wqh66886/past-present-future/common/model"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Cfg *Config
)

type Config struct {
	Mysql *model.Mysql `json:"mysql"`
}
