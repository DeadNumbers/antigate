# Antigate API (Golang)

API на языке Go для сервиса распознавания капч [Antigate](anti-captcha.com)

## Пример использования

```
#!go

package main

import (
	"fmt"
	"bitbucket.org/poetofcode/antigate" 
)

func main() {
	a := antigate.New("YOUR_ANTIGATE_API_KEY")

	// From URL
	captcha_text, _ := a.ProcessFromUrl("https://bytebucket.org/poetofcode/antigate/raw/061c18a443b8a2af6ed400da3da1e7d28959f909/captcha.png")
	fmt.Println("from url:", captcha_text)

	// From file
	captcha_text, _ = a.ProcessFromFile("captcha.png")	// the file "captcha.png" must exist
	fmt.Println("from file:", captcha_text)

	// Your balance
	balance, _ := a.GetBalance()
	fmt.Println("balance:", balance)
}

```

Сохраните этот код в файл под именем **main.go**, вставив свой Antigate-ключ вместо YOUR_ANTIGATE_API_KEY и выполните в терминале:
```
go get
go run main.go
```

## Запуск тестов

Клонируйте репозиторий:

```
hg clone https://bitbucket.org/poetofcode/antigate
cd antigate
```
Создайте в каталоге репозиторий файл под именем **key** (без расширения) и сохраните в нём свой ключ от Antigate

Для запуска тестов выполните:
```
go test
```