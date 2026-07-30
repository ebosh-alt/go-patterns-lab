package main

// OTPFlow описывает шаги, реализация которых зависит от способа доставки.
type OTPFlow interface {
	// Generate создаёт одноразовый код заданной длины.
	Generate(length int) string

	// Save сохраняет код для последующей проверки.
	Save(code string)

	// Message формирует сообщение с одноразовым кодом.
	Message(code string) string

	// Send доставляет подготовленное сообщение пользователю.
	Send(message string) error
}

// OTPService содержит шаблонный метод с неизменным порядком шагов.
type OTPService struct {
	flow OTPFlow
}

// NewOTPService создаёт сервис с выбранным способом доставки.
func NewOTPService(flow OTPFlow) *OTPService {
	return &OTPService{flow: flow}
}

// GenerateAndSend выполняет общий алгоритм создания и отправки OTP.
func (s *OTPService) GenerateAndSend(length int) error {
	code := s.flow.Generate(length)
	s.flow.Save(code)
	message := s.flow.Message(code)

	return s.flow.Send(message)
}
