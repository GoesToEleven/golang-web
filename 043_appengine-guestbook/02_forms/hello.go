package hello

import (
	"fmt"
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

func root(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, guestbookForm)
}

const guestbookForm = `
<html>
  <body>
    <form action="/sign" method="POST">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`

func sign(res http.ResponseWriter, req *http.Request) {
	err := signTemplate.Execute(res, req.FormValue("content"))
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

var signTemplate = template.Must(template.New("sign").Parse(signTemplateHTML))

const signTemplateHTML = `
<html>
  <body>
    <p>You wrote:</p>
    <pre>{{.}}</pre>
  </body>
</html>
`