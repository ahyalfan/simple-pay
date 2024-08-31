package service

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/dto"
	"ahyalfan/golang_e_money/internal/config"
	"encoding/json"
)

type emailService struct {
	cnf          *config.Config
	queueService domain.QueueService
}

func NewEmail(cnf *config.Config, queue domain.QueueService) domain.EmailService {
	return &emailService{cnf: cnf, queueService: queue}
}

// send implements domain.EmailService.

// ini yg sebelumnya tanpa antrian, yg diawal awal
// func (s *emailService) Send(to, subject, body string) error {
// 	// untuk identity kita kasih kosong karena biasanya smtp itu diisi kosongana
// 	auth := smtp.PlainAuth("", s.cnf.Mail.Username, s.cnf.Mail.Password, s.cnf.Mail.Host)

//		// Mempersiapkan pesan email berbentu byte
//		msg := []byte("From: simple-pay <" + s.cnf.Mail.Username + ">\n" +
//			"To: " + to + "\n" +
//			"Subject: " + subject + "\n" +
//			body)
//		// kirim email
//		return smtp.SendMail(s.cnf.Mail.Host+":"+s.cnf.Mail.Port, auth, s.cnf.Mail.Username, []string{to}, msg)
//	}
func (s *emailService) Send(to, subject, body string) error {
	payload := dto.SendEmailReq{
		To:      to,
		Subject: subject,
		Body:    body,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return s.queueService.Enqueue("send:email", data, 4)
}
