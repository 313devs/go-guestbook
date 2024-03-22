package handler

import (
	"fmt"
	"net/http"
)

type Signiture struct{}

func (s *Signiture) Sign(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signed!")
}

func (s *Signiture) GetSigniture(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sign")
}
