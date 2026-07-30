package main

import "fmt"

// EmailFlow определяет шаги подготовки и отправки письма.
type EmailFlow struct {
	BaseOTPFlow
}

// NewEmailFlow создаёт сценарий доставки одноразового кода по email.
func NewEmailFlow() *EmailFlow {
	return &EmailFlow{
		BaseOTPFlow: NewBaseOTPFlow("Email"),
	}
}

// Message формирует текст письма.
func (f *EmailFlow) Message(code string) string {
	return "Ваш одноразовый код: " + code
}

// Send имитирует отправку письма.
func (f *EmailFlow) Send(message string) error {
	fmt.Printf("Email: отправлено письмо %q\n", message)
	return nil
}
