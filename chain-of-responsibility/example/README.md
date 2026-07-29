# Реализация на Go

В Go паттерн удобно собрать через интерфейс и композицию:

- `Department` описывает выполнение шага и установку следующего отделения;
- `BaseDepartment` хранит ссылку на следующее звено и безопасно передаёт ему пациента;
- конкретные отделения встраивают `BaseDepartment`, выполняют свою операцию и продолжают цепочку;
- `main` создаёт отделения и задаёт их порядок.

Пример моделирует путь пациента по больнице:

```text
Reception -> Doctor -> Medical -> Cashier
```

Каждое отделение изменяет только свою часть состояния пациента. Последнее звено завершает обработку без специального терминатора.

## Структура примера

- [`department.go`](example/department.go) содержит интерфейс и базовый обработчик;
- [`patient.go`](example/patient.go) описывает состояние пациента;
- [`reception.go`](example/resecption.go), [`doctor.go`](example/doctor.go), [`medical.go`](example/medical.go) и [`cashier.go`](example/cashier.go) содержат конкретные звенья;
- [`main.go`](example/main.go) собирает и запускает цепочку;
- [`chain_test.go`](example/chain_test.go) проверяет полный маршрут и безопасное завершение.

## Запуск
Из корня репозитория:
```bash
go run ./chain-of-responsibility/example
go test ./...
```
