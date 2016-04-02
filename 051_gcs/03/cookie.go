package skyhdd

import (
"net/http"
	"encoding/json"
	"fmt"
)

func putCookie(res http.ResponseWriter, req *http.Request, fname string) ([]string, error) {
	var xs []string
	cookie, err := req.Cookie("file-names")
	if err != nil {
		return nil, fmt.Errorf("ERROR handler req.Cookie: %s", err)

	}
	if cookie != nil {
		err = json.Unmarshal([]byte(cookie.Value), &xs)
		if err != nil {
			return nil, fmt.Errorf("ERROR handler json.Unmarshal: %s", err)
		}
	}

	xs = append(xs, fname)
	bs, err := json.Marshal(xs)

	http.SetCookie(res, &http.Cookie{
		Name:  "file-names",
		Value: string(bs),
	})
	return xs, nil
}
