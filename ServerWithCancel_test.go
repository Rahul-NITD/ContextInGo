package context_test

import "testing"

type StubStorage struct {
	response string
}

func (ss *StubStorage) Fetch() string {
	return ss.response
}

func TestContextCancelling(t *testing.T) {

}
