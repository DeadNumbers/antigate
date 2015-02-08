package antigate

import (
	"fmt"
	"net/http"
)

func ProcessCaptchaByUrl(url string) (string, error) {
	fmt.Println(url)

	// var bytes []byte
	if _, err := loadImage(url); err != nil {
		return "", err
	}

	return "", nil
}

func loadImage(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil, err
	}
	defer resp.Body.Close()

	return nil, nil
}
