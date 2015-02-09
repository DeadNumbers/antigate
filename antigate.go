package antigate

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	// "net/url"
	// "strconv"
	"strings"
)

type antigate struct {
	key string
}

func New(key string) antigate {
	return antigate{key}
}

func (a *antigate) ProcessCaptchaByUrl(url string) (string, error) {
	image, err := loadImage(url)
	if err != nil {
		return "", err
	}

	captcha_id, err := uploadCaptcha(image, a.key)
	if err != nil {
		return "", err
	}

	return captcha_id, nil
}

func loadImage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	str := base64.StdEncoding.EncodeToString(body)
	return str, nil
}

func uploadCaptcha(imageBody string, key string) (string, error) {

	fmt.Println("Key:", key)

	/*
		resp, err := http.PostForm(
			"http://antigate.com/in.php",
			url.Values{
				"method": {"base64"},
				"key":    {key},
				"body":   {imageBody},
			})

		if err != nil {
			return "", err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		// fmt.Println(string(body))
		strBody := string(body)
	*/

	// strBody := "OK | 154209387"

	// if strings.Contains(strBody, "OK") == false {
	// 	return "", errors.New("Failed to get captcha-id")
	// }

	// id, err := strconv.Atoi(strBody)
	// if err != nil {
	// 	return "", err
	// }

	// fmt.Println(id)

	return "", nil
}

func parseCaptchaId(str string) (string, error) {
	list := strings.Split(str, "|")
	for i := range list {
		list[i] = strings.TrimSpace(list[i])
	}

	if list[0] != "OK" {
		return "", errors.New("Unable to get a captcha id")
	}

	return list[1], nil
}
