package main

import "fmt"

func main() {
	// Клиент выбирает конкретную реализацию шагов для доставки по SMS.
	smsService := NewOTPService(NewSMSFlow())
	if err := smsService.GenerateAndSend(4); err != nil {
		fmt.Printf("ошибка отправки SMS: %v\n", err)
	}

	fmt.Println()

	// Общий алгоритм не меняется при выборе доставки по email.
	emailService := NewOTPService(NewEmailFlow())
	if err := emailService.GenerateAndSend(4); err != nil {
		fmt.Printf("ошибка отправки email: %v\n", err)
	}
}
