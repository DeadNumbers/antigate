package antigate

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

var a *antigate

func initAntigate() {
	if a != nil {
		return
	}

	antigate_key, err := readTextFile("key")
	if err != nil {
		panic("Error while reading key file")
	}

	a = New(antigate_key)
}

func TestProcessFromUrl(t *testing.T) {
	captcha_url := "https://bytebucket.org/poetofcode/antigate/raw/061c18a443b8a2af6ed400da3da1e7d28959f909/captcha.png"
	expected := "83tsU"

	initAntigate()
	captcha, err := a.ProcessFromUrl(captcha_url)

	if err != nil {
		t.Error(err)
	} else if captcha != expected {
		t.Error("Expected:", expected, "Got:", captcha)
	}
}

func TestProcessFromFile(t *testing.T) {
	expected := "83tsU"

	initAntigate()
	captcha, err := a.ProcessFromFile("captcha.png")

	if err != nil {
		t.Error(err)
	} else if captcha != expected {
		t.Error("Expected:", expected, "Got:", captcha)
	}
}

func TestGetBalance(t *testing.T) {
	initAntigate()

	balance, err := a.GetBalance()
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("Balance:", balance)
	}
}

func TestParseCaptchaId(t *testing.T) {
	wrongBody := "| 154209387"
	_, err := parseCaptchaId(wrongBody)
	if err == nil {
		t.Error("Error expected")
	}

	correctBody := "OK | 154209387"
	id, err := parseCaptchaId(correctBody)
	if err != nil || id != "154209387" {
		t.Error("Expected no errors and id = 154209387. Got id = ", id)
	}
}

func readTextFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return trimAll(string(content)), nil
}

func trimAll(in string) string {
	in = strings.Replace(in, " ", "", -1)
	in = strings.Replace(in, "\n", "", -1)
	return in
}
