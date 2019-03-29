package models

import (
	"encoding/json"
	"image/color"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
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

type WebhookLabel struct {
	ID    string     `json:"id"`
	Name  string     `json:"name"`
	Color color.RGBA `json:"color"`
}

type WebhookPartialUser struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type WebhookDescription struct {
	Text string `json:"text"`
}

type WebhookPartialCard struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description WebhookDescription `json:"description"`
}

type WebhookLabels struct {
	Added   []WebhookLabel `json:"added"`
	Removed []WebhookLabel `json:"removed"`
}

type WebhookAssignees struct {
	Added   []WebhookPartialUser `json:"added"`
	Removed []WebhookPartialUser `json:"removed"`
}

type WebhookCard struct {
	ID                 string             `json:"id"`
	Name               string             `json:"name"`
	CreatedDate        time.Time          `json:"created_date"`
	BoardID            string             `json:"board_id"`
	ColumnID           string             `json:"column_id"`
	Labels             []WebhookLabel     `json:"labels"`
	Assignees          []WebhookAssignees `json:"assignees"`
	CompletedTaskCount int                `json:"completed_task_count"`
	TotalTaskCount     int                `json:"total_task_count"`
	AttachmentCount    int                `json:"attachment_count"`
	CommentCount       int                `json:"comment_count"`
	Description        WebhookDescription `json:"description"`
	CreatedBy          WebhookPartialUser `json:"created_by"`
	Previous           WebhookPartialCard `json:"previous"`
	Position           int                `json:"position"`
}

type WebhookCardEvent struct {
	Action    string           `json:"action"`
	Board     WebhookBoard     `json:"board"`
	Sender    WebhookSender    `json:"sender"`
	Card      WebhookCard      `json:"card"`
	NewCard   WebhookCard      `json:"new_card"`
	OldCard   WebhookCard      `json:"old_card"`
	Labels    WebhookLabels    `json:"labels"`
	Assigness WebhookAssignees `json:"assignees"`
	Sequence  int              `json:"sequence"`
}

type Webhook struct {
	ID             uuid.UUID `json:"id" db:"id"`
	Action         string    `json:"action" db:"action"`
	BoardID        string    `json:"board.id" db:"board_id"`
	CardID         string    `json:"card.id" db:"card_id"`
	SenderUsername string    `json:"sender.username" db:"sender_username"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
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
