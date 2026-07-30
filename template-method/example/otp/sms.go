package main

import "fmt"

// SMSFlow определяет шаги подготовки и отправки SMS.
type SMSFlow struct {
	BaseOTPFlow
}

// NewSMSFlow создаёт сценарий доставки одноразового кода по SMS.
func NewSMSFlow() *SMSFlow {
	return &SMSFlow{
		BaseOTPFlow: NewBaseOTPFlow("SMS"),
	}
}

// Message формирует текст SMS.
func (f *SMSFlow) Message(code string) string {
	return "Код для входа: " + code
}

// Send имитирует отправку SMS.
func (f *SMSFlow) Send(message string) error {
	fmt.Printf("SMS: отправлено сообщение %q\n", message)
	return nil
}
