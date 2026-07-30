package main

import (
	"errors"
	"reflect"
	"testing"
)

// formatSpy записывает вызовы изменяемых шагов импорта.
type formatSpy struct {
	steps       *[]string
	validateErr error
}

func (s *formatSpy) Parse(data []byte) ([]Record, error) {
	*s.steps = append(*s.steps, "parse")
	return []Record{{"name": "Анна", "email": "anna@example.com"}}, nil
}

func (s *formatSpy) Validate(records []Record) error {
	*s.steps = append(*s.steps, "validate")
	return s.validateErr
}

func (s *formatSpy) Transform(records []Record) ([]User, error) {
	*s.steps = append(*s.steps, "transform")
	return []User{{Name: "Анна", Email: "anna@example.com"}}, nil
}

// repositorySpy записывает вызов финального шага сохранения.
type repositorySpy struct {
	steps *[]string
}

func (s *repositorySpy) Save(users []User) error {
	*s.steps = append(*s.steps, "save")
	return nil
}

// TestImporterPreservesStepOrder проверяет неизменный порядок шагов импорта.
func TestImporterPreservesStepOrder(t *testing.T) {
	var steps []string
	importer := NewImporter(
		&formatSpy{steps: &steps},
		&repositorySpy{steps: &steps},
	)

	if err := importer.Import([]byte("source data")); err != nil {
		t.Fatalf("Import returned an error: %v", err)
	}

	want := []string{"parse", "validate", "transform", "save"}
	if !reflect.DeepEqual(steps, want) {
		t.Fatalf("steps = %v, want %v", steps, want)
	}
}

// TestImporterStopsAfterError проверяет остановку после ошибки текущего шага.
func TestImporterStopsAfterError(t *testing.T) {
	wantErr := errors.New("invalid records")
	var steps []string
	importer := NewImporter(
		&formatSpy{steps: &steps, validateErr: wantErr},
		&repositorySpy{steps: &steps},
	)

	err := importer.Import([]byte("source data"))

	if !errors.Is(err, wantErr) {
		t.Fatalf("error = %v, want %v", err, wantErr)
	}

	wantSteps := []string{"parse", "validate"}
	if !reflect.DeepEqual(steps, wantSteps) {
		t.Fatalf("steps = %v, want %v", steps, wantSteps)
	}
}

// TestCSVImport проверяет импорт пользователей из CSV.
func TestCSVImport(t *testing.T) {
	data := []byte("name,email\nАнна,anna@example.com\nИван,ivan@example.com\n")
	repository := NewMemoryRepository()
	importer := NewImporter(&CSVFormat{}, repository)

	if err := importer.Import(data); err != nil {
		t.Fatalf("Import returned an error: %v", err)
	}

	want := []User{
		{Name: "Анна", Email: "anna@example.com"},
		{Name: "Иван", Email: "ivan@example.com"},
	}
	if users := repository.Users(); !reflect.DeepEqual(users, want) {
		t.Fatalf("users = %#v, want %#v", users, want)
	}
}

// TestJSONImport проверяет импорт пользователей из JSON.
func TestJSONImport(t *testing.T) {
	data := []byte(`[
		{"name":"Анна","email":"anna@example.com"},
		{"name":"Иван","email":"ivan@example.com"}
	]`)
	repository := NewMemoryRepository()
	importer := NewImporter(&JSONFormat{}, repository)

	if err := importer.Import(data); err != nil {
		t.Fatalf("Import returned an error: %v", err)
	}

	want := []User{
		{Name: "Анна", Email: "anna@example.com"},
		{Name: "Иван", Email: "ivan@example.com"},
	}
	if users := repository.Users(); !reflect.DeepEqual(users, want) {
		t.Fatalf("users = %#v, want %#v", users, want)
	}
}
