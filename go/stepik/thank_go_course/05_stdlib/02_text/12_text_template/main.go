package main

import (
	"bytes"
	"fmt"
	"text/template"
)

// начало решения

var templateText = `{{- .Name}}, добрый день! Ваш баланс - {{.Balance}}₽. {{- if (ge .Balance 100)}} Все в порядке. {{- else if (gt .Balance 0) }} Пора пополнить. {{- else }} Доступ заблокирован. {{- end}}`

// конец решения

type User struct {
	Name    string
	Balance int
}

// renderToString рендерит данные по шаблону в строку
func renderToString(tpl *template.Template, data any) string {
	var buf bytes.Buffer
	tpl.Execute(&buf, data)
	return buf.String()
}

func main() {
	tpl := template.New("message")
	tpl = template.Must(tpl.Parse(templateText))

	user := User{"Алиса", 500}
	got := renderToString(tpl, user)

	const want = "Алиса, добрый день! Ваш баланс - 500₽. Все в порядке."
	if got != want {
		fmt.Printf("%v: got '%v'\n", user, got)
	} else {
		fmt.Println("PASSED")
	}
}
