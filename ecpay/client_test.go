package ecpay

import (
	"log"
	"strings"
	"testing"

	"github.com/toastcheng/ecpay-sdk-go/ecpay/order"
	"github.com/toastcheng/ecpay-sdk-go/ecpay/payment"
	"github.com/toastcheng/ecpay-sdk-go/ecpay/trade"
)

func getTestClient() *Client {
	client, err := NewClient("2000132", "5294y06JbISpM5x9", "v77hoKGq4kWxNNIS", WithSandbox)
	if err != nil {
		log.Fatalf("failed to new client: %v", err)
	}
	return client
}

func TestCreateOrderAll(t *testing.T) {
	client := getTestClient()
	order := order.Order{
		MerchantTradeNo:   "NO123",
		MerchantTradeDate: "2020/10/10 10:10:10",
		PlatformID:        "3002599",
		ChoosePayment:     order.ChoosePaymentTypeAll,
		TotalAmount:       100,
		PaymentType:       order.PaymentTypeAIO,
		ItemNames:         []string{"item1", "item2"},
		TradeDesc:         "description",
		ReturnURL:         "https://abc.com",
		NeedExtraPaidInfo: false,
		Invoice: &order.InvoiceParam{
			CustomerEmail: "abc@gmail.com",
		},
		Credit: &order.CreditParam{},
	}

	resp, err := client.AioCheckOut(order)
	if err != nil {
		log.Fatalf("failed to AioCheckOut: %v", err)
	}
	if strings.Contains(resp, "交易失敗 Transaction failed") {
		log.Fatalf("failed to AioCheckOut: %s", resp)
	}
}

func TestQueryTradeInfo(t *testing.T) {
	client := getTestClient()
	info := trade.Info{
		MerchantTradeNo: "NO123",
		TimeStamp:       "2020/10/10 10:10:10",
	}

	_, err := client.QueryTradeInfo(info)
	if err != nil {
		log.Fatalf("failed to QueryTradeInfo: %v", err)
	}
}

func TestQueryTrade(t *testing.T) {
	client := getTestClient()
	trade := trade.Trade{
		CreditRefundID: "NO123",
	}

	_, err := client.QueryTrade(trade)
	if err != nil {
		log.Fatalf("failed to QueryTradeInfo: %v", err)
	}
}

func TestQueryPaymentInfo(t *testing.T) {
	client := getTestClient()
	info := payment.Info{
		MerchantTradeNo: "NO123",
		TimeStamp:       "2020/10/10 10:10:10",
	}

	_, err := client.QueryPaymentInfo(info)
	if err != nil {
		log.Fatalf("failed to QueryTradeInfo: %v", err)
	}
}

func TestQueryCreditCardPeriodInfo(t *testing.T) {
	client := getTestClient()
	info := payment.CreditCardPeriodInfo{
		MerchantTradeNo: "NO123",
		TimeStamp:       "2020/10/10 10:10:10",
	}

	_, err := client.QueryCreditCardPeriodInfo(info)
	if err != nil {
		log.Fatalf("failed to QueryTradeInfo: %v", err)
	}
}
