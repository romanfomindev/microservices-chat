package storage

type storage struct {
	AccessToken string
}

var gloabalStoarge storage

func Init() {
	gloabalStoarge = storage{}
}

func SetAccessToken(token string) {
	gloabalStoarge.AccessToken = token
}

func GetAccessToken() string {
	return gloabalStoarge.AccessToken
}
