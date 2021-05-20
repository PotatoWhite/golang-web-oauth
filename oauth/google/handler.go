package google

import (
	"fmt"
	"log"
	"net/http"
)

// redirect to google login page
func RedirectToGoogleLoginPage(writer http.ResponseWriter, request *http.Request) {
	// generate transantionId
	state := generateStateOauthCookie(writer)
	url := GoogleOAuthConfig.AuthCodeURL(state)
	http.Redirect(writer, request, url, http.StatusTemporaryRedirect)
}

// redirected from google login page
func CallBackOAuthResultAndPrintUserInfo(writer http.ResponseWriter, request *http.Request) {
	oauthstate, _ := request.Cookie("oauthstate")

	// check fail and fast exit
	if request.FormValue("state") != oauthstate.Value {
		log.Printf("invalid google oauth state cookie:%s state:%s\n", oauthstate.Value, request.FormValue("state"))
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}

	// get user info
	data, err := getGoogleUserInfo(request.FormValue("code"))
	// handle error
	if err != nil {
		log.Println(err.Error())
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}

	// print user info to web browser
	fmt.Fprint(writer, string(data))
}
