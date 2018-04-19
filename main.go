package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "max-age=4203600")
	fmt.Fprint(w, `<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8">
    <title>syncretic idiot</title>
    <link href="https://fonts.googleapis.com/css?family=Inconsolata|Montserrat" rel="stylesheet">
    <style>
    html, body {
      height: 100%;
    }
    body {
      margin: 0;
      font-family: Montserrat, sans-serif;
    }
    #universe {
      height: 100%;
      padding: 0;
      margin: 0;
      display: -webkit-box;
      display: -moz-box;
      display: -ms-flexbox;
      display: -webkit-flex;
      display: flex;
      align-items: center;
      justify-content: center;
    }
    h1 {
      font-family: Inconsolata, monospace;
    }
    a, a:visited {
      color: black;
    }
    a:hover, a:active {
      color: #666;
    }
    p.small {
      font-size: .8em;
      text-align: center;
    }
    </style>
  </head>
  <body>
  <div id="universe">
    <div id="everything">
        <h1># &nbsp; syncretic<br># &nbsp; &nbsp; &nbsp; idiot</h1>
        <p>Don't say I didn't warn you.</p>
        <p class="small"><a href="https://keybase.io/veonik">identity</a> - <a href="mailto:nobody@syncreticidiot.com">contact</a></p>
    </div>
  </div>
  </body>
</html>`)
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
	srv := &http.Server{Addr: "syncreticidiot.com:80", Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://syncreticidiot.com/", http.StatusMovedPermanently)
		fmt.Fprintf(w, `<a href="%s">Redirecting...</a>`, "https://syncreticidiot.com/")
	})}
	return srv.ListenAndServe()
}

func listenAndServeTLS() error {
	srv := &http.Server{Addr: "syncreticidiot.com:443", Handler: http.HandlerFunc(handler)}
	return srv.ListenAndServeTLS(".syncreticidiot.com/cert/fullchain.pem", ".syncreticidiot.com/cert/privkey.pem")
}
