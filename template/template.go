package template

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/SuCicada/su-action-server/model"
	"github.com/gin-gonic/gin"

	"github.com/SuCicada/su-action-server/logger"

	"github.com/SuCicada/su-action-server/utils"
)

var gotmpl *template.Template

func InitTemplate() {
	file := utils.Get("TEMPLATE_PATH")
	_gotmpl, err := template.ParseFiles(file)
	if err != nil {
		logger.Error(err)
	}
	logger.Info("template init success: " + file)
	gotmpl = _gotmpl
}

func GetTemplate(data interface{}) string {
	//name := "notification.gohtml"
	//gohtml, err := Asset(name)
	//if err != nil {
	//	panic(err)
	//	return ""
	//}
	//tmpl := template.Must(template.New(name).Parse(string(gohtml)))

	//if err != nil {
	//	fmt.Println(err)
	//	return ""
	//}

	buffer := new(strings.Builder)
	fmt.Println("data: ", data)
	err := gotmpl.Execute(buffer, data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return buffer.String()
}

func ParseRequest(c *gin.Context) string {
	action := model.Action{}
	err := c.Bind(&action)
	if err != nil {
		fmt.Println("parse request failed", err)
		return ""
	}
	github := action.Github

	status := action.Job.Status
	// github.event.commits[0].message
	commits := github.Event.Commits
	commitMessage := ""
	if len(commits) > 0 {
		commitMessage = commits[0].Message // github.Event["commits"].([]interface{})[0].(map[string]string)["message"]
	}
	// https://github.com/${{ github.repository }}/commit/${{github.sha}})
	commitUrl := fmt.Sprintf("https://github.com/%s/commit/%s", github.Repository, github.Sha)
	// https: //github.com/${{github.repository}}/actions/runs/${{github.run_id}}
	actionUrl := fmt.Sprintf("https://github.com/%s/actions/run/%s", github.Repository, github.RunID)

	return GetTemplate(map[string]interface{}{
		"status":        status,
		"commitMessage": commitMessage,
		"commitUrl":     commitUrl,
		"actionUrl":     actionUrl,
		"github":        github,
		"job":           action.Job,
	})
}
