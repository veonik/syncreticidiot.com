package main

//go:generate go-bindata -prefix "../../assets" -pkg internal --nometadata -o "internal/assets.go" ../../assets

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"

	"./internal"

	"github.com/tyler-sommer/stick"
)

var listenHost string
var listenPortHTTPS int
var listenPortHTTP int

var tlsKeyFile string
var tlsCertFile string

var renderedMarkup []byte

func init() {
	flag.StringVar(&listenHost, "host", "syncreticidiot.com", "Listen hostname")
	flag.IntVar(&listenPortHTTPS,"port", 443, "HTTPS listen port")
	flag.IntVar(&listenPortHTTP, "http-port", 80, "HTTP listen port")
	flag.StringVar(&tlsKeyFile, "key-file", ".syncreticidiot.com/cert/privkey.pem", "TLS private key file")
	flag.StringVar(&tlsCertFile, "cert-file", ".syncreticidiot.com/cert/fullchain.pem", "TLS certificate chain file")
	flag.Parse()

	tpl := string(internal.MustAsset("index.html.twig"))
	sk := string(internal.MustAsset("id_rsa.pub"))
	gk := string(internal.MustAsset("gpg-pubkey.asc"))

	b := &bytes.Buffer{}
	env := stick.New(&stick.StringLoader{})
	err := env.Execute(tpl, b, map[string]stick.Value{
		"ssh_key": sk,
		"pgp_key": gk,
	})
	if err != nil {
		panic(err)
	}

	renderedMarkup = b.Bytes()
}

func main() {
	go func() {
		err := listenAndServe()
		if err != nil {
			log.Println(err.Error())
		}
	}()
	go func() {
		err := listenAndServeTLS()
		if err != nil {
			log.Println(err.Error())
		}
	}()
	select {}
}

func listenAndServe() error {
	l := fmt.Sprintf("%s:%d", listenHost, listenPortHTTP)
	ls := fmt.Sprintf("https://%s:%d/", listenHost, listenPortHTTPS)
	srv := &http.Server{Addr: l, Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, ls, http.StatusMovedPermanently)
		if _, err := fmt.Fprintf(w, `<a href="%s">Redirecting...</a>`, ls); err != nil {
			log.Println("error writing response:", err.Error())
		}
	})}
	return srv.ListenAndServe()
}

func listenAndServeTLS() error {
	l := fmt.Sprintf("%s:%d", listenHost, listenPortHTTPS)
	h := http.HandlerFunc(handler)
	srv := &http.Server{Addr: l, Handler: h}
	return srv.ListenAndServeTLS(tlsCertFile, tlsKeyFile)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=4203600")
	w.Header().Set("Content-Type", "text/html")
	if _, err := w.Write(renderedMarkup); err != nil {
		log.Println("error writing response:", err.Error())
	}
}
