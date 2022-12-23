package account

type GetAccountPositionResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol         string  `json:"symbol"`
		ContractCode   string  `json:"contract_code"`
		Volume         float64 `json:"volume"`
		Available      float64 `json:"available"`
		Frozen         float64 `json:"frozen"`
		CostOpen       float64 `json:"cost_open"`
		CostHold       float64 `json:"cost_hold"`
		ProfitUnreal   float64 `json:"profit_unreal"`
		ProfitRate     float64 `json:"profit_rate"`
		LeverRate      int     `json:"lever_rate"`
		PositionMargin float64 `json:"position_margin"`
		Direction      string  `json:"direction"`
		Profit         float64 `json:"profit"`
		LastPrice      float64 `json:"last_price"`
		MarginAsset    string  `json:"margin_asset"`
		MarginMode     string  `json:"margin_mode"`
		MarginAccount  string  `json:"margin_account"`
		ContractType   string  `json:"contract_type"`
		Pair           string  `json:"pair"`
		BusinessType   string  `json:"business_type"`
		PositionMode   string  `json:"position_mode"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
