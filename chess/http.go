package chess

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"

	http "github.com/Carcraftz/fhttp"
)

func formatHeaders(s string) http.Header {
	h := http.Header{}
	for _, header := range strings.Split(s, "|") {
		parts := strings.Split(header, ": ")
		h.Set(parts[0], parts[1])
	}
	return h
}

var csrfRegex = regexp.MustCompile(`"token":"[a-z90-9]+\.[A-z0-9-_]+\.[A-z0-9-_]+`)

// GetCSRF - gets init cookies and CSRF token
func (c *Client) GetCSRF() error {
	req, err := http.NewRequest("GET", "https://www.chess.com/register", nil)
	if err != nil {
		return err
	}
	req.Header = formatHeaders(`accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7|accept-encoding: gzip, deflate, br|accept-language: en-US,en;q=0.9|cache-control: no-cache|pragma: no-cache|sec-ch-ua: "Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111"|sec-ch-ua-mobile: ?0|sec-ch-ua-platform: "Windows"|sec-fetch-dest: document|sec-fetch-mode: navigate|sec-fetch-site: none|sec-fetch-user: ?1|upgrade-insecure-requests: 1`)
	req.Header.Set("User-Agent", c.UserAgent)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	c.SetCookies(resp.Header)
	b, err := c.DecodeBody(resp.Header, resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("response code was " + fmt.Sprint(resp.StatusCode))
	}
	// ! get csrf
	match := csrfRegex.FindString(string(b))
	if match == "" {
		return errors.New("unable to find csrf")
	}
	c.CSRF = strings.Split(match, ":")[1][1:]
	return nil
}

var sessIDRegex = regexp.MustCompile(`[0-9a-z]{32}`)

// GetFPInit - gets the init fingerprinting script
func (c *Client) GetFPInit() error {
	req, err := http.NewRequest("GET", "https://prod01.kaxsdc.com/collect/sdk?m=850100", nil)
	if err != nil {
		return err
	}
	req.Header = formatHeaders(`Accept: */*|Accept-Encoding: gzip, deflate, br|Accept-Language: en-US,en;q=0.9|Cache-Control: no-cache|Connection: keep-alive|Host: prod01.kaxsdc.com|Pragma: no-cache|Referer: https://www.chess.com/|sec-ch-ua: "Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111"|sec-ch-ua-mobile: ?0|sec-ch-ua-platform: "Windows"|Sec-Fetch-Dest: script|Sec-Fetch-Mode: no-cors|Sec-Fetch-Site: cross-site`)
	req.Header.Set("User-Agent", c.UserAgent)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := c.DecodeBody(resp.Header, resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("response code was " + fmt.Sprint(resp.StatusCode))
	}
	// ! get sessionID
	match := sessIDRegex.FindString(string(b))
	if match == "" {
		return errors.New("unable to find sessid")
	}
	c.SessID = match
	return nil
}

var kaCookieValueRegex = regexp.MustCompile(`cvalue="[a-z0-9]+`)

// GetKaCookie - gets the cookie value
func (c *Client) GetKaCookie() error {
	req, err := http.NewRequest("POST", "https://prod01.kaxsdc.com/collect/kasupport", strings.NewReader(`m=850100&s=`+c.SessID))
	req.Header = formatHeaders(`Accept: */*|Accept-Encoding: gzip, deflate, br|Accept-Language: en-US,en;q=0.9|Cache-Control: no-cache|Connection: keep-alive|Content-Length: 43|Content-type: application/x-www-form-urlencoded|Host: prod01.kaxsdc.com|Origin: https://www.chess.com|Pragma: no-cache|Referer: https://www.chess.com/|sec-ch-ua: "Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111"|sec-ch-ua-mobile: ?0|sec-ch-ua-platform: "Windows"|Sec-Fetch-Dest: empty|Sec-Fetch-Mode: cors|Sec-Fetch-Site: cross-site`)
	req.Header.Set("User-Agent", c.UserAgent)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := c.DecodeBody(resp.Header, resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("response code was " + fmt.Sprint(resp.StatusCode))
	}
	// ! get cookie value
	match := kaCookieValueRegex.FindString(string(b))
	if match == "" {
		return errors.New("unable to find sessid")
	}
	c.Cookies["cdn.chesscom.850100.ka.ck"] = strings.Split(match, `"`)[1]
	return c.submitCookieStore()
}

