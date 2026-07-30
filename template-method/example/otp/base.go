package main

import (
	"fmt"
	"strings"
)

// BaseOTPFlow содержит общую реализацию одинаковых шагов.
type BaseOTPFlow struct {
	channel string
}

// NewBaseOTPFlow создаёт базовую реализацию для указанного канала.
func NewBaseOTPFlow(channel string) BaseOTPFlow {
	return BaseOTPFlow{channel: channel}
}

// Generate создаёт демонстрационный одноразовый код.
func (f *BaseOTPFlow) Generate(length int) string {
	code := strings.Repeat("1", length)
	fmt.Printf("%s: создан код %s\n", f.channel, code)

	return code
}

// Save имитирует сохранение кода для последующей проверки.
func (f *BaseOTPFlow) Save(code string) {
	fmt.Printf("%s: код %s сохранён\n", f.channel, code)
}
