package model

type AccountListResponse struct {
	Status string `json:"status"`
	Meta   struct {
		Count int `json:"count"`
	} `json:"meta"`
	Data []struct {
		Nickname  string `json:"nickname"`
		AccountID int64  `json:"account_id"`
	} `json:"data"`
}
