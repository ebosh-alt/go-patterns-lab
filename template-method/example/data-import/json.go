package main

import (
	"encoding/json"
	"fmt"
)

// JSONFormat реализует изменяемые шаги импорта JSON.
type JSONFormat struct{}

// Parse разбирает JSON-массив записей.
func (f *JSONFormat) Parse(data []byte) ([]Record, error) {
	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		return nil, fmt.Errorf("разобрать JSON: %w", err)
	}

	return records, nil
}

// Validate проверяет обязательные поля записей JSON.
func (f *JSONFormat) Validate(records []Record) error {
	return validateUserRecords(records)
}

// Transform преобразует записи JSON в пользователей.
func (f *JSONFormat) Transform(records []Record) ([]User, error) {
	return transformUserRecords(records)
}
