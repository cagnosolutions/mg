package mg

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var DOMAIN = "api.mailgun.net/v3/sandbox73d66ccb60f948708fcaf2e2d1b3cd4c.mailgun.org"
var KEY = "key-173701b40541299bd3b7d40c3ac6fd43"
var ENDPOINT = "https://api:" + KEY + "@" + DOMAIN + "/messages"

type Email struct {
	To      []string
	From    string
	CC      []string
	BCC     []string
	Subject string
	Text    string
	HTML    string
}

func Send(To string, From string, Subject string, HTML string, BCC ...string) (string, error) {
	vals := make(url.Values, 0)
	vals.Add("to", To)
	vals.Add("from", From)
	vals.Add("subject", Subject)
	vals.Add("html", HTML)
	vals["bcc"] = BCC

	resp, err := http.PostForm(ENDPOINT, vals)
	if err != nil {
		log.Printf("mailgun.go >> Send() >> http.PostForm() >> %v\n", err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("mailgun.go >> Send() >> ioutil.ReadlAll() >> %v\n", err)
		return "", err
	}

	return fmt.Sprintf("%s", body), nil
}

func SendEmail(email Email) (string, error) {

	vals := make(url.Values, 0)

	vals["to"] = email.To
	vals["cc"] = email.CC
	vals["bcc"] = email.BCC
	vals.Add("from", email.From)
	vals.Add("subject", email.Subject)
	vals.Add("text", email.Text)
	vals.Add("html", email.HTML)

	resp, err := http.PostForm(ENDPOINT, vals)
	if err != nil {
		log.Printf("mailgun.go >> Send() >> http.PostForm() >> %v\n", err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("mailgun.go >> Send() >> ioutil.ReadlAll() >> %v\n", err)
		return "", err
	}

	return fmt.Sprintf("%s", body), nil
}
