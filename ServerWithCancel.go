package context

import "net/http"

type Storage interface {
	Fetch() string
}

func Server(s Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Fetch()
	}
}
