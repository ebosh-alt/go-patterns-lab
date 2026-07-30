package main

import "fmt"

func main() {
	csvData := []byte("name,email\nАнна,anna@example.com\nИван,ivan@example.com\n")
	runImport("CSV", &CSVFormat{}, csvData)

	fmt.Println()

	jsonData := []byte(`[
		{"name":"Мария","email":"maria@example.com"},
		{"name":"Пётр","email":"petr@example.com"}
	]`)
	runImport("JSON", &JSONFormat{}, jsonData)
}

// runImport запускает один и тот же шаблонный метод для выбранного формата.
func runImport(name string, format ImportFormat, data []byte) {
	repository := NewMemoryRepository()
	importer := NewImporter(format, repository)

	if err := importer.Import(data); err != nil {
		fmt.Printf("%s: ошибка импорта: %v\n", name, err)
		return
	}

	fmt.Printf("%s: импортировано пользователей: %d\n", name, len(repository.Users()))
	for _, user := range repository.Users() {
		fmt.Printf("- %s <%s>\n", user.Name, user.Email)
	}
}
