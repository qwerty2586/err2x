package assertx

import "github.com/lainio/err2/assert"

type AsserterX assert.Asserter

var X = AsserterX(assert.AsserterToError)

func (x AsserterX) True(v bool, errx error) {
	if !v {
		panic(errx)
	}
}

func (x AsserterX) NoError(err error, errx error) {
	if err != nil {
		panic(errx)
	}
}
