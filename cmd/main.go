package main

import "todo-list/internal/app"

// Хэнлдер идет в сервис - (валидация в сервисе) - сервис идет в бд

func main() {
	app.Start()
}
