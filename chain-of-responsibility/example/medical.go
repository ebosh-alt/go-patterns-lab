package main

import "fmt"

// Medical выдаёт пациенту лекарства после осмотра.
type Medical struct {
	BaseDepartment
}

// Execute выдаёт лекарства или пропускает этап, если он уже завершён.
func (m *Medical) Execute(patient *Patient) {
	if patient.MedicineDone {
		fmt.Println("Medical: medicine already given")
		m.executeNext(patient)
		return
	}

	fmt.Printf("Medical: giving medicine to %s\n", patient.Name)
	patient.MedicineDone = true
	m.executeNext(patient)
}
