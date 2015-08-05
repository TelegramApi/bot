# Bot
Telegram bot API library written in golang

# Example

```go
package main

import (
	"encoding/json"
	tb "github.com/TelegramApi/bot"
)

func main() {
	bot := tb.Create("API_TOKEN")
	bot.Listen()

	for update := range bot.Updates {

		var outputMessage string

		switch update.Message.Text {
		case "/start":
			outputMessage = "I am your new Bot.\n\n"
		case "Hi, Bot!":
			outputMessage = "Hello, " + update.Message.From.FirstName
		default:
			outputMessage = ""
		}

		var chat tb.User
		json.Unmarshal(update.Message.Chat, &chat)

		var keyboard = tb.ReplyKeyboardMarkup{Keyboard: [][]string{[]string{"Hi, Bot!"}}}
		bot.SendMessage(chat.Id, outputMessage, false, 0, keyboard)
	}
}
```