package errors

import (
	. "github.com/coderlindacheng/balabalago/special_string"
)

type errorString struct {
	string
}

func (e errorString) Error() string {
	return e.string
}

type errorWrapper struct {
	string
	error
}

func (p *errorWrapper) Error() string {
	if p.error != nil {
		return p.error.Error() + UNIX_LINE_SEPARATOR + p.string
	} else {
		return p.string
	}
}

func NewOnlyStr(s string) error {
	return errorString{s}
}

func NewWrapper(err error, s string) error {
	return &errorWrapper{s, err}
}
