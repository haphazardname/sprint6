package server

/* TO DO
В этом пакете вы реализуете функцию для создания http-сервера.
Алгоритм реализации:
Создайте структуру сервера с полями для логгера (log.Logger) и http-сервера (http.Server).
Создайте функцию, в которой нужно создать http-роутер. Функция принимает log.Logger и возвращает экземпляр структуры вашего сервера.
Зарегистрируйте ваши хендлеры в http-роутере.
Создайте экземпляр структуры http.Server. Для настройки вашего сервера используйте следующие поля:
Addr — используйте порт 8080.
Handler — передайте ваш http-роутер.
ErrorLog — передайте ваш логгер.
ReadTimeout — таймаут для чтения. 5 секунд.
WriteTimeout — таймаут для записи. 10 секунд.
IdleTimeout — таймаут ожидания следующего запроса. 15 секунд.
Верните ссылку на ваш сервер.
*/

import (
	"log"
	"net/http"
	"time"
)

// Структура сервера
type Server struct {
	Logger     *log.Logger
	HTTPServer http.Server
}

// хз, видимо для регистрации
type Handler interface {
	HandlerRoot(w http.ResponseWriter, r *http.Request)
	HandlerUpload(w http.ResponseWriter, r *http.Request)
}

// Функция для создания HTTP-роутера
func NewServer(logger *log.Logger, hand Handler) *Server {

	// Создаем новый роутер
	mux := http.NewServeMux()

	// Регистрируем хендлеры в роутере
	mux.HandleFunc("/", hand.HandlerRoot)
	mux.HandleFunc("/upload", hand.HandlerUpload)

	// Создаем экземпляр структуры http.Server.
	server := &Server{
		Logger: logger,
		HTTPServer: http.Server{
			Addr:         ":8080",
			Handler:      mux,
			ErrorLog:     logger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}

	// Возвращаем ссылку на наш сервер
	return server
}
