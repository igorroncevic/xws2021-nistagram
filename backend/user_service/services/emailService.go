package services

import (
	"context"
	"fmt"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"math/rand"
	"net/smtp"
)
type EmailService struct {
}

func NewEmailService() (*EmailService, error) {
	return &EmailService{},nil
}

func  (service *EmailService)  SendEmail(ctx context.Context, in string) (bool, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "UpdateUserProfile")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)


	// Sender data.
	from := "bsep2021@gmail.com"
	password := "BSEP2021"

	// Receiver email address.
	to := []string{
		"t.kovacevic98@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, 8)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	fmt.Println(string(s))

	message := []byte("Hi,\nwe received a request to reset your password.Your old password has been locked for security reasons.\nTo unlock your profile you must verify your identity.\n\nYour password reset code is:"+string(s));


	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return false,err
	}
	fmt.Println("Email Sent Successfully!")

	return true, nil
}
