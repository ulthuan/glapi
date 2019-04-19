package models

import (
	"log"
)

// CardWebHook model that contains all the execution event logic
type CardWebHook struct {
	cardEvent *WebhookCardEvent
}

func (cwh CardWebHook) ExecuteAction(cEvent *WebhookCardEvent) {
	cwh.cardEvent = cEvent
	log.Println(cwh.cardEvent.Action)
}
