package transport

import (
	"encoding/base64"
	"net/http"
)

func NewTransport(rt http.RoundTripper, user, pass string) Transport {
	return Transport{
		username: user,
		password: pass,
		rt:       rt,
	}
}

type Transport struct {
	username string
	password string
	// keep a reference to the client's original transport
	rt http.RoundTripper
}

func (t *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	// set the Authorization headers here
	r.Header.Set("Authorization", "Basic "+encodeB64(t.username+":"+t.password))
	res, err := t.rt.RoundTrip(r)

	return res, err
}

func encodeB64(d string) string {
	return base64.StdEncoding.EncodeToString([]byte(d))
}
