package antigate

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type antigate struct {
	key string
}

func New(key string) *antigate {
	return &antigate{key}
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

	captcha, err := checkCaptcha(captcha_id, a.key)
	if err != nil {
		return "", err
	}

	return captcha, nil
}

func (a *antigate) GetBalance() (float64, error) {
	var url = "http://antigate.com/res.php?key=" + a.key + "&action=getbalance"

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	balance, err := strconv.ParseFloat(string(body), 64)
	if err != nil {
		return 0, err
	}

	return balance, nil
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
	if err != nil {
		return "", err
	}

	id, err := parseCaptchaId(string(body))
	if err != nil {
		return "", err
	}

	return id, nil
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

func checkCaptcha(id, key string) (string, error) {
	url := "http://antigate.com/res.php?key=" + key + "&action=get&id=" + id

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// parsing reponse
	response := string(body)

	if strings.Contains(response, "OK") {
		list := strings.Split(response, "|")
		for i := range list {
			list[i] = strings.TrimSpace(list[i])
		}

		return list[1], nil
	}

	if response == "CAPCHA_NOT_READY" {
		time.Sleep(5 * time.Second)
		return checkCaptcha(id, key)
	}

	return "", errors.New("Error while receiving captcha")
}
