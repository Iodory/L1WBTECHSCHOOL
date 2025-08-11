package main

var justString string

func createHugeString(num int) string {
	runes := make([]rune, num)
	for i := range runes {
		runes[i] = 'F'
	}
	return string(runes)
}

func someFunc() {
	v := createHugeString(1 << 10)

	// justString = v[:100] - это просто ссылка на строку v а новый способ это копирование нужных нам символов  и создание нового среза сборщик сам удалит старый срез
	justString = string([]rune(v[:100]))
}

func main() {
	someFunc()
}
