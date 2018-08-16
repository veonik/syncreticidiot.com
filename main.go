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
    #gold-plated #key {
      margin: 0;
      padding: 0 15px;
      border: 1px solid #eee;
      font-family: Inconsolata, monospace;
      font-size: .9em;
      color: #666;
      white-space: pre-wrap;
      word-wrap: break-word;
    }
    </style>
  </head>
  <body>
  <div id="universe">
    <div id="everything">
      <h1># &nbsp; syncretic<br># &nbsp; &nbsp; &nbsp; idiot</h1>
      <p>Don't say I didn't warn you.</p>
      <p class="small">
        <a href="javascript:;" id="pubkey">pubkey</a> - 
        <a href="https://keybase.io/veonik">identity</a> - 
        <a href="mailto:nobody@syncreticidiot.com">contact</a>
      </p>
    </div>
    <div id="gold-plated" style="display: none; opacity: 0">
      <div id="key" class="hidden">
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAD9QDd6bZvuei8hhCn0BVuwBqbjSmiCwLeSHkb2azS+ix
dDeYOXH6FUg+EOvTpj5ExomDKrVx5wdGNrPYRH9XibFRlAqQi542T645eQxKsxYjnTTnhko572KpXQ2
g6CHK7uEw8e3CBjP7OQM5T9UTuXgkjzTxaEqbROheKBSzmPVaubY8TNJbza9TxV0JgfDB1/YzDhSbl8
QkSxnK5HTsmGu3i8SbIoVImFmUZfN+v0LUr42KssupnhRmd5yAbdw1rsT/0yHLBd0I6yf1hM+7S/WKg
lM5uI4FKsJNI0Ee6JRpGofqa2NDievBbduALyXWSMn++2K6uRY/vP2GpLZY/xFXSV5tuILfFPQff5ql
mOCjIR15vzjtVSJez4Qc+ThVr/gEbUsv61BEVobzAY5DqSNmAmuSvr8SQKAr0NiGsdoaTizE2lw7Bz0
GIxamjQVQjaJY9Y1BYG5CyujjDfW/1reuUWy7+a26FwE0zImISrorjlgTK1pU26qf8N4u7IdVya27gS
qldD6jSrJv7UZtSKGSPK9yo8djDlt1JzZy41FPGF8OlToz4kxsSzqN+ULOKl5PSxSKljC3LtrNnmsz+
IjrlypHw4jjmTckzvOL6t5/gCnfD1Gvt/MQ4zb/whHHhHHsYulVnUdZrSpN4MweVEgqNnoWIqlo30fU
OR6UP6iQ9OLepGTiNubpCV+hd+yFzs2GosL3M0YX3ooNKYaKAUPTVa88CJOQyfNfkHkkU09lE0/xjqr
43cENezjYynH6Hd7S6omTr569jZRTuJDHYdVXIekt82du7ABrQOKBK+JqbWIEcxwwne6Zs0phSG4i9+
KYfOCd7cfWRG0enj2P7kBxZ1AmEpRAX17qdd69Ky+sJPZcyduS9jhg3aYoOgcmWdxt6AlWnpLcTQWMC
p2t5N+XOomEJ+9EyPNR0oiXTPGvnYnb0+z7bi22+IrHa2fdiERb+VzrBnB+rJHGXhLBca2e+HlkSdmW
813kLlnG/sz/LcnyrdqZzVYhSHN832+N8gXvR650ZdMYgTY9uARbvQvpArYcuWGwlvL3Ox/rmyugTcC
XU4f3QHMnqgCUpHd0J23/rFxdVNAxLi/cU/TpnT/2lM/sXxUGnUs2bqFm1e/XxfTA98Bb8Z8+fuls5Z
mwf06LzIkchI2LP2LAbeJJ0NzthSl5wg37vznXdSFqiYQHkha11rzom1BoueWw0ke1PSO3LDTplcYgU
+W56SIsvaNqv2B74pEaR9iTO7S7hyZXHffGkNq2s4OfcIvC2yQ0U4TBeI5rDiSa5S5dceTHDmjs95H/
+/AXC/gBN0JF6ioQsWlbSnxqs0Y6kLP7pdjx3FCZ9rCMJ contact@tylersommer.pro
      </div>
    </div>
  </div>
  <script>
  document.addEventListener('DOMContentLoaded', function() {
    let anchor = document.getElementById('pubkey');
    let div = document.getElementById('key');
    anchor.addEventListener('click', function(e) {
      e.preventDefault();
      fadeOut(document.getElementById('everything'), function() {
        fadeIn(document.getElementById('gold-plated'));
      });
    });
  });
  const speed = 250; // in ms
  const delay = 10;
  const step = 1 / (speed / delay);
  const fadeOut = function(el, then) {
    let opacity = 1;
    let interval = setInterval(function() {
      opacity = opacity - step;
      if (opacity <= 0) {
        clearInterval(interval);
      }
      el.style.opacity = opacity;
      if (opacity <= 0) {
        el.style.display = 'none';
        if (then !== undefined) {
          then();
        }
      }
    }, delay);
  }
  const fadeIn = function(el, then) {
    let opacity = 0;
    el.style.opacity = opacity;
    el.style.display = 'block';
    let interval = setInterval(function() {
      opacity = opacity + step;
      if (opacity >= 1) {
        clearInterval(interval);
      }
      el.style.opacity = opacity;
      if (opacity >= 1) {
        if (then !== undefined) {
          then();
        }
      }
    }, delay);
  }
  </script>
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
