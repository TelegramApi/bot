package telegram_bot

import (
	"encoding/json"
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"strconv"
)

var BaseURL string = "https://api.telegram.org/bot%s/%s?%s"

type Bot struct {
	token   string
	Updates chan Update
}

type Response struct {
	Ok          bool            `json:"ok"`
	Description string          `json:"description"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
}

func (b *Bot) newRequest(method string, params url.Values) (Response, error) {
	res, err := http.Get(fmt.Sprintf(BaseURL, b.token, method, params.Encode()))
	if err != nil {
		return Response{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		return Response{}, err
	}

	if res.StatusCode != 200 {
		return Response{}, nil //todo
	}

	var resp Response
	json.Unmarshal(body, &resp)

	return resp, nil
}

func Create(token string) *Bot {
	return &Bot{token: token, Updates: make(chan Update)}
}

func (b *Bot) GetMe() User {
	resp, _ := b.newRequest("getMe", url.Values{})

	var user User
	json.Unmarshal(resp.Result, &user)
	return user
}

func (b *Bot) GetUpdates(offset int, limit int, timeout int) []Update {

	params := url.Values{}

	if offset != 0 {
		params.Set("offset", strconv.Itoa(offset))
	}

	params.Set("limit", strconv.Itoa(limit))
	params.Set("timeout", strconv.Itoa(timeout))

	resp, _ := b.newRequest("getUpdates", params)

	var updates []Update
	json.Unmarshal(resp.Result, &updates)
	return updates
}

func (b *Bot) Listen() {
	var offset int = 0
	var limit int = 100
	var timeout int = 10

	go func() {
		for {
			updates := b.GetUpdates(offset, limit, timeout)
			for _, update := range updates {
				if update.UpdateId >= offset {
					b.Updates <- update
					offset = update.UpdateId + 1
				}
			}
		}
	}()
}

func (b *Bot) SetWebhook(webhookUrl string) {
	params := url.Values{}
	params.Set("url", webhookUrl)
	b.newRequest("setWebhook", params)
}

//todo: implement optional parameters
func (b *Bot) SendMessage(chat_id int, text string, disable_web_page_preview bool, reply_to_message_id int, reply_markup interface{}) (Message, error) {
	params := url.Values{}
	params.Set("chat_id", strconv.Itoa(chat_id))
	params.Set("text", text)
	resp, err := b.newRequest("sendMessage", params)
	var message Message
	json.Unmarshal(resp.Result, &message)
	return message, err
}

func (b *Bot) forwardMessage() {
	panic("Not Implemented")
}

func (b *Bot) sendPhoto() {
	panic("Not Implemented")
}

func (b *Bot) sendAudio() {
	panic("Not Implemented")
}

func (b *Bot) sendDocument() {
	panic("Not Implemented")
}

func (b *Bot) sendSticker() {
	panic("Not Implemented")
}

func (b *Bot) sendVideo() {
	panic("Not Implemented")
}

func (b *Bot) sendLocation() {
	panic("Not Implemented")
}

func (b *Bot) sendChatAction() {
	panic("Not Implemented")
}

func (b *Bot) getUserProfilePhotos() {
	panic("Not Implemented")
}