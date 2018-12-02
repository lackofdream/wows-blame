package model

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

type ArenaVehiclesInfo struct {
	ShipId   int    `json:"shipId"`
	Relation int    `json:"relation"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
}

type ArenaInfo struct {
	ClientVersionFromXml string              `json:"clientVersionFromXml"`
	GameMode             int                 `json:"gameMode"`
	ClientVersionFromExe string              `json:"clientVersionFromExe"`
	MapDisplayName       string              `json:"mapDisplayName"`
	MapId                int                 `json:"mapId"`
	MatchGroup           string              `json:"matchGroup"`
	Duration             int                 `json:"duration"`
	GameLogic            string              `json:"gameLogic"`
	Name                 string              `json:"name"`
	Scenario             string              `json:"scenario"`
	PlayerID             int                 `json:"playerID"`
	Vehicles             []ArenaVehiclesInfo `json:"vehicles"`
	PlayersPerTeam       int                 `json:"playersPerTeam"`
	DateTime             string              `json:"dateTime"`
	MapName              string              `json:"mapName"`
	PlayerName           string              `json:"playerName"`
	ScenarioConfigId     int                 `json:"scenarioConfigId"`
	TeamsCount           int                 `json:"teamsCount"`
	Logic                string              `json:"logic"`
	PlayerVehicle        string              `json:"playerVehicle"`
}
