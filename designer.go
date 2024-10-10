package main

type Designer struct {
	ID        int
	WorkTime  int // Время работы в днях
	Busy      bool // Занят или свободен
}

func (d *Designer) Work() {
	// Логика работы проектировщика
	d.Busy = true
	// Здесь можно добавить задержку, чтобы симулировать время работы
}
