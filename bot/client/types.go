package client

type UpdatesResponse struct {
	Ok     bool      `json:"ok"`
	Result []Updates `json:"result"`
}

type Updates struct {
	Id      int          `json:"update_id"`
	Message *NewMessages `json:"message"`
}

type NewMessages struct {
	Text string   `json:"text"`
	From UserInfo `json:"from"`
	Chat ChatInfo `json:"chat"`
}

type UserInfo struct {
	UserName string `json:"username"`
}

type ChatInfo struct {
	Id int `json:"id"`
}
