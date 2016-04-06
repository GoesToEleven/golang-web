package skyhdd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func putCookie(res http.ResponseWriter, req *http.Request, fname string) error {
	mss := make(map[string]bool)
	cookie, _ := req.Cookie("file-names")
	if cookie != nil {
		bs, err := base64.URLEncoding.DecodeString(cookie.Value)
		if err != nil {
			return fmt.Errorf("ERROR putCookie base64.URLEncoding.DecodeString: %s", err)
		}
		err = json.Unmarshal(bs, &mss)
		if err != nil {
			return fmt.Errorf("ERROR putCookie json.Unmarshal: %s", err)
		}
	}

	mss[fname] = true
	bs, err := json.Marshal(mss)
	if err != nil {
		return fmt.Errorf("ERROR putCookie json.Marshal: ", err)
	}
	b64 := base64.URLEncoding.EncodeToString(bs)

	// FYI
	ctx := appengine.NewContext(req)
	log.Infof(ctx, "COOKIE JSON: %s", string(bs))

	http.SetCookie(res, &http.Cookie{
		Name:  "file-names",
		Value: b64,
	})
	return nil
}


func getCookie(res http.ResponseWriter, req *http.Request) (map[string]bool, error) {

	mss := make(map[string]bool)

	cookie, err := req.Cookie("file-names")
	if err != nil {
		return nil, fmt.Errorf("ERROR getCookie req.Cookie: %s", err)
	}

	bs, err := base64.URLEncoding.DecodeString(cookie.Value)
	if err != nil {
		return nil, fmt.Errorf("ERROR getCookie base64.URLEncoding.DecodeString: %s", err)
	}

	err = json.Unmarshal(bs, &mss)
	if err != nil {
		return nil, fmt.Errorf("ERROR getCookie json.Unmarshal: %s", err)
	}

	return mss, nil
}
