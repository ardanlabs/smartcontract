package currency

// Status represents the information about the call.
type Status struct {
	TimeStamp    string `json:"timestamp"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Elapsed      int    `json:"elapsed"`
	CreditCount  int    `json:"credit_count"`
}

// DataETH represents the data provided when asking for ETH.
type DataETH struct {
	Symbol      string `json:"symbol"`
	Amount      int    `json:"amount"`
	LastUpdated string `json:"last_updated"`
	Quote       struct {
		ETH struct {
			Price       float64 `json:"price"`
			LastUpdated string  `json:"last_updated"`
		} `json:"ETH"`
	} `json:"quote"`
}

// DataUSD represents the data provided when asking for USD.
type DataUSD struct {
	Symbol      string `json:"symbol"`
	Amount      int    `json:"amount"`
	LastUpdated string `json:"last_updated"`
	Quote       struct {
		USD struct {
			Price       float64 `json:"price"`
			LastUpdated string  `json:"last_updated"`
		} `json:"USD"`
	} `json:"quote"`
}

// ResponseETH2USD represents the complete response when converting ETH to USD.
type ResponseETH2USD struct {
	Status Status    `json:"status"`
	Data   []DataUSD `json:"data"`
}

// ResponseUSD2ETH represents the complete response when converting USD to ETH.
type ResponseUSD2ETH struct {
	Status Status    `json:"status"`
	Data   []DataETH `json:"data"`
}

// TransactionDetails provides details about a transaction and it's cost.
type TransactionDetails struct {
	Hash              string
	Nonce             uint64
	GasLimit          uint64
	GasOfferPriceGWei string
	Value             string
	MaxGasPriceGWei   string
	MaxGasPriceUSD    string
}

// ReceiptDetails provides details about a receipt and it's cost.
type ReceiptDetails struct {
	Status        uint64
	GasUsed       uint64
	GasPriceGWei  string
	GasPriceUSD   string
	FinalCostGWei string
	FinalCostUSD  string
}

// BalanceDiff performs calculations on the starting and ending balance.
type BalanceDiff struct {
	BeforeGWei string
	AfterGWei  string
	DiffGWei   string
	DiffUSD    string
}

// LogData represents data we can pull from events in the receipt logs.
type LogData struct {
	EventName string
	Data      map[string]interface{}
}
