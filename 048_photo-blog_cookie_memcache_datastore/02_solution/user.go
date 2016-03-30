package mem

import (
	"github.com/nu7hatch/gouuid"
	"net/http"
	"log"
)

func newVisitor(req *http.Request) (*http.Cookie, error) {
	m := initialModel()
	id, err := uuid.NewV4()
	if err != nil {
		log.Println("ERROR newVisitor uuid.NewV4", err)
		return nil, err
	}
	return makeCookie(m, id.String(), req)
}

func currentVisitor(m model, id string, req *http.Request) (*http.Cookie, error) {
	return makeCookie(m, id, req)
}

func initialModel() model {
	m := model{
		Name:  "",
		State: false,
		Pictures: []string{
			"one.jpg",
		},
	}
	return m
}
