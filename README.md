## what is this
这个是[su-action-webhook](https://github.com/SuCicada/su-action-webhook)的服务后端。
用于接受 github action 消息并推送到指定的提醒服务。
支持自定义模版
目前支持的提醒服务有：
1. Telegram Bot

## how to use
1. modify [notification.gohtml](notification.gohtml) to your love format
2. make build
3. upload all files in `dist` to your server

## Next Version
fully support [NoneBot2](https://github.com/nonebot/nonebot2)
