package line

type LineMessage struct {
	Destination string             `json:"destination"`
	Events      []LineMessageEvent `json:"events"`
}

type LineMessageEvent struct {
	ReplyToken string `json:"replyToken"`
	Type       string `json:"type"`
	Timestamp  int64  `json:"timestamp"`
	Source     struct {
		Type   string `json:"type"`
		UserID string `json:"userId"`
	} `json:"source"`
	Message Message `json:"message"`
}

type BotBrain struct {
	InputMessage Message
	ReplyMessage ReplyMessage
	Action       string
	Title        string
	Code         string
}

type Message struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Text      string `json:"text,omitempty"`
	PackageId int    `json:"packageId,omitempty"`
	StickerId int    `json:"stickerId,omitempty"`
}

type ReplyMessage struct {
	Message
	ReplyToken string `json:"replyToken"`
}
