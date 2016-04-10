package pblog

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine/memcache"

	"github.com/nu7hatch/gouuid"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// Session holds the user's data
type Session struct {
	ID       string
	Bucket   string
	Pictures map[string]string
	res      http.ResponseWriter
	ctx      context.Context
}

func getSession(res http.ResponseWriter, req *http.Request) *Session {

	var s Session
	ctx := appengine.NewContext(req)

	cookie, err := req.Cookie("sessionid")
	if err != nil || cookie.Value == "" {
		sessionID, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "sessionid",
			Value: sessionID.String(),
		}
	}

	item, err := memcache.Get(ctx, cookie.Value)
	if err != nil {

		// put file names from gcs in s.Pictures
		s.ID = cookie.Value
		s.ctx = ctx
		s.listBucket()

		// create memcache.Item
		bs, err := json.Marshal(s)
		if err != nil {
			log.Errorf(ctx, "ERROR memcache.Get json.Marshal: %s", err)
		}
		item = &memcache.Item{
			Key:   cookie.Value,
			Value: bs,
		}
	}

	json.Unmarshal(item.Value, &s)
	s.ID = cookie.Value
	s.ctx = ctx
	s.res = res

	// store in memcache
	s.putSession()

	return &s
}

func (s *Session) putSession() {
	bs, err := json.Marshal(s)
	if err != nil {
		return
	}

	memcache.Set(s.ctx, &memcache.Item{
		Key:   s.ID,
		Value: bs,
	})

	http.SetCookie(s.res, &http.Cookie{
		Name:  "sessionid",
		Value: s.ID,
	})
}
