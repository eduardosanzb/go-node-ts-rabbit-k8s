package usecases

import (
	"encoding/json"
	"log"

	"github.com/eduardosanzb/consumer/domains"
	"github.com/eduardosanzb/consumer/interfaces"
)

func CreateMessageWithSender(rawJson []byte) bool {

	senderRepo := interfaces.SenderRepository{}
	messageRepo := interfaces.MessageRepository{}

	sender := domains.Sender{}
	message := domains.Message{}

	if err := json.Unmarshal(rawJson, &message); err != nil {
		log.Println("Error unmarshalling message")
		log.Println(err)
		return false
	}
	if err := json.Unmarshal(rawJson, &sender); err != nil {
		log.Println("Error unmarshalling sender")
		log.Println(err)
		return false
	}

	var err error
	var senderStoredInDB *domains.Sender
	senderStoredInDB, err = senderRepo.FindByName(sender.Name)
	if err != nil {
		log.Println(err)
		return false
	}

	if senderStoredInDB == nil {
		if err := senderRepo.Store(&sender); err != nil {
			return false
		}
	}
	if senderStoredInDB != nil {
		sender.ID = senderStoredInDB.ID
	}

	message.Sender = sender
	if err := messageRepo.Store(message); err != nil {
		log.Println(err)
		return false
	}

	return true
}
