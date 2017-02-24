package errors

import . "github.com/coderlindacheng/balabalago/special_string"

type ErrorWrapper struct {
	Err    error
	Attach string
}

func (p *ErrorWrapper) Error() string {
	if p.Err != nil {
		return p.Err.Error() + UNIX_LINE_SEPARATOR + p.Attach
	} else {
		return p.Attach
	}
}
