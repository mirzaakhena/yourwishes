package vo

import "fmt"

type WishesID string

func NewWishesID(randomID string) (WishesID, error) {

	var obj = WishesID(fmt.Sprintf("WS-%s", randomID))

	return obj, nil
}

func (r WishesID) String() string {
	return string(r)
}
