package models

type Tweet_url struct {
	ExpandedUrl string `json:"expanded_url"`
}

type Entities struct {
	Urls []Tweet_url `json:"urls"`
}

type Tweet struct {
	Entities Entities `json:"entities"`
	Lang     string   `json:"lang"`
	Text     string   `json:"text"`
	Id       string   `json:"id"`
}

type ResponseData struct {
	Tweets []Tweet `json:"data"`
}
