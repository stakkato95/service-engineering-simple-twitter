package dto

type ResponseDto struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

type TweetCountDto struct {
	TweetCount int `json:"tweetCount"`
}