func (c *Client) submitCookieStore() error {
	req, err := http.NewRequest("POST", "https://prod01.kaxsdc.com/collect/cookiestore", strings.NewReader(fmt.Sprintf(`m=850100&s=%s&k=%s`, c.SessID, c.Cookies["cdn.chesscom.850100.ka.ck"])))
	if err != nil {
		return err
	}
	req.Header = formatHeaders(`Accept: */*|Accept-Encoding: gzip, deflate, br|Accept-Language: en-US,en;q=0.9|Cache-Control: no-cache|Connection: keep-alive|Content-type: application/x-www-form-urlencoded|Host: prod01.kaxsdc.com|Origin: https://www.chess.com|Pragma: no-cache|Referer: https://www.chess.com/|sec-ch-ua: "Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111"|sec-ch-ua-mobile: ?0|sec-ch-ua-platform: "Windows"|Sec-Fetch-Dest: empty|Sec-Fetch-Mode: cors|Sec-Fetch-Site: cross-site`)
	req.Header.Set("User-Agent", c.UserAgent)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// SendMD1 gets the logo.html (which gets CID) then sends the md1 payload
func (c *Client) SendMD1() error {
	err := c.getCID()
	if err != nil {
		return err
	}
	c.StartTime = time.Now().UnixMilli()
	req, err := http.NewRequest("POST", "https://prod01.kaxsdc.com/md", strings.NewReader(fmt.Sprintf(`et=1&ln=en-US&e=%v&t0=240&tf=300&ta=240&sa=1032x1920&cd=24&sd=1080x1920&fd=0&lh=9fc4c19727632d3c59628c98d704ff5e&s=%s&m=850100&n=clientdata`, time.Now().UnixMilli(), c.SessID)))
	req.Header = formatHeaders(`Accept: */*|Accept-Encoding: gzip, deflate, br|Accept-Language: en-US,en;q=0.9|Cache-Control: no-cache|Connection: keep-alive|Content-Type: application/x-www-form-urlencoded|Host: prod01.kaxsdc.com|Origin: https://prod01.kaxsdc.com|Pragma: no-cache|Referer: https://prod01.kaxsdc.com/logo.htm?m=850100&s=` + c.SessID + `|sec-ch-ua: "Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111"|sec-ch-ua-mobile: ?0|sec-ch-ua-platform: "Windows"|Sec-Fetch-Dest: empty|Sec-Fetch-Mode: cors|Sec-Fetch-Site: same-origin`)
	req.Header.Set("User-Agent", c.UserAgent)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (c *Client) getCID() error {
	req, err := http.NewRequest("GET", fmt.Sprintf(`https://prod01.kaxsdc.com/logo.htm?m=850100&s=%s`, c.SessID), nil)
	if err != nil {
		return err
	}
	req.Header = formatHeaders(`Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7|Accept-Encoding: gzip, deflate, br|Accept-Language: en-US,en;q=0.9|Cache-Control: no-cache|Connection: keep-alive|Host: prod01.kaxsdc.com|Pragma: no-cache|Referer: https://www.chess.com/|sec-ch-ua: "Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111"|sec-ch-ua-mobile: ?0|sec-ch-ua-platform: "Windows"|Sec-Fetch-Dest: iframe|Sec-Fetch-Mode: navigate|Sec-Fetch-Site: cross-site|Upgrade-Insecure-Requests: 1`)
	req.Header.Set("User-Agent", c.UserAgent)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// SendFinish finishes data collection
func (c *Client) SendFinish() error {
	req, err := http.NewRequest("POST", "https://prod01.kaxsdc.com/fin", strings.NewReader(fmt.Sprintf(`s=%s&m=850100&n=collect-end&com=true&et=%v`, c.SessID, time.Now().UnixMilli()-c.StartTime)))
	if err != nil {
		return err
	}
	req.Header = formatHeaders(`Accept: */*|Accept-Encoding: gzip, deflate, br|Accept-Language: en-US,en;q=0.9|Cache-Control: no-cache|Connection: keep-alive|Content-Type: application/x-www-form-urlencoded|Host: prod01.kaxsdc.com|Origin: https://prod01.kaxsdc.com|Pragma: no-cache|Referer: https://prod01.kaxsdc.com/logo.htm?m=850100&s=` + c.SessID + `|sec-ch-ua: "Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111"|sec-ch-ua-mobile: ?0|sec-ch-ua-platform: "Windows"|Sec-Fetch-Dest: empty|Sec-Fetch-Mode: cors|Sec-Fetch-Site: same-origin`)
	req.Header.Set("User-Agent", c.UserAgent)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

// SubmitSignup - creates a new chess.com account
func (c *Client) SubmitSignup(username, password, email string) error {
	//! kountSessionId = session ID
	//! fingeprint = fingerprintjs fingerprint
	payload := url.Values{}
	payload.Set("registration[username]", username)
	payload.Set("registration[email]", email)
	payload.Set("registration[password]", password)
	payload.Set("registration[skillLevel]", "1")
	payload.Set("registration[timezone]", "America/New_York")
	payload.Set("registration[_token]", c.CSRF)
	payload.Set("kountSessionId", c.SessID)
	payload.Set("fingerprint", c.makeFP())
	payload.Set("registration[friend]", "")

	req, err := http.NewRequest("POST", "https://www.chess.com/register", strings.NewReader(payload.Encode()))
	req.Header = formatHeaders(`accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7|accept-encoding: gzip, deflate, br|accept-language: en-US,en;q=0.9|cache-control: no-cache|content-type: application/x-www-form-urlencoded|origin: https://www.chess.com|pragma: no-cache|referer: https://www.chess.com/register|sec-ch-ua: "Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111"|sec-ch-ua-mobile: ?0|sec-ch-ua-platform: "Windows"|sec-fetch-dest: document|sec-fetch-mode: navigate|sec-fetch-site: same-origin|sec-fetch-user: ?1|upgrade-insecure-requests: 1`)
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Cookie", c.FormatCookies())
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 302 {
		return errors.New("status code was not 302")
	}
	if resp.Header["Location"][0] != "https://www.chess.com/home" {
		return errors.New("location header is " + resp.Header["Location"][0])
	}
	c.SetCookies(resp.Header)
	return nil
}

const letterBytes = "abcdef0123456789"

// RandStringBytes generates a random string x letters long
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// KEY is a global key used in the sha1 hash
const KEY = "ab0a13311f98bf844de70c496f53ead356780866753b532a096ab26483601574"

// CreateAccount registers a new account
func (c *Client) CreateAccount(Username, Password, Email string) error {
	payload := fmt.Sprintf(`username=%s&password=%s&email=%s&deviceId=%s&clientId=1bc9f2f2-4961-11ed-8971-f9a8d47c7a48&onboardingType=`, Username, Password, Email, RandStringBytes(32))
	req, err := http.NewRequest("POST", "https://api.chess.com/v1/users?signed=Android4.5.14-"+createHash("POST/v1/users"+payload+KEY), strings.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header = formatHeaders(`Host: api.chess.com|User-Agent: Chesscom-Android/4.5.14-googleplay (Android/9; SM-G988N; en_US; contact #android in Slack)|Accept-Language: en-US|Content-Type: application/x-www-form-urlencoded|Accept-Encoding: gzip, deflate`)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := c.DecodeBody(resp.Header, resp.Body)
	if err != nil {
		return err
	}
	// fmt.Println(string(b))
	if resp.StatusCode != 200 {
		fmt.Println(string(b))
		return errors.New("response code was " + fmt.Sprint(resp.StatusCode))
	}
	f := &BearerResponse{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		return err
	}
	if f.Status != "success" {
		return errors.New("invalid response (real)")
	}
	if f.Data.Oauth.AccessToken == "" {
		return errors.New("no bearer found (real)")
	}
	c.Bearer = "Bearer " + f.Data.Oauth.AccessToken
	c.RefreshToken = f.Data.Oauth.RefreshToken
	c.LoginToken = f.Data.LoginToken
	c.SessionID = f.Data.SessionID
	return nil
}

// AddFriend adds someone as a friend
func (c *Client) AddFriend(Username string) error {
	payload := fmt.Sprintf(`username=%s&message=test&recommendation=false&source=onboardingFlow&method=searchByNameEmailUsername&inContactList=false`, Username)
	req, err := http.NewRequest("POST", "https://api.chess.com/v1/friends/requests?signed=Android4.5.14-"+createHash("POST/v1/friends/requests"+payload+KEY), strings.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header = formatHeaders(`Host: api.chess.com|User-Agent: Chesscom-Android/4.5.14-googleplay (Android/9; SM-G988N; en_US; contact #android in Slack)|Accept-Language: en-US|Content-Type: application/x-www-form-urlencoded|Accept-Encoding: gzip, deflate`)
	req.Header.Set("Authorization", c.Bearer)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := c.DecodeBody(resp.Header, resp.Body)
	if err != nil {
		return err
	}
	// fmt.Println(string(b))
	if resp.StatusCode != 200 {
		fmt.Println(string(b))
		return errors.New("response code was " + fmt.Sprint(resp.StatusCode))
	}
	return nil
}

// ReportPlayer adds someone as a friend
func (c *Client) ReportPlayer(Username string) error {
	payload := fmt.Sprintf(`abuserUsername=%s&reasonId=11`, Username)
	req, err := http.NewRequest("POST", "https://api.chess.com/v1/users/abuse-report?signed=Android4.5.14-"+createHash("POST/v1/users/abuse-report"+payload+KEY), strings.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header = formatHeaders(`Host: api.chess.com|User-Agent: Chesscom-Android/4.5.14-googleplay (Android/9; SM-G988N; en_US; contact #android in Slack)|Accept-Language: en-US|Content-Type: application/x-www-form-urlencoded|Accept-Encoding: gzip, deflate`)
	req.Header.Set("Authorization", c.Bearer)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := c.DecodeBody(resp.Header, resp.Body)
	if err != nil {
		return err
	}
	// fmt.Println(string(b))
	if resp.StatusCode != 200 {
		fmt.Println(string(b))
		return errors.New("response code was " + fmt.Sprint(resp.StatusCode))
	}
	return nil
}
