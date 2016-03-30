package mem

import (
	"github.com/nu7hatch/gouuid"
	"net/http"
	"log"
)

func newVisitor(req *http.Request) (*http.Cookie, error) {
	m := initialModel()
	return makeCookie(m, req)
}

func currentVisitor(m model, req *http.Request) (*http.Cookie, error) {
	return makeCookie(m, req)
}

func initialModel() model {
	id, err := uuid.NewV4()
	if err != nil {
		log.Println("ERROR newVisitor uuid.NewV4", err)
		return nil, err
	}
	m := model{
		Name:  "",
		State: false,
		Pictures: []string{
			"one.jpg",
		},
		ID: id.String(),
	}
	return m
}
