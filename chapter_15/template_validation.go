package main

import (
	"fmt"
	"text/template"
)

func main() {
	tOk := template.New("ok")
	template.Must(tOk.Parse("/* and a comment */ some static text:{{ .Name }}"))
	fmt.Println("The first one parsed OK.")
	fmt.Println("The next one ought to fail.")
	tErr := template.New("error_template")
	template.Must(tErr.Parse(" some static text {{ .Name}"))
}
