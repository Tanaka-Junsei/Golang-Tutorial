package greetings

import (
	"fmt"

	"errors"
	"math/rand"
)

// Goでは、try-catchがないから、関数の返り値でerrorを返す
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name") // ここでerrorを新しく定義してる
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {

	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	format := []string{ // 挨拶の一覧を入れるスライス
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return format[rand.Intn(len(format))] // Intnで整数を返す
}
