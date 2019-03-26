package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"github.com/ulthuan/go-glo"
)

type WebhookBoard struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type WebhookSender struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
}

type WebhookArchived struct {
	Action string        `json:"action"`
	Board  WebhookBoard  `json:"board"`
	Sender WebhookSender `json:"sender"`
}

type WebhookCardAdded struct{}
type WebhookCardUpdate struct{}
type WebhookCardCopied struct{}

type Webhook struct {
	ID             uuid.UUID       `json:"id" db:"id"`
	Action         string          `json:"action" db:"action"`
	BoardID        string          `json:"board.id" db:"board_id"`
	Board          glo.Board       `json:"board"`
	Sender         glo.User        `json:"sender"`
	SenderUsername string          `json:"sender.username" db:"sender_username"`
	Members        []glo.Member    `json:"members"`
	Column         glo.BoardColumn `json:"column"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (w Webhook) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Webhooks is not required by pop and may be deleted
type Webhooks []Webhook

// String is not required by pop and may be deleted
func (w Webhooks) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (w *Webhook) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (w *Webhook) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (w *Webhook) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
