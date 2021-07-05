package domain

type PricedOrder struct{}

type GetProductPrice = func(productCode string) float64

type PriceOrder = func(input ValidateOrder, GetProductPriceFn GetProductPrice) (PricedOrder, error)

type HtmlString string

type OrderAcknowledgment struct {
	EmailAddress EmailAddress
	Letter       HtmlString
}

type CreateOrderAcknowledgmentLetter = func(input PricedOrder) HtmlString

type SendOrderAcknowledgment = func(input OrderAcknowledgment) (OrderAcknowledgmentSent, error)

type OrderAcknowledgmentSent struct {
	OrderID      string
	EmailAddress EmailAddress
}

type AcknowledgmentOrder = func(
	input PricedOrder,
	createOrderAcknowledgmentLetterFn CreateOrderAcknowledgmentLetter,
	sendOrderAcknowledgmentFn SendOrderAcknowledgment,
) (OrderAcknowledgmentSent, error)

// 配送用
type OrderPlaced PricedOrder

type Address string

type BillingAmount float64

// 請求用
type BillableOrderPlaced struct {
	OrderID        string
	BillingAddress Address
	AmountToBill   BillingAmount
}

func (a OrderPlaced) Event() {}

func (a BillableOrderPlaced) Event() {}

func (a OrderAcknowledgmentSent) Event() {}

type PlaceOrderEvent interface {
	Event()
}

type CreateEvents = func(input PricedOrder) []PlaceOrderEvent

type PlaceOrderWorkFlow = func(input PlaceOrder) (PlaceOrderEvents, error)
