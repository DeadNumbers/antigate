package antigate

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	fmt.Println(string(body))

	return "", nil
}
