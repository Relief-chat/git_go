// cf. https://github.com/0c34/govwa/blob/139693e56406b5684d2a6ae22c0af90717e149b8/user/session/session.go

package session

import (
	"log"
	"fmt"
	"net/http"
	"govwa/util/config"
	"github.com/gorilla/sessions"
)

type Self struct{}

func New() *Self {
	return &Self{}
}

var store = sessions.NewCookieStore([]byte(config.Cfg.Sessionkey))

func (self *Self) SetSession(w http.ResponseWriter, r *http.Request, data map[string]string) {
	session, err := store.Get(r, "govwa")

	if err != nil {
		log.Println(err.Error())
	}

    // bad
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: false, //set to false for xss :)
        Secure: true,
	}

	session.Values["govwa_session"] = true

	//create new session to store on server side
	if data != nil {
		for key, value := range data {
			session.Values[key] = value
		}
	}
	err = session.Save(r, w) //safe session and send it to client as cookie

		if err != nil {
			log.Println(err.Error())
		}
}

func (self *Self) SetSession_ok(w http.ResponseWriter, r *http.Request, data map[string]string) {
	session, err := store.Get(r, "govwa")

	if err != nil {
		log.Println(err.Error())
	}

    // ok
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true, 
        Secure: true,
	}

	session.Values["govwa_session"] = true
 
	if data != nil {
		for key, value := range data {
			session.Values[key] = value
		}
	}
	err = session.Save(r, w) 

		if err != nil {
			log.Println(err.Error())
		}
}