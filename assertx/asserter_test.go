package assertx_test

import (
	"errors"
	"github.com/lainio/err2"
	"github.com/qwerty2586/err2x.git/assertx"
	"log"
	"testing"
)

type MyErr struct {
	Code     int
	Message  string
	Original error
}

func (m MyErr) Error() string {
	return m.Message
}

func (m MyErr) Handle() {
	log.Printf("Simulating handling error, sending code %d with body %s", m.Code, m.Message)
}

func handleErr(err error) {
	switch v := err.(type) {
	case MyErr:
		code := v.Code
		log.Println("code: ", code)
		v.Handle()
		break
	default:
		log.Println("default error, ", v.Error())
	}
}

func TestWithout(t *testing.T) {
	defer err2.Catch(handleErr)
	_, err := fReturningError()
	assertx.X.NoError(err, MyErr{Code: 500, Message: "Function returned error %w"})
	t.Error("should not be reached")
}

func fReturningError() (string, error) {
	return "", errors.New("error")
}
