package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
	"github.com/elazarl/goproxy/ext/auth"
)

var portNumber = flag.String("port", "12345", "port number")
var basicAuthUser = flag.String("user", "admin", "basic auth user name")
var basicAuthPass = flag.String("pass", "secret", "basic auth user pass")
var verbose = flag.Bool("verbose", false, "print verbose log")

func main() {
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose

	if *basicAuthUser != "" && *basicAuthPass != "" {
		auth.ProxyBasic(proxy, "RELM", func(user, pass string) bool {
			return user == *basicAuthUser && pass == *basicAuthPass
		})
	} else {
		log.Println("basic auth is disabled")
	}

	log.Println("listen: " + *portNumber)
	log.Fatal(http.ListenAndServe(":"+*portNumber, proxy))
}
