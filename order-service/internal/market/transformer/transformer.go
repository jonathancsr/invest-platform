package transformer

import (
	"github.com/jonathancsr/invest-platform/go/internal/market/dto"
	"github.com/jonathancsr/invest-platform/go/internal/market/entity"
)

func TransformInput(input dto.TradeInput) *entity.Order {
	asset := entity.NewAsset(input.AssetID, input.AssetID, 1000)
	investor := entity.NewInvestor(input.InvestorID)
	order := entity.NewOrder(input.OrderId, investor, asset, input.Shares, input.Price, input.OrderType)

	if input.CurrentShares > 0 {
		assetPosition := entity.NewInvestorAssetPosition(input.AssetID, input.CurrentShares)
		investor.AddAssetPosition(assetPosition)
	}

	return order
}

func TransformOutput(order *entity.Order) *dto.OrderOutput {
	output := dto.OrderOutput{
		OrderId:    order.ID,
		InvestorID: order.Investor.ID,
		AssetID:    order.Asset.ID,
		OrderType:  order.OrderType,
		Partial:    order.PendingShares,
		Shares:     order.Shares,
		Status:     order.Status,
	}

	var transactionsOutput []*dto.TransactionOutput
	for _, t := range order.Transactions {
		transactionOutput := &dto.TransactionOutput{
			TransactionID: t.ID,
			BuyerId:       t.BuyingOrder.ID,
			SellerID:      t.SellingOrder.ID,
			AssetID:       t.SellingOrder.Asset.ID,
			Price:         t.Price,
			Shares:        t.SellingOrder.Shares - t.BuyingOrder.Shares,
		}
		transactionsOutput = append(transactionsOutput, transactionOutput)
	}

	output.TransactionOutput = transactionsOutput

	return &output
}
