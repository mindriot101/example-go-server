package server

import "net/http"

func (s *server) AddRoutes() {
	s.router.HandleFunc("/", s.handleIndex())
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("x-server-name", "goserver")
	}
}
