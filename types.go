package tmpmailgo

type Email struct {
	username string
	domain   string
}

type Domains []string

type MessageSummary struct {
	ID      int    `json:"id"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Date    string `json:"date"`
}

type Message struct {
	ID          int          `json:"id"`
	From        string       `json:"from"`
	Subject     string       `json:"subject"`
	Date        string       `json:"date"`
	Attachments []Attachment `json:"attachments"`
	Body        string       `json:"body"`
	TextBody    string       `json:"textBody"`
	HTMLBody    string       `json:"htmlBody"`
}

type Attachment struct {
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
	Size        int    `json:"size"`
}
