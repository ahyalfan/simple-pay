package service

import (
	"ahyalfan/golang_e_money/domain"
	"ahyalfan/golang_e_money/internal/config"
	"net/smtp"
)

type emailService struct {
	cnf *config.Config
}

func NewEmail(cnf *config.Config) domain.EmailService {
	return &emailService{cnf: cnf}
}

// send implements domain.EmailService.
func (s *emailService) Send(to, subject, body string) error {
	// untuk identity kita kasih kosong karena biasanya smtp itu diisi kosongana
	auth := smtp.PlainAuth("", s.cnf.Mail.Username, s.cnf.Mail.Password, s.cnf.Mail.Host)

	// Mempersiapkan pesan email berbentu byte
	msg := []byte("From: simple-pay <" + s.cnf.Mail.Username + ">\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		body)
	// kirim email
	return smtp.SendMail(s.cnf.Mail.Host+":"+s.cnf.Mail.Port, auth, s.cnf.Mail.Username, []string{to}, msg)
}
