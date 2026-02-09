package main

import (
	"fmt"
	"study/http"
)

func main() {
	fmt.Println("Запуск Сервера")
	err := http.StartHTTPServer()
	if err != nil {
		fmt.Println("Во время работы произошла ошибка", err)
	} else {
		fmt.Println("Сервер запустился успешно! ")
	}

}
