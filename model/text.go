package model

type (
	TextGenReq struct {
		JsonString string `json:"json_string"`
		Template   string `json:"template"`
	}

	TextGenRes struct {
		Text string `json:"text"`
	}
)
