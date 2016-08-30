package mg

import (
	"bytes"
	"html/template"
	"log"
)

func Body(file string, val interface{}) (string, error) {
	t, err := template.ParseFiles(file)
	if err != nil {
		log.Printf("package mailgun >> template.go >> Body() >> template.ParseFiles() >> %v\n\n", err)
		return "", err
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, val); err != nil {
		log.Printf("package mailgun >> template.go >> Body() >> t.Execute() >> %v\n\n", err)
		return "", err
	}
	return buf.String(), nil
}
