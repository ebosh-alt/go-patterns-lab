package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
)

// CSVFormat реализует изменяемые шаги импорта CSV.
type CSVFormat struct{}

// Parse разбирает заголовок и строки CSV.
func (f *CSVFormat) Parse(data []byte) ([]Record, error) {
	reader := csv.NewReader(bytes.NewReader(data))

	headers, err := reader.Read()
	if err != nil {
		if err == io.EOF {
			return nil, fmt.Errorf("CSV не содержит заголовок")
		}
		return nil, fmt.Errorf("прочитать заголовок CSV: %w", err)
	}

	var records []Record
	for {
		values, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("прочитать строку CSV: %w", err)
		}

		record := make(Record, len(headers))
		for index, header := range headers {
			record[header] = values[index]
		}
		records = append(records, record)
	}

	return records, nil
}

// Validate проверяет обязательные поля записей CSV.
func (f *CSVFormat) Validate(records []Record) error {
	return validateUserRecords(records)
}

// Transform преобразует записи CSV в пользователей.
func (f *CSVFormat) Transform(records []Record) ([]User, error) {
	return transformUserRecords(records)
}
