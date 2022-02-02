package main

import (
	"net"
	"os"
	"time"

	"github.com/matryer/go-oauth/oauth"
)

var conn net.Conn

var (
	authClient *oauth.Client
	creds      *oauth.Credentials
)

func setupTwitterAuth() {
	ts := struct {
		ConsumerKey    string
		ConsumerSecret string
		AccessToken    string
		AccessSecret   string
	}{
		ConsumerKey:    os.Getenv("API_KEY"),
		ConsumerSecret: os.Getenv("API_KEY_SECRET"),
		AccessToken:    os.Getenv("ACCESS_TOKEN"),
		AccessSecret:   os.Getenv("ACCESS_TOKEN_SECRET"),
	}
	creds = &oauth.Credentials{
		Token:  ts.AccessToken,
		Secret: ts.AccessSecret,
	}
	authClient = &oauth.Client{
		Credentials: oauth.Credentials{
			Token:  ts.ConsumerKey,
			Secret: ts.ConsumerSecret,
		},
	}
}

func dial(netw, addr string) (net.Conn, error) {
	if conn != nil {
		conn.Close()
		conn = nil
	}
	netc, err := net.DialTimeout(netw, addr, 5*time.Second)
	if err != nil {
		return nil, err
	}
	conn = netc

	return netc, nil
}
