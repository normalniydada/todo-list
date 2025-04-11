package main

import "todo-list/internal/app"

// Хэнлдер идет в сервис - (валидация в сервисе) - сервис идет в бд
// GET - Redis (если в редисе не найдено, то идем в БД, предварительно сохранив данные в Redis) -> БД
func main() {
	app.Start()
}
