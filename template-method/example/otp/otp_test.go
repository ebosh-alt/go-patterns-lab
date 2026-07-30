package main

import (
	"errors"
	"reflect"
	"testing"
)

// flowSpy записывает вызванные шаги шаблонного алгоритма.
type flowSpy struct {
	steps   []string
	sendErr error
}

func (s *flowSpy) Generate(length int) string {
	s.steps = append(s.steps, "generate")
	return "1234"
}

func (s *flowSpy) Save(code string) {
	s.steps = append(s.steps, "save")
}

func (s *flowSpy) Message(code string) string {
	s.steps = append(s.steps, "message")
	return "Код: " + code
}

func (s *flowSpy) Send(message string) error {
	s.steps = append(s.steps, "send")
	return s.sendErr
}

// TestOTPServicePreservesStepOrder проверяет неизменность порядка шагов.
func TestOTPServicePreservesStepOrder(t *testing.T) {
	flow := &flowSpy{}
	service := NewOTPService(flow)

	if err := service.GenerateAndSend(4); err != nil {
		t.Fatalf("GenerateAndSend returned an error: %v", err)
	}

	want := []string{"generate", "save", "message", "send"}
	if !reflect.DeepEqual(flow.steps, want) {
		t.Fatalf("steps = %v, want %v", flow.steps, want)
	}
}

// TestOTPServiceReturnsSendError проверяет возврат ошибки конкретной реализации.
func TestOTPServiceReturnsSendError(t *testing.T) {
	wantErr := errors.New("notification service is unavailable")
	flow := &flowSpy{sendErr: wantErr}
	service := NewOTPService(flow)

	err := service.GenerateAndSend(4)

	if !errors.Is(err, wantErr) {
		t.Fatalf("error = %v, want %v", err, wantErr)
	}
}
