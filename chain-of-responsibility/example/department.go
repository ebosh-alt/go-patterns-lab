package main

// Department задаёт общий интерфейс для всех звеньев цепочки.
type Department interface {
	// Execute обрабатывает пациента и при необходимости передаёт его дальше.
	Execute(*Patient)

	// SetNext устанавливает следующее звено и возвращает его для удобной сборки цепочки.
	SetNext(Department) Department
}

// BaseDepartment хранит общую для всех отделений ссылку на следующее звено.
type BaseDepartment struct {
	next Department
}

// SetNext связывает текущий обработчик со следующим.
func (d *BaseDepartment) SetNext(next Department) Department {
	d.next = next
	return next
}

// executeNext передаёт пациента дальше, если цепочка ещё не закончилась.
func (d *BaseDepartment) executeNext(patient *Patient) {
	if d.next != nil {
		d.next.Execute(patient)
	}
}
