package domains

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type SenderRepository interface {
	Store(sender Sender) error
	FindById(id int) (Sender, error)
}

type MessageRepository interface {
	Store(message Message) error
}

type Sender struct {
	ID   int64
	Name string `json:"sender"`
}

type Message struct {
	ID        int64
	Timestamp string            `json:"ts"`
	Message   MessageAttributes `json:"message"` // change to jsonb
	IP        string            `json:"sent-from-ip"`
	Priority  int               `json:"priority"`
	Sender    Sender            `json:"-"`
}

type MessageAttributes map[string]interface{}

func (this MessageAttributes) Value() (driver.Value, error) {
	return json.Marshal(this)
}

func (this *MessageAttributes) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to [] failed")
	}
	return json.Unmarshal(b, &this)
}
