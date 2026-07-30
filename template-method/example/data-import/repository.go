package main

// User представляет пользователя после преобразования исходной записи.
type User struct {
	Name  string
	Email string
}

// Repository описывает сохранение импортированных пользователей.
type Repository interface {
	Save(users []User) error
}

// MemoryRepository сохраняет пользователей в памяти для демонстрации.
type MemoryRepository struct {
	users []User
}

// NewMemoryRepository создаёт пустое хранилище.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}

// Save сохраняет копию переданного списка пользователей.
func (r *MemoryRepository) Save(users []User) error {
	r.users = append([]User(nil), users...)
	return nil
}

// Users возвращает копию сохранённых пользователей.
func (r *MemoryRepository) Users() []User {
	return append([]User(nil), r.users...)
}
