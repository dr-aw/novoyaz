package main

import (
	"fmt"
	"os"
	"strings"
)

var keys = []string{"ник", "щик", "чик", "ова", "ева", "ива", "енн", "ка", "чив", "ущ", "ющ", "яч", "оват", "оров", "дны", "дно", "лог", "чес", "щест", "еч", "сно", "сот", "изм", "етн", "сти"}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Пожалуйста, укажите словесы.")
		return
	}
	fmt.Println("______________________")

	//Проверка каждого слова на соответствие
	for _, word := range os.Args[1:] {
		found := false
		for _, key := range keys {
			if strings.Contains(word, key) {
				parts := strings.SplitN(word, key, 2)
				part1 := parts[0]
				if len(parts[0]) < 3 {
					part1 = changer(strings.ToLower(parts[0] + key))
				} else {
					part1 = changer(strings.ToLower(parts[0])) + key
				}
				part2 := parts[1]
				newWord := part1 + part2
				fmt.Printf("%s ", newWord)
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("%s ", word)
		}
	}
	fmt.Println() // Перенос строки после вывода всех слов
}

func changer(part string) string {
	length := len(part)
	//if length == 0 {
	//	fmt.Fprint(os.Stderr, "Ошибка - строка пустая\n")
	//	os.Exit(1)
	//}

	runes := []rune(part)
	lastChar := runes[len(runes)-1]
	var secondLastChar rune
	if length > 1 {
		secondLastChar = runes[len(runes)-2]
	}
	switch lastChar {
	case 'а', 'я':
		return "хуя"
	case 'е', 'э':
		return "хуе"
	case 'ё', 'о':
		return "хуё"
	case 'и', 'ы':
		return "хуи"
	case 'й':
		return "хуй"
	case 'у', 'ю':
		return "хую"
	}

	switch secondLastChar {
	case 'а', 'я':
		return "хуя" + string(lastChar)
	case 'е', 'э':
		return "хуе" + string(lastChar)
	case 'ё', 'о':
		return "хуё" + string(lastChar)
	case 'и', 'ы':
		return "хуи" + string(lastChar)
	case 'й':
		return "хуй" + string(lastChar)
	case 'у', 'ю':
		return "хую" + string(lastChar)
	}

	return part
}
