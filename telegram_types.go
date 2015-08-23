package telegram_bot

import "encoding/json"

type Update struct {
	UpdateId int     `json:"update_id"` //The update‘s unique identifier
	Message  Message `json:"message"`   //Optional. New incoming message of any kind — text, photo, sticker, etc.
}

type User struct {
	Id        int    `json:"id"`         //Unique identifier for this user or bot
	FirstName string `json:"first_name"` //User‘s or bot’s first name
	LastName  string `json:"last_name"`  //Optional. User‘s or bot’s last name
	Username  string `json:"username"`   //Optional. User‘s or bot’s username
}

type GroupChat struct {
	Id    int    `json:"id"`    //Unique identifier for this group chat
	Title string `json:"title"` //Group name
}

type Message struct {
	MessageID           int             `json:"message_id"`            //Unique message identifier
	From                User            `json:"from"`                  //Sender
	Date                int             `json:"date"`                  //Date the message was sent in Unix time
	Chat                json.RawMessage `json:"chat"`                  //Conversation the message belongs to — user in case of a private message, GroupChat in case of a group
	ForwardFrom         User            `json:"forward_from"`          //Optional. For forwarded messages, sender of the original message
	ForwardDate         int             `json:"forward_date"`          //Optional. For forwarded messages, date the original message was sent in Unix time
	ReplyToMessage      *Message        `json:"reply_to_message"`      //Optional. For replies, the original message.
	Text                string          `json:"text"`                  //Optional. For text messages, the actual UTF-8 text of the message
	Audio               Audio           `json:"audio"`                 //Optional. Message is an audio file, information about the file
	Document            Document        `json:"document"`              //Optional. Message is a general file, information about the file
	Photo               []PhotoSize     `json:"photo"`                 //Optional. Message is a photo, available sizes of the photo
	Sticker             Sticker         `json:"sticker"`               //Optional. Message is a sticker, information about the sticker
	Video               Video           `json:"video"`                 //Optional. Message is a video, information about the video
	Caption             string          `json:"caption"`               //Optional. Caption for the photo or video
	Contact             Contact         `json:"contact"`               //Optional. Message is a shared contact, information about the contact
	Location            Location        `json:"location"`              //Optional. Message is a shared location, information about the location
	NewChatParticipant  User            `json:"new_chat_participant"`  //Optional. A new member was added to the group, information about them (this member may be bot itself)
	LeftChatParticipant User            `json:"left_chat_participant"` //Optional. A member was removed from the group, information about them (this member may be bot itself)
	NewChatTitle        string          `json:"new_chat_title"`        //Optional. A group title was changed to this value
	NewChatPhoto        []PhotoSize     `json:"new_chat_photo"`        //Optional. A group photo was change to this value
	DeleteChatPhoto     bool            `json:"delete_chat_photo"`     //Optional. Informs that the group photo was deleted
	GroupChatCreated    bool            `json:"group_chat_created"`    //Optional. Informs that the group has been created
}

type PhotoSize struct {
	FileId   string `json:"file_id"`   //Unique identifier for this file
	Width    int    `json:"width"`     //Photo width
	Height   int    `json:"height"`    //Photo height
	FileSize int    `json:"file_size"` //Optional. File size
}

type Audio struct {
	FileId   string `json:"file_id"`   //Unique identifier for this file
	Duration int    `json:"duration"`  //Duration of the audio in seconds as defined by sender
	MimeType string `json:"mime_type"` //Optional. MIME type of the file as defined by sender
	FileSize int    `json:"file_size"` //Optional. File size
}

type Document struct {
	FileId   string    `json:"file_id"`   //Unique file identifier
	Thumb    PhotoSize `json:"thumb"`     //Optional. Document thumbnail as defined by sender
	FileName int       `json:"file_name"` //Optional. Original filename as defined by sender
	MimeType string    `json:"mime_type"` //Optional. MIME type of the file as defined by sender
	FileSize int       `json:"file_size"` //Optional. File size
}

type Sticker struct {
	FileId   string    `json:"file_id"`   //Unique identifier for this file
	Width    int       `json:"width"`     //Sticker width
	Height   int       `json:"height"`    //Sticker height
	Thumb    PhotoSize `json:"thumb"`     //Optional. Sticker thumbnail in .webp or .jpg format
	FileSize int       `json:"file_size"` //Optional. File size
}

type Video struct {
	FileId   string    `json:"file_id"`   //Unique identifier for this file
	Width    int       `json:"width"`     //Video width as defined by sender
	Height   int       `json:"height"`    //Video height as defined by sender
	Duration int       `json:"duration"`  //Duration of the video in seconds as defined by sender
	Thumb    PhotoSize `json:"thumb"`     //Optional. Video thumbnail
	MimeType string    `json:"mime_type"` //Optional. Mime type of a file as defined by sender
	FileSize int       `json:"file_size"` //Optional. File size
}

type Contact struct {
	PhoneNumber string `json:"phone_number"` //Contact's phone number
	FirstName   string `json:"first_name"`   //Contact's first name
	LastName    string `json:"last_name"`    //Optional. Contact's last name
	UserId      int    `json:"user_id"`      //Optional. Contact's user identifier in Telegram
}

type Location struct {
	Longitude float32 `json:"longitude"` //Longitude as defined by sender
	Latitude  float32 `json:"latitude"`  //Latitude as defined by sender
}

type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"` //Total number of profile pictures the target user has
	Photos     [][]PhotoSize `json:"photos"`      //Requested profile pictures (in up to 4 sizes each)
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

type InputFile []byte
