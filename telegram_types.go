package telegram_bot

import "encoding/json"

type Update struct {
	UpdateId int     `json:"update_id"` //The update‘s unique identifier
	Message  Message `json:"message"`   //Optional. New incoming message of any kind — text, photo, sticker, etc.
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

type GroupChat struct {
	Id    int    `json:"id"`    //Unique identifier for this group chat
	Title string `json:"title"` //Group name
}

type Message struct {
	MessageID           int             `json:"message_id"`   //Unique message identifier
	From                User            `json:"from"`         //Sender
	Date                int             `json:"date"`         //Date the message was sent in Unix time
	Chat                json.RawMessage `json:"chat"`         //Conversation the message belongs to — user in case of a private message, GroupChat in case of a group
	ForwardFrom         User            `json:"forward_from"` //Optional. For forwarded messages, sender of the original message
	ForwardDate         int             `json:"forward_date"` //Optional. For forwarded messages, date the original message was sent in Unix time
	ReplyToMessage      *Message        `json:"reply_to_message"`
	Text                string          `json:"text"`
	Audio               Audio           `json:"audio"`
	Document            Document        `json:"document"`
	Photo               []PhotoSize     `json:"photo"`
	Sticker             Sticker         `json:"sticker"`
	Video               Video           `json:"video"`
	Caption             string          `json:"caption"`
	Contact             Contact         `json:"contact"`
	Location            Location        `json:"location"`
	NewChatParticipant  User            `json:"new_chat_participant"`
	LeftChatParticipant User            `json:"left_chat_participant"`
	NewChatTitle        string          `json:"new_chat_title"`
	NewChatPhoto        []PhotoSize     `json:"new_chat_photo"`
	DeleteChatPhoto     bool            `json:"delete_chat_photo"`
	GroupChatCreated    bool            `json:"group_chat_created"`
}

type PhotoSize struct {
	FileId   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}

type Audio struct {
	FileId   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
}

type Document struct {
}

type Sticker struct {
}

type Video struct {
}

type Contact struct {
}

type Location struct {
}

type UserProfilePhotos struct {
}

type ReplyKeyboardMarkup struct {
	Keyboard       [][]string `json:"keyboard"`        //Array of button rows, each represented by an Array of Strings
	ResizeKeyboard bool       `json:"resize_keyboard"` //Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). 	OneTimeKeyboard bool       `json:"one_time_keyboard"` //Optional. Requests clients to hide the keyboard as soon as it's been used. Defaults to false.
	Selective      bool       `json:"selective"`       //Optional. Use this parameter if you want to show the keyboard to specific users only.
}

type ReplyKeyboardHide struct {
	HideKeyboard bool `json:"hide_keyboard"` //Requests clients to hide the custom keyboard
	Selective    bool `json:"selective"`     //Optional. Use this parameter if you want to hide keyboard for specific users only.
}

type ForceReply struct {
	ForceReply bool `json:"force_reply"` //Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
	Selective  bool `json:"selective"`   //Optional. Use this parameter if you want to force reply from specific users only.
}

type InputFile struct {
}
