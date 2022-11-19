package order

type SwitchPositionModeResponse struct {
	Status string `json:"status"`
	Data   []struct {
		MarginAccount string `json:"margin_account"`
		PositionMode  string `json:"position_mode"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}