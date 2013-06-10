package httpclient

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Header map[string]string
type Auth struct {
	Type     string
	Username string
	Password string
	Token    string
}

var (
	user *Auth = nil
)

func newTransport() *http.Transport {
	return &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
}

func SetTokenAuth(token string) {
	user = &Auth{Type: "token", Token: token}
}

func SetBasicAuth(username, password string) {
	user = &Auth{Type: "basic", Username: username, Password: password}
}

func Do(t string, url string, header Header, data io.Reader) (*http.Response, error) {
	//trans := newTransport()
	output := &bytes.Buffer{}
	req, err := http.NewRequest(t, url, data)
	req.TransferEncoding = []string{"identity"}
	if err != nil {
		return nil, err
	}
	if user != nil {
		if user.Type == "basic" {
			req.SetBasicAuth(user.Username, user.Password)
		} else {
			req.Header.Add("Authorization", "OAuth "+user.Token)
		}
	}
	for k, v := range header {
		print(k + ": " + v + "\n")
		req.Header.Add(k, v)
	}
	fmt.Printf("%#v\n", req.Header)
	req.Write(output)
	fmt.Println(output)
	return nil, errors.New("blah") //trans.RoundTrip(req)
}

func Get(url string, header Header, data io.Reader) (resp *http.Response, err error) {
	return Do("GET", url, header, data)
}

func Post(url string, header Header, data io.Reader) (resp *http.Response, err error) {
	return Do("POST", url, header, data)
}

func Put(url string, header Header, data io.Reader) (resp *http.Response, err error) {
	return Do("PUT", url, header, data)
}

func Delete(url string, header Header, data io.Reader) (resp *http.Response, err error) {
	return Do("DELETE", url, header, data)
}
