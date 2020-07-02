package cav2

var (
	//OAuthToken - Twitch oauth's token
	OAuthToken string
	//WriteSession - Unknown
	WriteSession = false
)

const (
	twitchCodeFile = "twitch_code.txt"
)

//SetOAuth Takes in a twitch oauth code generated with OAuth.html
func SetOAuth(token string) {
	OAuthToken = token
}

//ReadTokenFromDisk Reads the auth token from a file
func ReadTokenFromDisk() {
	code := readStringFromFile(twitchCodeFile)
	SetOAuth(code)
}
