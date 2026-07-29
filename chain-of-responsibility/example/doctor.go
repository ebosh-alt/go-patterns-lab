package main

import "fmt"

// Doctor проводит осмотр пациента после регистрации.
type Doctor struct {
	BaseDepartment
}

// Execute выполняет осмотр или пропускает его, если этап уже завершён.
func (d *Doctor) Execute(patient *Patient) {
	if patient.DoctorCheckUpDone {
		fmt.Println("Doctor: checkup already completed")
		d.executeNext(patient)
		return
	}

	fmt.Printf("Doctor: checking %s\n", patient.Name)
	patient.DoctorCheckUpDone = true
	d.executeNext(patient)
}
