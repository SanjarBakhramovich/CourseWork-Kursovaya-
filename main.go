package main

import (
	"fmt"
	"time"
)

func main() {
	var numProjects int
	fmt.Print("Введите количество проектов для обработки: ")
	_, err := fmt.Scan(&numProjects)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	designers := []Designer{
		{ID: 1, WorkTime: 4},
		{ID: 2, WorkTime: 6},
		{ID: 3, WorkTime: 5},
		{ID: 4, WorkTime: 7},
	}

	for i := 1; i <= numProjects; i++ { // Обработка проектов от 1 до numProjects
		project := NewProject()
		if project.Urgency {
			// Логика для обработки срочного проекта
			fmt.Printf("Проект %d/%d: Срочный проект\n", i, numProjects)
			for _, designer := range designers {
				designer.Work()
				time.Sleep(time.Duration(designer.WorkTime) * time.Second)
				designer.Busy = false // Освобождаем проектировщика
			}
		} else {
			// Логика для обычного проекта
			fmt.Printf("Проект %d/%d: Обычный проект\n", i, numProjects)
			for _, designer := range designers {
				designer.Work()
				time.Sleep(time.Duration(designer.WorkTime) * time.Second)
				designer.Busy = false
			}
		}
	}
	fmt.Println("Все проекты обработаны!")
}
