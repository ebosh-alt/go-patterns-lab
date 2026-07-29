package main

// Patient представляет запрос, который проходит через цепочку отделений.
// Флаги показывают, какие этапы обработки уже завершены.
type Patient struct {
	Name              string
	RegistrationDone  bool
	DoctorCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}
