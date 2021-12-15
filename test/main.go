package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Token struct {
	AccessToken string `json:"access_token"`

	TokenType string `json:"token_type,omitempty"`

	RefreshToken string `json:"refresh_token,omitempty"`

	Expiry time.Time `json:"expiry,omitempty"`
}

type MyHandler struct {
	http.Handler
}

//实现接口的ServeHTTP方法
func (xxx *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "通过实现接口Handler的ServeHTTP方法来创建服务器")
	token := Token{AccessToken: "132456", TokenType: "aa", RefreshToken: "bb", Expiry: time.Now()}
	e := json.NewEncoder(w)
	e.Encode(token)
}

func main() {
	myHandler := &MyHandler{}
	http.Handle("/", myHandler)
	http.ListenAndServe(":8888", nil)
}
