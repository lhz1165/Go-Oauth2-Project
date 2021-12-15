package controller

import (
	"encoding/json"
	"fmt"
	"goauthtest/server2/config"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-session/session"
)

func AuthorizeController(c *gin.Context) {
	srv := config.GetServ()
	w := c.Writer
	r := c.Request
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	r.Form = form
	fmt.Printf("request map        %v\n", form)
	store.Delete("ReturnUri")
	store.Save()
	//用户没登陆过跳转到登录页面 UserAuthorizationHandler来处理
	err = srv.HandleAuthorizeRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func TokenController(c *gin.Context) {
	srv := config.GetServ()
	w := c.Writer
	r := c.Request
	err := srv.HandleTokenRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TestController(c *gin.Context) {
	srv := config.GetServ()
	w := c.Writer
	r := c.Request
	token, err := srv.ValidationBearerToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"client_id":  token.GetClientID(),
		"user_id":    token.GetUserID(),
	}
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	e.Encode(data)
}
