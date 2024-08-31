package domain

type QueueService interface {
	// daftarkan ke antrian
	Enqueue(queueName string, data []byte, retyr int) error
}
