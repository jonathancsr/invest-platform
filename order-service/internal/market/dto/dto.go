package dto

type TradeInput struct {
	OrderId       string  `json:"order_id"`
	InvestorID    string  `json:"investor_id"`
	AssetID       string  `json:"asset_id"`
	CurrentShares int     `json:"current_shares"`
	Shares        int     `json:"shares"`
	Price         float64 `json:"price"`
	OrderType     string  `json:"order_type"`
}

type OrderOutput struct {
	OrderId           string               `json:"order_id"`
	InvestorID        string               `json:"investor_id"`
	AssetID           string               `json:"asset_id"`
	OrderType         string               `json:"order_type"`
	Status            string               `json:"status"`
	Partial           int                  `json:"partial"`
	Shares            int                  `json:"shares"`
	Order             int                  `json:"orders"`
	TransactionOutput []*TransactionOutput `json:"transactions"`
}

type TransactionOutput struct {
	TransactionID string  `json:"transaction_id"`
	BuyerId       string  `json:"buyer_id"`
	SellerID      string  `json:"seller_id"`
	AssetID       string  `json:"asset_id"`
	Price         float64 `json:"price"`
	Shares        int     `json:"shares"`
}
