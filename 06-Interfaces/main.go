package main

import "fmt"

type Notifier interface {
	Send(message string)
}
type EmailUser struct {
	Email string
}
type SMSUser struct {
	PhoneNumber string
}

func (e EmailUser) Send(massage string) {
	fmt.Printf("Sending Email to %v : %v \n ", e.Email, massage)

}

func (s SMSUser) Send(massage string) {
	fmt.Printf("Sending SMS to %v : %v \n ", s.PhoneNumber, massage)

}

func Notify(n Notifier, msg string) {
	n.Send(msg)
}

func main() {
	email := EmailUser{
		Email: "yasser@gmail.com",
	}
	sms := SMSUser{
		PhoneNumber: "01000002200",
	}
	Notify(email, "Hi")
	Notify(sms, "Welcome")
}
