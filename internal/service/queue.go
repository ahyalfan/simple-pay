package service

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/internal/component"
	"ahyalfan/golang_e_money/internal/config"

	"github.com/hibiken/asynq"
)

// ini sangat disarankan apapun itu jenis servicenya, jika servicenya butuh waktu yg lumayan bisa pakai queue
type queueService struct {
	queueClient *asynq.Client
}

func NewQueueService(cnf *config.Config) domain.QueueService {
	// ini kita samakan dengan server yg sudah dibuat
	redisConn := asynq.RedisClientOpt{
		Addr:     cnf.Queue.Addr,
		Password: cnf.Queue.Password,
	}

	client := asynq.NewClient(redisConn)
	return &queueService{
		queueClient: client,
	}
}

// Enqueue implements domain.QueueService.
func (q *queueService) Enqueue(queueName string, data []byte, retyr int) error {
	task := asynq.NewTask(queueName, data, asynq.MaxRetry(retyr)) // retry optional // jika tidak dikasih akan nyoba terus

	info, err := q.queueClient.Enqueue(task) // ada option lainya itu bebas mau diisi atau tidak
	if err != nil {
		component.Log.Error("error when  enqueu: ", err.Error())
	}
	component.Log.Infof("task %s enqueued with ID: %s", queueName, info.ID)
	return err
}
