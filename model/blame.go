package model

type WowsBlameVersion struct {
	Ok   bool `json:"ok"`
	Data struct {
		Version string `json:"version"`
	} `json:"data"`
	Message string `json:"message"`
}

type WowsBlamePlayerPayload struct {
	WinRate          float64 `json:"win_rate"`
	TotalBattleCount int     `json:"total_battle_count"`
	AccountName      string  `json:"account_name"`
	AccountID        string  `json:"account_id"`
	ShipName         string  `json:"ship_name"`
	ShipID           string  `json:"ship_id"`
	ShipType         string  `json:"ship_type"`
	ShipWinRate      float64 `json:"ship_win_rate"`
	ShipBattleCount  int     `json:"ship_battle_count"`
}

type WowsBlamePlayer struct {
	Ok      bool                   `json:"ok"`
	Data    WowsBlamePlayerPayload `json:"data"`
	Message string                 `json:"message"`
}

type WowsBlameSetupParam struct {
	ApplicationID string `json:"application_id"`
	GamePath      string `json:"game_path"`
}

type WowsBlameSetupResponse struct {
	Ok           bool   `json:"ok"`
	AppIDOk      bool   `json:"app_id_ok"`
	AppIDMessage string `json:"app_id_message"`
	PathOk       bool   `json:"path_ok"`
	PathMessage  string `json:"path_message"`
	Message      string `json:"message"`
}

type WowsBlameMatchResponse struct {
	Ok      bool      `json:"ok"`
	Message string    `json:"message"`
	Data    ArenaInfo `json:"data"`
}
