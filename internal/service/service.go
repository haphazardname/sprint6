package service

import (
	"os"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

/* TO DO
В этом пакете вы реализуете функцию автоматического определения кода Морзе или обычного текста из переданной строки.
Если передан обычный текст, функция должна переконвертировать его в код Морзе и вернуть;
и наоборот — если был передан код Морзе, функция должна переконвертировать его в обычный текст и вернуть.
Для реализации этой функции придётся обратиться к стандартной библиотеке, а именно — к пакету strings.
В этом пакете есть хорошие примеры, которые демонстрируют, как можно решить эту задачу.
Не забудьте обработать ошибки и вернуть их.
*/

//import "strings"

type Service struct {
	converter morse.Converter
}

func NewService(converter morse.Converter) *Service {
	return &Service{
		converter: converter,
	}
}

func (s *Service) ConvertString(input string) (string, error) {

	res := ""
	// Проверяем, содержит ли строка символы, которые не используются в коде Морзе
	if strings.ContainsAny(input, ".-") {
		// Если содержит, то считаем, что это код морзе
		res = s.converter.ToText(input)
	} else {
		// Иначе считаем, что это текст
		res = s.converter.ToMorse(strings.ToUpper(input))
		//res = input
	}
	//os.WriteFile(time.Now().String()+".txt", []byte(res), 0644)

	err := os.WriteFile(time.Now().Format("2006-01-02_15-04-05")+".txt", []byte(res), 0644)
	if err != nil {

		return "", err
	}
	return res, nil

}
