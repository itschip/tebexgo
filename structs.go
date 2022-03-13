package tebexgo

type Session struct {
	Secret string
}

// Server information

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

// Packages

type Package struct {
	Id                int             `json:"id"`
	Name              string          `json:"name"`
	Image             string          `json:"image"`
	Price             int             `json:"price"`
	ExpiryLength      int             `json:"expiry_length"`
	ExpiryPeriod      string          `json:"expiry_period"`
	Type              string          `json:"type"`
	Category          PackageCategory `json:"category"`
	GlobalLimit       int             `json:"global_limit"`
	GlobalLimitPeriod string          `json:"global_limit_period"`
	UserLimit         int             `json:"user_limit"`
	Servers           []PackageServer `json:"servers"`
	RequiredPackages  []interface{}   `json:"required_packages"`
	RequireAny        bool            `json:"require_any"`
	CreateGiftcard    bool            `json:"create_giftcard"`
	ShowUntil         bool            `json:"show_until"`
	GuiItem           string          `json:"gui_item"`
	Disabled          bool            `json:"disabled"`
	DisableQuantity   bool            `json:"disable_quantity"`
	CustomPrice       bool            `json:"custom_price"`
	ChooseServer      bool            `json:"choose_server"`
	LimitExpires      bool            `json:"limit_expires"`
	InheritCommands   bool            `json:"inherit_commands"`
	VariableGiftcard  bool            `json:"variable_giftcard"`
}

type PackageCategory struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type PackageServer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UpdatePackageObject struct {
	Disabled bool
	Name     string
	Price    int
}

// Listing

type Listing struct {
	Categories []ListingCategory `json:"categories"`
}

type ListingCategory struct {
	Id                int            `json:"id"`
	Order             int            `json:"order"`
	Name              string         `json:"name"`
	OnlySubcategories bool           `json:"only_subcategories"`
	Subcategories     []interface{}  `json:"subcategories"`
	Packages          ListingPackage `json:"packages"`
}

type ListingPackage struct {
	Id    int                `json:"id"`
	Order int                `json:"order"`
	Name  string             `json:"name"`
	Price string             `json:"price"`
	Sale  ListingPackageSale `json:"sale"`
}

type ListingPackageSale struct {
	Active   bool   `json:"active"`
	Discount string `json:"discount"`
}

// Bans

type Bans struct {
	Data []BanData `json:"data"`
}

type BanData struct {
	Id           int     `json:"id"`
	Time         string  `json:"time"`
	Ip           string  `json:"ip"`
	PaymentEmail string  `json:"payment_email"`
	Reason       string  `json:"reason"`
	User         BanUser `json:"user"`
}

type BanUser struct {
	Ign  string `json:"ign"`
	Uuid string `json:"uuid"`
}
