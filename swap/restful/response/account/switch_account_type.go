package account

type SwitchAccountType struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AccountType int `json:"account_type"`
	} `json:"data"`
	Ts int64 `json:"ts"`
}
