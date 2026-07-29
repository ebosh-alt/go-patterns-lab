package main

import "fmt"

// Reception регистрирует пациента и служит первым звеном цепочки.
type Reception struct {
	BaseDepartment
}

// Execute выполняет регистрацию или пропускает её, если этап уже завершён.
func (r *Reception) Execute(patient *Patient) {
	if patient.RegistrationDone {
		fmt.Println("Reception: registration already completed")
		r.executeNext(patient)
		return
	}

	fmt.Printf("Reception: registering %s\n", patient.Name)
	patient.RegistrationDone = true
	r.executeNext(patient)
}
