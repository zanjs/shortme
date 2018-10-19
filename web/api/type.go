package api

type version struct {
	Version string `json:"version"`
}

type errorResp struct {
	Msg string `json:"msg"`
}

type shortReq struct {
	LongURL string `json:"longURL"`
}

type shortResp struct {
	ShortURL string `json:"shortURL"`
}

type expandReq struct {
	ShortURL string `json:"shortURL"`
}

type expandResp struct {
	LongURL string `json:"longURL"`
}
