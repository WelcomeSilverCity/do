package global

import (
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"

	"github.com/WelcomeSilverCity/do/config"
)

var (
	Db        *gorm.DB
	AllConfig config.Config
	EsClient  *elastic.Client
)
