package tebexgo

type Session struct {
	Secret string
}

type ServerInformation struct {
	Account
	Server
}

type Account struct {
	Id         int      `json:"id"`
	Domain     string   `json:"domain"`
	Name       string   `json:"name"`
	Currency   Currency `json:"currency"`
	OnlineMode bool     `json:"online_mode"`
	GameType   string   `json:"game_type"`
	LogEvents  bool     `json:"log_events"`
	Server     Server   `json:"server"`
}

type Currency struct {
	Iso4217 string `json:"iso_4217"`
	Symbol  string `json:"symbol"`
}

type Server struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Packages struct {}