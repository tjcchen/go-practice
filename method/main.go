package main

import (
	// "net/http"
	a "github.com/tjcchen/go-practice/method/a"
)

type Server struct {
	URL string
	client string
}

// For reference purpose - receiver usages
// func (s *Server) StartTLS() {
// 	if s.URL != "" {
// 		panic("Server already started")
// 	}
// 	if s.client == nil {
// 		s.client = &http.Client{Transport: &http.Transport{}}
// 	}
// }

func main() {
	// method example:
	// func (recv receiver_type) methodName(parameter_list) (return_value_list)

	a.PrintMsg("hello world") // hello world
}