package mg

import (
	"bytes"
	"html/template"
	"log"
)

func BodyFile(file string, val interface{}) (string, error) {
	t, err := template.ParseFiles(file)
	if err != nil {
		log.Printf("package mailgun >> template.go >> BodyFile() >> template.ParseFiles() >> %v\n\n", err)
		return "", err
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, val); err != nil {
		log.Printf("package mailgun >> template.go >> BodyFile() >> t.Execute() >> %v\n\n", err)
		return "", err
	}
	return buf.String(), nil
}

func Body(body string, val interface{}) (string, error) {
	t, err := template.New("body").Parse(body)
	if err != nil {
		log.Printf("package mailgun >> template.go >> Body() >> template.New(\"body\").Parse(body) >> %v\n\n", err)
		return "", err
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, val); err != nil {
		log.Printf("package mailgun >> template.go >> Body() >> t.Execute() >> %v\n\n", err)
		return "", err
	}
	return buf.String(), nil
}
