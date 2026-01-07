package main

/*TO DO
Здесь, в функции main() нужно создать логгер, далее создать сервер с помощью вашей функции из пакета server, и запустить его.
Если при запуске сервера возникают ошибки, выведите её с помощью логгера на уровне Fatal.
*/

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func main() {

	// хз зачем, но на вебинаре делали
	converter := morse.NewConverter(morse.DefaultMorse)
	mainService := service.NewService(converter)

	// Создаем логгер
	logger := log.New(os.Stdout, "logger: ", log.LstdFlags)

	mainHandler := handlers.New(logger, mainService) //хз

	// Создаем сервер с помощью функции из пакета server
	srv := server.NewServer(logger, mainHandler)

	// Запускаем сервер
	err := srv.HTTPServer.ListenAndServe()
	if err != nil {
		// Если возникает ошибка при запуске сервера, выводим её с помощью логгера на уровне Fatal
		logger.Fatal("Error starting server: ", err)
	}
}
