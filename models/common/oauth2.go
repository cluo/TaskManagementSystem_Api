package common

import (
	"context"

	"github.com/astaxie/beego"

	"log"

	"golang.org/x/oauth2"
)

// OAuthServiceStruct 定义
type OAuthServiceStruct struct {
}

var Endpoint = oauth2.Endpoint{
	AuthURL:  beego.AppConfig.String("oauth2_AuthURL"),
	TokenURL: beego.AppConfig.String("oauth2_TokenURL"),
}
var state = "task"

func (oauth *OAuthServiceStruct) GetOAuth2Token(username, password string) (oauth2_token string, err error) {

	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     "O3gMpRJtROqw5v9WrYqm9zQlGAUjhFkojlNHFN7V",
		ClientSecret: "57ANxb2GRY3DZ3eQE7u6f1QN7TZKxCtaBWavy5ZPKTZokUjOlqjkIi3w6ZAL0yLd1ajdlhAYLYcfPitYFobLSghtDOIYIhJVlNxACumMyJfObEo7AymQvClM6pQQpIQC",
		Endpoint:     Endpoint,
	}

	tok, err := conf.PasswordCredentialsToken(ctx, username, password)
	if err != nil {
		log.Println(err)
		return
	}
	oauth2_token = tok.AccessToken
	return
}

// func ExampleConfig() {
// 	ctx := context.Background()
// 	conf := &oauth2.Config{
// 		ClientID:     "YOUR_CLIENT_ID",
// 		ClientSecret: "YOUR_CLIENT_SECRET",
// 		Scopes:       []string{"SCOPE1", "SCOPE2"},
// 		Endpoint: oauth2.Endpoint{
// 			AuthURL:  "https://provider.com/o/oauth2/auth",
// 			TokenURL: "https://provider.com/o/oauth2/token",
// 		},
// 	}

// 	// Redirect user to consent page to ask for permission
// 	// for the scopes specified above.
// 	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
// 	fmt.Printf("Visit the URL for the auth dialog: %v", url)

// 	// Use the authorization code that is pushed to the redirect
// 	// URL. Exchange will do the handshake to retrieve the
// 	// initial access token. The HTTP Client returned by
// 	// conf.Client will refresh the token as necessary.
// 	var code string
// 	if _, err := fmt.Scan(&code); err != nil {
// 		log.Fatal(err)
// 	}
// 	tok, err := conf.Exchange(ctx, code)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	client := conf.Client(ctx, tok)
// 	client.Get("...")
// }
