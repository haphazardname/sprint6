package handlers

/* TO DO
В этом пакете вы реализуете два хендлера.
Для корневого эндпоинта / нужно реализовать хендлер, который возвращает HTML из файла index.html.
Второй хендлер для эндпоинта /upload должен выполнять следующие действия:
Парсить html-форму из файла index.html.
Получить файл из формы (не забудьте его закрыть).
Прочитать данные из файла.
Передать эти данные в функцию автоопределения из пакета service, которую вы создали, чтобы получить переконвертируемую строку.
Создать локальный файл. Эта операция обычно небезопасна и так делать не рекомендуется, но в рамках нашего задания хотелось бы более наглядного результата, поэтому мы решились на этот шаг, ради видимого результата. А вообще, обычно используют временные файлы.
Записать в локальный файл результат конвертации строки. Для генерации имени файла вы можете использовать время с помощью time.Now().UTC().String(). Чтобы получить расширения файла, используйте filepath.Ext().
Вернуть результат конвертации строки.
Там, где это необходимо, обработайте возможные ошибки. Статус при возникновении ошибок http.StatusInternalServerError.
*/

import (
	"io"
	"log"
	"net/http"
)

type service interface {
	ConvertString(input string) (string, error)
}

type Handler struct {
	logger  *log.Logger
	service service
}

// хз зачем
func New(logger *log.Logger, service service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

// для эндпоинта / - возвращает HTML из файла index.html
func (h *Handler) HandlerRoot(w http.ResponseWriter, r *http.Request) {
	filepath := "index.html"
	http.ServeFile(w, r, filepath)
}

// для эндпоинта /upload
func (h *Handler) HandlerUpload(w http.ResponseWriter, r *http.Request) {

	//Парсить html-форму из файла index.html
	if err := r.ParseMultipartForm(10000000000); err != nil {
		h.logger.Printf("Error parsing form: %v", err)
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	//Получить файл из формы (не забудьте его закрыть).
	file, _, err := r.FormFile("myFile")
	if err != nil {
		h.logger.Printf("Error getting file: %v", err)
		http.Error(w, "Error getting file", http.StatusInternalServerError)
		return
	}

	//Прочитать данные из файла.
	input, err := io.ReadAll(file)
	if err != nil {
		h.logger.Printf("Error reading file: %v", err)
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}
	//закрываем файл
	defer file.Close()

	//Передать эти данные в функцию автоопределения из пакета service, которую вы создали, чтобы получить переконвертируемую строку.
	result, err := h.service.ConvertString(string(input))
	if err != nil {
		h.logger.Printf("Error convert string: %v", err)
		http.Error(w, "Error convert string", http.StatusInternalServerError)
		return
	}

	//Создать локальный файл. Эта операция обычно небезопасна и так делать не рекомендуется, но в рамках нашего задания хотелось бы более наглядного результата,
	// поэтому мы решились на этот шаг, ради видимого результата. А вообще, обычно используют временные файлы.
	//Записать в локальный файл результат конвертации строки. Для генерации имени файла вы можете использовать время с помощью time.Now().UTC().String().
	// Чтобы получить расширения файла, используйте filepath.Ext().
	// - сделано внутри service.ConvertString

	//Вернуть результат конвертации строки.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
