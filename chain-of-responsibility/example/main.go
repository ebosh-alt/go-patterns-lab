package main

func main() {
	// Клиент создаёт обработчики независимо друг от друга.
	reception := &Reception{}
	doctor := &Doctor{}
	medical := &Medical{}
	cashier := &Cashier{}

	// SetNext возвращает добавленное звено, поэтому цепочку можно собрать одной строкой.
	reception.SetNext(doctor).SetNext(medical).SetNext(cashier)

	patient := &Patient{Name: "Alex"}

	// Клиент знает только первое звено, остальные вызовутся последовательно.
	reception.Execute(patient)
}
