package common

import (
	"github.com/SuCicada/su-action-server/logger"
	"github.com/SuCicada/su-action-server/template"
	"github.com/SuCicada/su-action-server/utils"
)

func Init() {
	utils.InitEnv("")
	logger.InitLog()
	template.InitTemplate()
}
