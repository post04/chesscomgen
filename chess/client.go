package chess

import (
	"fmt"
	"time"

	http "github.com/Carcraftz/fhttp"

	"github.com/Carcraftz/cclient"
	tls "github.com/Carcraftz/utls"
)

// Client holds http.Client
type Client struct {
	Proxy        string
	HTTPClient   http.Client
	UserAgent    string
	Cookies      map[string]string
	CSRF         string
	SessID       string
	StartTime    int64
	Bearer       string
	UUID         string
	RefreshToken string
	LoginToken   string
	SessionID    string
}

// NewClient returns a *Client
func NewClient(proxy string) *Client {
	fmt.Println(proxy)
	s := &Client{
		Proxy: proxy,
	}
	httpClient, err := cclient.NewClient(tls.HelloIOS_Auto, proxy, true, 60*time.Second)
	if err != nil {
		panic(err)
	}
	s.HTTPClient = httpClient
	s.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
	s.Cookies = make(map[string]string)
	return s
}
