package antigate

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestProcessCaptchaByUrl(t *testing.T) {
	t.Error("Not implemented")
}

/*
func TestLoadImage(t *testing.T) {
	url := "https://bytebucket.org/poetofcode/antigate/raw/061c18a443b8a2af6ed400da3da1e7d28959f909/captcha.png"

	readed, err := loadImage(url)
	if err != nil {
		t.Error(err)
	}

	expected := getExpectedFileLength("captcha_base64.dat")

	if expected != readed {
		t.Error("Strings are different\n\nEXPECTED\n", expected, "\n\nGOT\n", readed)
	}
}
*/

func getExpectedFileLength(path string) string {
	content, _ := ioutil.ReadFile(path)
	return trimAll(string(content))
}

func trimAll(in string) string {
	in = strings.Replace(in, " ", "", -1)
	in = strings.Replace(in, "\n", "", -1)
	return in
}
