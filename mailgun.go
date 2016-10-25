package mg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var DOMAIN = ""
var KEY = ""
var ENDPOINT = "https://api:" + KEY + "@" + DOMAIN

var API = errors.New("No KEY or DOMAIN is set")
var TO = errors.New("No email addres is set")

type Email struct {
	To      []string
	From    string
	CC      []string
	BCC     []string
	Subject string
	Text    string
	HTML    string
	Tags    []string
}

func Send(to string, from string, subject string, html string, tags []string, bcc ...string) (string, error) {
	if KEY == "" || DOMAIN == "" {
		return "", API
	}
	if to == "" {
		return "", TO
	}
	vals := make(url.Values, 0)
	vals.Add("to", to)
	vals.Add("from", from)
	vals.Add("subject", subject)
	vals.Add("html", html)
	vals["bcc"] = bcc
	vals["o:tag"] = tags

	resp, err := http.PostForm(ENDPOINT+"/messages", vals)
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
	if KEY == "" || DOMAIN == "" {
		return "", API
	}
	if len(email.To) < 1 {
		return "", TO
	}

	vals := make(url.Values, 0)
	vals["to"] = email.To
	vals["cc"] = email.CC
	vals["bcc"] = email.BCC
	vals["o:tag"] = email.Tags
	vals.Add("from", email.From)
	vals.Add("subject", email.Subject)
	vals.Add("text", email.Text)
	vals.Add("html", email.HTML)

	// fmt.Println(vals.Get("o:tag"))

	resp, err := http.PostForm(ENDPOINT+"/messages", vals)
	if err != nil {
		log.Printf("mailgun.go >> SendEmail() >> http.PostForm() >> %v\n", err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("mailgun.go >> SendEmail() >> ioutil.ReadlAll() >> %v\n", err)
		return "", err
	}

	return fmt.Sprintf("%s", body), nil
}

func SetCredentials(domain, key string) {
	if domain != "" {
		DOMAIN = domain
	}
	if key != "" {
		KEY = key
	}
	ENDPOINT = "https://api:" + KEY + "@" + DOMAIN
}

func GetTag(tag string) (string, error) {

	if KEY == "" || DOMAIN == "" {
		return "", API
	}

	resp, err := http.Get(ENDPOINT + "/tags/" + tag)
	if err != nil {
		log.Printf("mailgun.go >> GetTag() >> http.Get() >> %v\n", err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("mailgun.go >> GetTag() >> ioutil.ReadlAll() >> %v\n", err)
		return "", err
	}
	return string(body), nil
}
