package domain

type ValidatedOrder struct{}

type UnvalidatedOrder struct{}

type CheckProductCodeExists = func(productCode string) bool

type UnvalidatedAddress string

type CheckedAddress string

type GetCheckedAddress = func(v UnvalidatedAddress) CheckedAddress

type CheckedAddressExists = func(v UnvalidatedAddress) (CheckedAddress, error)

type ValidateOrder = func(input UnvalidatedOrder,
	checkProductCodeExistsFn CheckProductCodeExists,
	checkedAddressExistsFn CheckedAddressExists,
) (ValidatedOrder, error)

type PlaceOrderEvents struct {}

type PlaceOrder = func(input UnvalidatedOrder) PlaceOrderEvents
