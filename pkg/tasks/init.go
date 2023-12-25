package tasks

import (
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB
	queueClient *asynq.Client
)

func AddDB(database *gorm.DB) {
	db = database
}

func AddQueueClient(client *asynq.Client) {
	queueClient = client
}
