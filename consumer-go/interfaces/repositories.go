package interfaces

import (
	"database/sql"
	"log"

	"github.com/eduardosanzb/consumer/domains"
	"github.com/eduardosanzb/consumer/infrastructure"
)

type SenderRepository struct{}
type MessageRepository struct{}

func (repo *SenderRepository) Store(sender *domains.Sender) error {
	statement, err := infrastructure.DB.Prepare("INSERT INTO senders(name) VALUES($1) RETURNING id;")
	log.Println(statement)

	if err != nil {
		log.Println("Error when trying to prepare statement for Store sender")
		log.Println(err)
		return err
	}
	defer statement.Close()

	var insertedID int64
	if err := statement.QueryRow(sender.Name).Scan(&insertedID); err != nil {
		log.Println("Error when trying to insert a sender")
		log.Println(err)
		return err
	}

	log.Printf("insertedID: %v, %T", insertedID, insertedID)
	sender.ID = insertedID
	return nil
}

func (repo *SenderRepository) FindByName(name string) (*domains.Sender, error) {
	statement, err := infrastructure.DB.Prepare("SELECT id FROM  senders WHERE name=$1;")
	if err != nil {
		log.Println("Error when trying to prepare statement to find sender")
		log.Println(err)
		return nil, err
	}
	defer statement.Close()

	result := statement.QueryRow(name)

	sender := domains.Sender{}
	err = result.Scan(&sender.ID)
	if err != nil && err != sql.ErrNoRows {
		log.Println("Error quen trying to get a Sender by Name")
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &sender, nil
}

func (repo *MessageRepository) Store(message domains.Message) error {
	statement, err := infrastructure.DB.Prepare(
		"INSERT INTO " +
			"messages(timestamp, message, ip, priority, sender_id) " +
			"values(to_timestamp($1), $2, $3, $4, $5) " +
			"RETURNING id;")

	if err != nil {
		log.Println("Error when trying to prepare statement for Store Message")
		log.Println(err)
		return err
	}
	defer statement.Close()

	var insertedID int64
	if err := statement.QueryRow(
		message.Timestamp,
		message.Message,
		message.IP,
		message.Priority,
		message.Sender.ID,
	).Scan(&insertedID); err != nil {
		log.Println("Erro when trying to insert a message")
		log.Println(err)
		return err
	}

	message.ID = insertedID
	return nil
}
