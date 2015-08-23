package telegram_bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
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

func (b *Bot) newRequest(method string, params url.Values, file []byte) (Response, error) {

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	if file != nil {
		var f bytes.Buffer

		_, err := f.Write(file)

		fw, err := w.CreateFormFile("photo", "file.png")
		if err != nil {
			return Response{}, err
		}

		f.WriteTo(fw)
	}

	for name, val := range params {
		fw, err := w.CreateFormField(name)
		if err != nil {
			return Response{}, err
		}
		if _, err := fw.Write([]byte(val[0])); err != nil {
			return Response{}, err
		}
	}
	w.Close()

	client := &http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf(BaseURL, b.token, method), &buf)

	req.Header.Set("Content-Type", w.FormDataContentType())

	res, err := client.Do(req)

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

func (b *Bot) GetUpdates(offset int, limit int, timeout int) []Update {

	params := url.Values{}

	if offset != 0 {
		params.Set("offset", strconv.Itoa(offset))
	}

	params.Set("limit", strconv.Itoa(limit))
	params.Set("timeout", strconv.Itoa(timeout))

	resp, _ := b.newRequest("getUpdates", params, nil)

	var updates []Update
	json.Unmarshal(resp.Result, &updates)
	return updates
}

func (b *Bot) SetWebhook(webhookUrl string) {
	params := url.Values{}
	params.Set("url", webhookUrl)
	b.newRequest("setWebhook", params, nil)
}

func (b *Bot) GetMe() User {
	resp, _ := b.newRequest("getMe", url.Values{}, nil)

	var user User
	json.Unmarshal(resp.Result, &user)
	return user
}

func (b *Bot) ForwardMessage(chat_id, from_chat_id, message_id int) (Message, error) {
	params := url.Values{}
	params.Set("chat_id", strconv.Itoa(chat_id))
	params.Set("from_chat_id", strconv.Itoa(from_chat_id))
	params.Set("message_id", strconv.Itoa(message_id))

	resp, err := b.newRequest("forwardMessage", params, nil)
	var message Message
	json.Unmarshal(resp.Result, &message)

	return message, err
}

func (b *Bot) SendMessage(chat_id int, text string, disable_web_page_preview bool, reply_to_message_id int, reply_markup interface{}) (Message, error) {
	params := url.Values{}
	params.Set("chat_id", strconv.Itoa(chat_id))
	params.Set("text", text)
	params.Set("disable_web_page_preview", strconv.FormatBool(disable_web_page_preview))
	params.Set("reply_to_message_id", strconv.Itoa(reply_to_message_id))

	keyboard, err := json.Marshal(reply_markup)

	params.Set("reply_markup", string(keyboard))

	resp, err := b.newRequest("sendMessage", params, nil)
	var message Message
	json.Unmarshal(resp.Result, &message)
	return message, err
}

func (b *Bot) SendPhoto(chat_id int, photo []byte, caption string) (Message, error) {
	params := url.Values{}
	params.Set("chat_id", strconv.Itoa(chat_id))
	params.Set("caption", caption)

	resp, err := b.newRequest("sendPhoto", params, photo)
	var message Message
	json.Unmarshal(resp.Result, &message)

	return message, err
}

func (b *Bot) sendAudio(chat_id int, audio []byte, duration int, performer string, title string, reply_to_message_id int, reply_markup interface{}) (Message, error) {
	params := url.Values{}
	params.Set("chat_id", strconv.Itoa(chat_id))
	params.Set("duration", strconv.Itoa(duration))
	params.Set("reply_to_message_id", strconv.Itoa(reply_to_message_id))
	params.Set("performer", performer)
	params.Set("title", title)

	keyboard, _ := json.Marshal(reply_markup)
	params.Set("reply_markup", string(keyboard))

	resp, err := b.newRequest("sendAudio", params, audio)
	var message Message
	json.Unmarshal(resp.Result, &message)

	return message, err
}

func (b *Bot) sendDocument(chat_id int, document []byte, reply_to_message_id int, reply_markup interface{}) (Message, error) {
	params := url.Values{}
	params.Set("chat_id", strconv.Itoa(chat_id))
	params.Set("reply_to_message_id", strconv.Itoa(reply_to_message_id))

	keyboard, _ := json.Marshal(reply_markup)
	params.Set("reply_markup", string(keyboard))

	resp, err := b.newRequest("sendDocument", params, document)
	var message Message
	json.Unmarshal(resp.Result, &message)

	return message, err
}

func (b *Bot) sendSticker(chat_id int, sticker []byte, reply_to_message_id int, reply_markup interface{}) (Message, error) {
	params := url.Values{}
	params.Set("chat_id", strconv.Itoa(chat_id))
	params.Set("reply_to_message_id", strconv.Itoa(reply_to_message_id))

	keyboard, _ := json.Marshal(reply_markup)
	params.Set("reply_markup", string(keyboard))

	resp, err := b.newRequest("sendSticker", params, sticker)
	var message Message
	json.Unmarshal(resp.Result, &message)

	return message, err
}

func (b *Bot) sendVideo(chat_id int, video []byte, duration int, caption string, reply_to_message_id int, reply_markup interface{}) (Message, error)  {
	params := url.Values{}
	params.Set("chat_id", strconv.Itoa(chat_id))
	params.Set("duration", strconv.Itoa(duration))
	params.Set("caption", caption)
	params.Set("reply_to_message_id", strconv.Itoa(reply_to_message_id))

	keyboard, _ := json.Marshal(reply_markup)
	params.Set("reply_markup", string(keyboard))

	resp, err := b.newRequest("sendSticker", params, video)
	var message Message
	json.Unmarshal(resp.Result, &message)

	return message, err
}

func (b *Bot) sendLocation(chat_id int, latitude float64, longitude float64, reply_to_message_id int, reply_markup interface {}) (Message, error) {
	params := url.Values{}
	params.Set("chat_id", strconv.Itoa(chat_id))
	params.Set("latitude", strconv.FormatFloat(latitude, 'f', -1, 64))
	params.Set("longitude", strconv.FormatFloat(longitude, 'f', -1, 64))

	params.Set("reply_to_message_id", strconv.Itoa(reply_to_message_id))

	keyboard, _ := json.Marshal(reply_markup)
	params.Set("reply_markup", string(keyboard))

	resp, err := b.newRequest("sendLocation", params, nil)
	var message Message
	json.Unmarshal(resp.Result, &message)

	return message, err
}

func (b *Bot) sendChatAction(chat_id int, action string) (Message, error) {
	params := url.Values{}
	params.Set("chat_id", strconv.Itoa(chat_id))
	params.Set("action", action)

	resp, err := b.newRequest("sendChatAction", params, nil)
	var message Message
	json.Unmarshal(resp.Result, &message)

	return message, err
}

func (b *Bot) getUserProfilePhotos(user_id int, offset int, limit int) (UserProfilePhotos, error) {
	params := url.Values{}
	params.Set("user_id", strconv.Itoa(user_id))
	params.Set("offset", strconv.Itoa(offset))
	params.Set("limit", strconv.Itoa(limit))

	resp, err := b.newRequest("getUserProfilePhotos", params, nil)
	var photos UserProfilePhotos
	json.Unmarshal(resp.Result, &photos)

	return photos, err
}
