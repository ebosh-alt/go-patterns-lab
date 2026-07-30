package main

import "fmt"

// Record содержит данные одной записи в формате ключ-значение.
type Record map[string]string

// ImportFormat описывает шаги, которые зависят от формата источника.
type ImportFormat interface {
	// Parse разбирает исходные данные на отдельные записи.
	Parse(data []byte) ([]Record, error)

	// Validate проверяет обязательные поля записей.
	Validate(records []Record) error

	// Transform преобразует записи в доменные объекты.
	Transform(records []Record) ([]User, error)
}

// Importer содержит шаблонный метод импорта.
type Importer struct {
	format     ImportFormat
	repository Repository
}

// NewImporter создаёт импортёр для выбранного формата и хранилища.
func NewImporter(format ImportFormat, repository Repository) *Importer {
	return &Importer{
		format:     format,
		repository: repository,
	}
}

// Import выполняет шаги импорта в фиксированном порядке.
func (i *Importer) Import(data []byte) error {
	records, err := i.format.Parse(data)
	if err != nil {
		return err
	}

	if err := i.format.Validate(records); err != nil {
		return err
	}

	users, err := i.format.Transform(records)
	if err != nil {
		return err
	}

	return i.repository.Save(users)
}

// validateUserRecords проверяет поля, общие для поддерживаемых форматов.
func validateUserRecords(records []Record) error {
	for index, record := range records {
		if record["name"] == "" {
			return fmt.Errorf("запись %d: имя не заполнено", index+1)
		}
		if record["email"] == "" {
			return fmt.Errorf("запись %d: email не заполнен", index+1)
		}
	}

	return nil
}

// transformUserRecords преобразует записи в пользователей.
func transformUserRecords(records []Record) ([]User, error) {
	users := make([]User, 0, len(records))
	for _, record := range records {
		users = append(users, User{
			Name:  record["name"],
			Email: record["email"],
		})
	}

	return users, nil
}
