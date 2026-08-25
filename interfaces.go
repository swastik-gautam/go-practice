package main

import "fmt"

// Interface
type Notifier interface {
	Send(message string)
}

// Email implementation
type Email struct{}

func (e Email) Send(message string) {
	fmt.Println("Sending Email:", message)
}

// SMS implementation
type SMS struct{}

func (s SMS) Send(message string) {
	fmt.Println("Sending SMS:", message)
}

// This function accepts any type
// that implements Notifier
func Notify(n Notifier, message string) {
	n.Send(message)
}

func main() {
	email := Email{}
	sms := SMS{}

	Notify(email, "Hello from Go!")
	Notify(sms, "Your OTP is 1234")
}
