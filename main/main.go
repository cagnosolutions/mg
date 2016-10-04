package main

import (
	"fmt"
	"log"

	"github.com/cagnosolutions/mg"
)

type User struct {
	Name   string
	Age    int
	Active bool
}

var DOMAIN = "api.mailgun.net/v3/sandbox73d66ccb60f948708fcaf2e2d1b3cd4c.mailgun.org"
var KEY = "key-173701b40541299bd3b7d40c3ac6fd43"

func main() {

	mg.SetCredentials(DOMAIN, KEY)

	// user := User{
	// 	Name:   "Greg Pechiro",
	// 	Age:    30,
	// 	Active: true,
	// }
	//
	// //body, err := mg.BodyFile("email.tmpl", map[string]interface{}{"user": user})
	// body, err := mg.Body(email, map[string]interface{}{"user": user})
	//
	// if err != nil {
	// 	panic(err)
	// }
	// //r, err := mg.Send("gregpechiro@gmail.com", "Totally not a virus or spam <info@test.com>", "TEST EMAIL", body, "gregpechiro@yahoo.com", "scottiecagno@gmaio.com", "cagnosolutions@gmail.com")
	// email := mg.Email{
	// 	To:      []string{"gregpechiro@gmail.com"},
	// 	From:    "Not a virus or spam <info@test.com>",
	// 	Subject: "Totally not a virus or spam... I promise",
	// 	HTML:    body,
	// 	//CC:      []string{"gregpechiro@yahoo.com", "cagnosolutions@gmail.com"},
	// 	//BCC:     []string{"scottiecagno@gmail.com"},
	// 	Tags: []string{"company-1234"},
	// }
	// r, err := mg.SendEmail(email)

	r, err := mg.GetTag("company:1234")

	if err == mg.API {
		log.Panic("Please set API domain and key")
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", r)

}

var email = `<!doctype html>
<html lang="en-US" style="font-family:'Nato Sans',sans-serif;font-size: 14px; color:#777;">
	<head>
		<meta charset="utf-8">
	</head>
	<body>
		<div style="margin: 25px;">
			<div style="width:600px;margin:0px auto;background-color:#f5f6f5;border:1px solid #dddddd;-moz-border-radius:3px;-webkit-border-radius:3px;">
				<div style="padding-left: 27px;padding-right: 27px;padding-bottom: 27px;">
					<div style="border:0px solid #999;margin-top:25px;padding:15px;">
						<h3>
							HELLO!
						</h3>
						<hr/>
						<p>My name is {{ .user.Name }}</p>
						<p>
							I am {{ .user.Age }}. This email was composed using the template wrapper I just created
						</p>
						<p>It was sent using the mailgun api wrapper I just wrote.</p>
					</div>
					<p style="font-weight:100;font-size:11px;text-align:center;">
						<em>Peace</em>
					</p>
				</div>
			</div>
		</div>
	</body>
</html>
`
