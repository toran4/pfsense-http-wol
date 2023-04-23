package main

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

func sendWoL(pfsenseHost string, user string, password string, mac string, targetIf string) bool {
	// Create a new cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	// Disable SSL verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Jar: jar}

	// Get CSRF token from login page
	resp, _ := client.Get(pfsenseHost + "/index.php")
	body, _ := io.ReadAll(resp.Body)
	csrfTokenPattern := regexp.MustCompile(`csrfMagicToken = "(.*?)[;"]`)
	csrfToken := getCrfToken(csrfTokenPattern, body)

	// Log in
	data := url.Values{
		"__csrf_magic": {csrfToken},
		"usernamefld":  {user},
		"passwordfld":  {password},
		"login":        {"Sign In"},
	}
	req, _ := http.NewRequest("POST", pfsenseHost+"/index.php", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ = client.Do(req)
	body, _ = io.ReadAll(resp.Body)

	// Get CSRF token from post-login page
	csrfToken = getCrfToken(csrfTokenPattern, body)

	// Send Wake-on-LAN packet
	data = url.Values{
		"__csrf_magic": {csrfToken},
		"mac":          {mac},
		"if":           {targetIf},
	}
	req, _ = http.NewRequest("POST", pfsenseHost+"/services_wol.php", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ = client.Do(req)
	return resp.StatusCode == 200
}

func getCrfToken(csrfTokenPattern *regexp.Regexp, body []byte) string {
	match := csrfTokenPattern.FindStringSubmatch(string(body))
	csrfToken := match[1]
	return csrfToken
}
