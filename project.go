package main

import "math/rand"

type Project struct {
	Urgency bool
}

func NewProject() Project {
	// Логика создания проекта с возможностью срочности
	return Project{Urgency: rand.Intn(2) == 1} // 50% шанс на срочный проект
}
