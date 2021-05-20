package router

import (
	"github.com/gorilla/pat"
	"github.com/potatowhite/web/study-oauth/oauth/google"
	"github.com/urfave/negroni"
)

func Router() *negroni.Negroni {
	mux := pat.New()
	mux.HandleFunc("/auth/google/login", google.RedirectToGoogleLoginPage)
	mux.HandleFunc("/auth/google/callback", google.CallBackOAuthResultAndPrintUserInfo)

	ng := negroni.Classic()
	ng.UseHandler(mux)

	return ng
}
