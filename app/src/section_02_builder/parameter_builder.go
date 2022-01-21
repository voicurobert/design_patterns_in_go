package builder

import (
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("email should contain @")
	}
	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.email.body = body
	return b
}

type build func(builder *EmailBuilder)

func SendEmail(action build) {
	builder := EmailBuilder{}

	action(&builder)

	sendEmailImpl(&builder.email)
}

func sendEmailImpl(email *email) {
	fmt.Println("EMAIL: ", email)
}

func ParameterExample() {
	SendEmail(func(b *EmailBuilder) {
		b.
			From("foo@bar.com").
			To("bar@baz.com").
			Subject("meeting").
			Body("hello, do you want to meet")
	})
}
