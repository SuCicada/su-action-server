package template

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/SuCicada/su-action-server/logger"
	"github.com/SuCicada/su-action-server/model"
	"github.com/SuCicada/su-action-server/utils"
)

func TestGetTemplate(t *testing.T) {
	utils.InitEnv("../.env")
	logger.InitLog()
	InitTemplate()
	jsonFile, _ := os.ReadFile("../test/request.json")
	action := model.Action{}
	json.Unmarshal(jsonFile, &action)
	// fmt.Println(action)
	res := GetTemplate(map[string]interface{}{
		//"status":        "success",
		"status":        "failure",
		"commitMessage": "test",
		"commitUrl":     "asfasf",
		"actionUrl":     "asfasf",
		"github":        utils.InterfaceToMap(action.Github),
		"job":           utils.InterfaceToMap(action.Job),
	})
	fmt.Println(res)
}
