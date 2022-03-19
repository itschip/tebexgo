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
	Image             interface{}     `json:"image"`
	Price             float32         `json:"price"`
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
	Disabled bool   `json:"disabled"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
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

type BanInput struct {
	Reason string `json:"reason"`
	Ip     string `json:"ip"`
	User   string `json:"user"`
}

// Checkout

type PutCheckoutObject struct {
	PackageId string `json:"package_id"`
	Username  string `json:"username"`
}

type Checkout struct {
	Url     string `json:"url"`
	Expires string `json:"expires"`
}

// Gift cards

type GiftCards struct {
	Data []GiftCard `json:"data"`
}

type GiftCard struct {
	Id      int             `json:"id"`
	Code    string          `json:"code"`
	Balance GiftCardBalance `json:"balance"`
	Note    string          `json:"note"`
	Void    bool            `json:"void"`
}

type GiftCardBalance struct {
	Starting  string `json:"starting"`
	Remaining string `json:"remaining"`
	Currency  string `json:"currency"`
}

type PutGiftCardObject struct {
	ExpiresAt string `json:"expires_at"`
	Note      string `json:"note"`
	Amount    int    `json:"amount"`
}

type Sales struct {
	Data []Sale `json:"data"`
}

type Sale struct {
	Id        int           `json:"id"`
	Effective SaleEffective `json:"effective"`
	Discount  SaleDiscount  `json:"discount"`
	Start     int           `json:"start"`
	Expire    int           `json:"expire"`
	Order     int           `json:"order"`
}

type SaleEffective struct {
	Type       string `json:"package"`
	Packages   []int  `json:"packages"`
	Categories []int  `json:"categories"`
}

type SaleDiscount struct {
	Type       string `json:"type"`
	Percentage int    `json:"percentage"`
	Value      int    `json:"value"`
}

type PlayerData struct {
	Player         Player          `json:"player"`
	BanCount       int             `json:"banCount"`
	ChargebackRate int             `json:"chargebackRate"`
	Payments       []PlayerPayment `json:"payments"`
	PurchaseTotals PurchaseTotals  `json:"purchaseTotals"`
}

type Player struct {
	Id               string `json:"id"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	CacheExpire      string `json:"cache_expire"`
	Username         string `json:"username"`
	Meta             string `json:"meta"`
	PluginUsernameId string `json:"plugin_username_id"`
}

type PlayerPayment struct {
	TxnId    string `json:"txn_id"`
	Time     int    `json:"time"`
	Price    int64  `json:"price"`
	Currency string `json:"currency"`
	Status   int    `json:"status"`
}

type PurchaseTotals struct {
	USD int64 `json:"USD"`
	GBP int64 `json:"GBP"`
}

// Payments

type Payment struct {
	Id             int              `json:"id"`
	Amount         string           `json:"amount"`
	Date           string           `json:"date"`
	Currency       Currency         `json:"currency"`
	PaymentGateway PaymentGateway   `json:"gateway"`
	Status         string           `json:"status"`
	Email          string           `json:"email"`
	Player         PaymentPlayer    `json:"player"`
	Packages       []PaymentPackage `json:"packages"`
}

type PaymentGateway struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type PaymentPlayer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}

type PaymentPackage struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type PaymentPackageFields struct {
	Name    string                      `json:"name"`
	Value   string                      `json:"value"`
	Type    string                      `json:"type"`
	Options []PaymentPackageFieldOption `json:"options"`
}

type PaymentPackageFieldOption struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

type UpdatePaymentObject struct {
	Username string `json:"username"`
	Status   string `json:"status"`
}

// Community

type CommunityGoal struct {
	Id            int    `json:"id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	Account       int    `json:"account"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Image         string `json:"image"`
	Target        string `json:"target"`
	Current       string `json:"current"`
	Repeatable    int    `json:"repeatable"`
	LastAchieved  int    `json:"last_achieved"`
	TimesAchieved int    `json:"times_achieved"`
	Satus         string `json:"status"`
	Sale          int    `json:"sale"`
}

// Coupons

type Coupons struct {
	Pagination CouponPagination `json:"pagination"`
	Data       CouponData       `json:"data"`
}

type Coupon struct {
	Data CouponData `json:"data"`
}

type CouponPagination struct {
	TotalResults int    `json:"totalResults"`
	CurrentPage  int    `json:"currentPage"`
	LastPage     int    `json:"lastPage"`
	Previous     int    `json:"previous"`
	Next         string `json:"next"`
}

type CouponData struct {
	Id         int             `json:"id"`
	Code       string          `json:"code"`
	Effective  CouponEffective `json:"effective"`
	Expire     CouponExpire    `json:"expire"`
	BasketType string          `json:"basket_type"`
	StartDate  string          `json:"start_date"`
	UserLimit  string          `json:"user_limit"`
	Minimum    int             `json:"minimum"`
	Username   string          `json:"username"`
	Note       string          `json:"note"`
}

type CouponEffective struct {
	Type       string `json:"type"`
	Packages   []int  `json:"packages"`
	Categories []int  `json:"categories"`
}

type CouponExpire struct {
	ReedemUnlimited string `json:"reedem_unlimited"`
	ExpireNever     string `json:"expire_never"`
	Limit           int    `json:"limit"`
	Date            string `json:"date"`
}

type CreateCouponObject struct {
	Code                      string `json:"code"`
	EffectiveOn               string `json:'effective_on"`
	Packages                  []int  `json:"packages"`
	Categories                []int  `json:"categories"`
	DiscountType              string `json:"discount_type"`
	DiscountAmount            int    `json:"discount_amount"`
	DiscountPrecentage        int    `json:"discount_precentage"`
	ReedemUnlimited           bool   `json:"reedem_unlimited"`
	ExpireNever               bool   `json:"expire_never"`
	ExpireLimit               int    `json:"expire_limit"`
	ExpireDate                string `json:"expire_date"`
	StartDate                 string `json:"start_date"`
	BasketType                string `json:"basket_type"`
	Minium                    int    `json:"minimum"`
	DiscountApplicationMethod int    `json:"discount_application_method"`
	Username                  string `json:"username"`
	Note                      string `json:"note"`
}
