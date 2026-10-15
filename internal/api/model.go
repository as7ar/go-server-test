package api

type Response struct {
	Code    int8   `json:code`
	Content string `json:content`
}
