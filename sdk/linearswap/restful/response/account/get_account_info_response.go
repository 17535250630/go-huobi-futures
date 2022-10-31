package account

type GetAccountInfoResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		FuturesContractDetail []struct {
			Symbol           string      `json:"symbol"`
			ContractCode     string      `json:"contract_code"`
			MarginPosition   int         `json:"margin_position"`
			MarginFrozen     int         `json:"margin_frozen"`
			MarginAvailable  float64     `json:"margin_available"`
			ProfitUnreal     float64     `json:"profit_unreal"`
			LiquidationPrice interface{} `json:"liquidation_price"`
			LeverRate        int         `json:"lever_rate"`
			AdjustFactor     float64     `json:"adjust_factor"`
			ContractType     string      `json:"contract_type"`
			Pair             string      `json:"pair"`
			BusinessType     string      `json:"business_type"`
		} `json:"futures_contract_detail"`
		MarginMode        string      `json:"margin_mode"`
		MarginAccount     string      `json:"margin_account"`
		MarginAsset       string      `json:"margin_asset"`
		MarginBalance     float64     `json:"margin_balance"`
		MarginStatic      float64     `json:"margin_static"`
		MarginPosition    int         `json:"margin_position"`
		MarginFrozen      int         `json:"margin_frozen"`
		ProfitReal        float64     `json:"profit_real"`
		ProfitUnreal      int         `json:"profit_unreal"`
		WithdrawAvailable float64     `json:"withdraw_available"`
		RiskRate          interface{} `json:"risk_rate"`
		PositionMode      string      `json:"position_mode"`
		ContractDetail    []struct {
			Symbol           string      `json:"symbol"`
			ContractCode     string      `json:"contract_code"`
			MarginPosition   int         `json:"margin_position"`
			MarginFrozen     int         `json:"margin_frozen"`
			MarginAvailable  float64     `json:"margin_available"`
			ProfitUnreal     float64     `json:"profit_unreal"`
			LiquidationPrice interface{} `json:"liquidation_price"`
			LeverRate        int         `json:"lever_rate"`
			AdjustFactor     float64     `json:"adjust_factor"`
			ContractType     string      `json:"contract_type"`
			Pair             string      `json:"pair"`
			BusinessType     string      `json:"business_type"`
		} `json:"contract_detail"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
