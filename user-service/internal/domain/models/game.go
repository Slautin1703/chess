package models

type Game struct {
	ID        int    `json:"id"`
	GameType  string `json:"game_type"`
	Duration  string `json:"duration"`
	Status    string `json:"status"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Score     string `json:"score"`
}
