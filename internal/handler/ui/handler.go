package ui

import (
	"Weather/internal/helper"
	"errors"
	"log"
	"net/http"
)

func (h Handler) Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := helper.GetCookie(r)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				cookie = helper.GenerateAndSetCookie(w)
				h.srv.UserSrv.AddUser(cookie)
			} else {
				log.Println(err)
				return
			}
		}

		next(w, r)
	}
}
