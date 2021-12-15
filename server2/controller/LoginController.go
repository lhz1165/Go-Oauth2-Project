package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-session/session"
)

func LoginController(c *gin.Context) {
	w := c.Writer
	r := c.Request
	store, err := session.Start(c.Request.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if r.Form == nil {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	store.Set("LoggedInUserID", r.Form.Get("username"))
	store.Save()
	c.Redirect(http.StatusFound, "/auth")
	//c.Redirect(http.StatusFound, "/oauth/authorize")
}

func LoginHtmlController(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{
		"title": "login",
	})
}

func AuthController(c *gin.Context) {
	w := c.Writer
	r := c.Request

	store, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}
	c.HTML(200, "auth.html", gin.H{
		"title": "auth",
	})
}
