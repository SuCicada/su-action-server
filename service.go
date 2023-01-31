package main

import (
	"html"
	"strconv"

	"github.com/SuCicada/su-action-server/template"

	"github.com/SuCicada/su-action-server/utils"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func GitHubActionNotify(c *gin.Context) {
	// user := c.Params.ByName("name")

	// var req = map[string]interface{}{}
	// c.BindJSON(&req)
	// req[""] = "test"
	// req["github"].(map[string]interface{})["event"].(map[string]interface{})["repository"]
	// fmt.Println(c)

	user := utils.Get("TELEGRAM_TO")
	text := template.ParseRequest(c)
	SendTelegramMessage(user, text)
}

func SendTelegramMessage(user string, txt string) {
	//txt, err := template.RenderTrim(value, p)
	//if err != nil {
	//	return err
	//}
	token := utils.Get("TELEGRAM_TOKEN")
	bot, _ := tgbotapi.NewBotAPI(token)
	txt = html.UnescapeString(txt)
	id, _ := strconv.ParseInt(user, 10, 64)
	msg := tgbotapi.NewMessage(id, txt)
	msg.ParseMode = "markdown"
	msg.DisableWebPagePreview = false
	msg.DisableNotification = false

	bot.Send(msg)
}
