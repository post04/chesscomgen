package chess

import (
	"strings"

	http "github.com/Carcraftz/fhttp"
)

// SetCookies takes Set-Cookie headers and makes them halal
func (c *Client) SetCookies(headers http.Header) {
	for _, cookie := range headers["Set-Cookie"] {
		parts := strings.Split(cookie, "; ")
		c.Cookies[strings.Split(parts[0], "=")[0]] = strings.Join(strings.Split(parts[0], "=")[1:], "=")
	}
}

// FormatCookies takes c.Cookies and makes it into a useable string
func (c *Client) FormatCookies() string {
	cookies := ""
	for key, value := range c.Cookies {
		cookies += key + "=" + value + "; "
	}
	return cookies
}
