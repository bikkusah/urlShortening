package controller

import (
	"net/http"
	"strings"
	"time"

	"github.com/bikkusah/urlShortening/constant"
	"github.com/bikkusah/urlShortening/database"
	"github.com/bikkusah/urlShortening/helper"
	"github.com/bikkusah/urlShortening/types"
)

func ShortTheUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	longUrl := r.FormValue("long_url")

	code := helper.GenRandomString(6)

	record, _ := database.Mgr.GetUrlFromCode(code)
	if record.UrlCode != "" {
		http.Error(w, "this code is already in use", http.StatusBadRequest)
		return
	}

	url := types.UrlDb{
		CreatedAt: time.Now().Unix(),
		ExpiredAt: time.Now().Unix(),
		UrlCode:   code,
		LongUrl:   longUrl,
		ShortUrl:  constant.BaseUrl + code,
	}

	_, err := database.Mgr.Insert(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Shortened URL: " + url.ShortUrl))
}

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, constant.RedirectUrlPath)

	record, _ := database.Mgr.GetUrlFromCode(code)
	if record.UrlCode == "" {
		http.Error(w, "URL not found", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, record.LongUrl, http.StatusPermanentRedirect)
}
