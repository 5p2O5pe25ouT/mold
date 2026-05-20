package scrubbers

import (
	"context"
	"testing"

	. "github.com/go-playground/assert/v2"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//

func TestEmails(t *testing.T) {
	scrub := New()
	email := "Dean.Karn@gmail.com"

	type Test struct {
		Email string `scrub:"emails"`
	}

	tt := Test{Email: email}
	err := scrub.Struct(context.Background(), &tt)
	Equal(t, err, nil)
	Equal(t, tt.Email, "<<scrubbed::email::sha1::c52a47d4f3cde7c83046fc1fc8f208df74833cfa>>@gmail.com")

	err = scrub.Field(context.Background(), &email, "emails")
	Equal(t, err, nil)
	Equal(t, email, "<<scrubbed::email::sha1::c52a47d4f3cde7c83046fc1fc8f208df74833cfa>>@gmail.com")

	var iface interface{}
	err = scrub.Field(context.Background(), &iface, "emails")
	Equal(t, err, nil)
	Equal(t, iface, nil)

	iface = "Dean.Karn@gmail.com"
	err = scrub.Field(context.Background(), &iface, "emails")
	Equal(t, err, nil)
	Equal(t, iface, "<<scrubbed::email::sha1::c52a47d4f3cde7c83046fc1fc8f208df74833cfa>>@gmail.com")

	emailText := "alice@example.com bob@example.com"
	err = scrub.Field(context.Background(), &emailText, "emails")
	Equal(t, err, nil)
	Equal(t, emailText,
		"<<scrubbed::email::sha1::522b276a356bdf39013dfabea2cd43e141ecc9e8>>@example.com "+
			"<<scrubbed::email::sha1::48181acd22b3edaebc8a447868a7df7ce629920a>>@example.com",
	)
}

func TestText(t *testing.T) {
	scrub := New()
	name := "Joey Bloggs"

	type Test struct {
		String string `scrub:"text"`
	}

	tt := Test{String: name}
	err := scrub.Struct(context.Background(), &tt)
	Equal(t, err, nil)
	Equal(t, tt.String, "<<scrubbed::text::sha1::028f74c1850aa1efb33a2e8746c0f4183e1e8e30>>")

	err = scrub.Field(context.Background(), &name, "text")
	Equal(t, err, nil)
	Equal(t, name, "<<scrubbed::text::sha1::028f74c1850aa1efb33a2e8746c0f4183e1e8e30>>")

	var iface interface{}
	err = scrub.Field(context.Background(), &iface, "text")
	Equal(t, err, nil)
	Equal(t, iface, nil)

	iface = "Joey Bloggs"
	err = scrub.Field(context.Background(), &iface, "text")
	Equal(t, err, nil)
	Equal(t, iface, "<<scrubbed::text::sha1::028f74c1850aa1efb33a2e8746c0f4183e1e8e30>>")

	// testing Text wrapped func
	name = "Joey Bloggs"
	err = scrub.Field(context.Background(), &name, "name")
	Equal(t, err, nil)
	Equal(t, name, "<<scrubbed::name::sha1::028f74c1850aa1efb33a2e8746c0f4183e1e8e30>>")
}
