package main

import "testing"

// TestChainProcessesPatient проверяет прохождение пациента через всю цепочку.
func TestChainProcessesPatient(t *testing.T) {
	reception := &Reception{}
	doctor := &Doctor{}
	medical := &Medical{}
	cashier := &Cashier{}

	reception.SetNext(doctor).SetNext(medical).SetNext(cashier)

	patient := &Patient{Name: "Alex"}
	reception.Execute(patient)

	if !patient.RegistrationDone {
		t.Error("registration was not completed")
	}
	if !patient.DoctorCheckUpDone {
		t.Error("doctor checkup was not completed")
	}
	if !patient.MedicineDone {
		t.Error("medicine was not given")
	}
	if !patient.PaymentDone {
		t.Error("payment was not completed")
	}
}

// TestDepartmentWithoutNextStopsSafely проверяет безопасное завершение последнего звена.
func TestDepartmentWithoutNextStopsSafely(t *testing.T) {
	patient := &Patient{Name: "Alex"}

	reception := &Reception{}
	reception.Execute(patient)

	if !patient.RegistrationDone {
		t.Fatal("reception did not process patient")
	}
}
