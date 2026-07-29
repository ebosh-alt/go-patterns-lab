package main

import "fmt"

// Cashier принимает оплату и завершает маршрут пациента.
type Cashier struct {
	BaseDepartment
}

// Execute принимает оплату или пропускает этап, если он уже завершён.
func (c *Cashier) Execute(patient *Patient) {
	if patient.PaymentDone {
		fmt.Println("Cashier: payment already completed")
		c.executeNext(patient)
		return
	}

	fmt.Printf("Cashier: accepting payment from %s\n", patient.Name)
	patient.PaymentDone = true
	c.executeNext(patient)
}
