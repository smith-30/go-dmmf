package domain

type Item struct {
	Price float64
}

type EmptyCart struct {
}

func (a *EmptyCart) Items() []Item {
	return nil
}

type ActiveCart struct {
	UnpaidItems []Item
}

func (a *ActiveCart) Items() []Item {
	return nil
}

type PaidCart struct {
	PaidItems []Item
	Payment   float64
}

func (a *PaidCart) Items() []Item {
	return nil
}

type ShoppingCart interface {
	Items() []Item
}

func AddItem(cart ShoppingCart, item Item) ShoppingCart {
	switch v := cart.(type) {
	case *EmptyCart:
		return &ActiveCart{
			UnpaidItems: []Item{item},
		}
	case *ActiveCart:
		re := &ActiveCart{
			UnpaidItems: []Item{item},
		}
		re.UnpaidItems = append(re.UnpaidItems, v.UnpaidItems...)
		return re
	case *PaidCart:
	}
	return cart
}

func MakePayment(cart ShoppingCart, payment float64) ShoppingCart {
	switch v := cart.(type) {
	case *EmptyCart:
	case *ActiveCart:
		re := &PaidCart{
			Payment: payment,
		}
		re.PaidItems = append(re.PaidItems, v.UnpaidItems...)
		return re
	case *PaidCart:
	}
	return cart
}
