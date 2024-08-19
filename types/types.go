package types

type ShortUrlBody struct {
	LongUrl string `json:"long_url"`
}

type UrlDb struct {
	ID        int64  `json:"id"`
	UrlCode   string `json:"url_code"`
	LongUrl   string `json:"long_url"`
	ShortUrl  string `json:"short_url"`
	CreatedAt int64  `json:"created_at"`
	ExpiredAt int64  `json:"expired_at"`
}
