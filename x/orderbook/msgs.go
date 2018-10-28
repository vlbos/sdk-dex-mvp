package orderbook

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgMakeOrder struct {
	MakerAddr      sdk.AccAddress
	SellCoins      sdk.Coin
	BuyDenom       string
	Price          sdk.Dec
	ExpirationTime time.Time
}

func NewMsgMakeOrder(makerAddr sdk.AccAddress, sellCoins sdk.Coin, buyDenom string, price sdk.Dec, expirationTime time.Time) MsgMakeOrder {
	return MsgMakeOrder{
		MakerAddr:      makerAddr,
		SellCoins:      sellCoins,
		BuyDenom:       buyDenom,
		Price:          price,
		ExpirationTime: expirationTime,
	}
}

// Implements Msg.
func (msg MsgMakeOrder) Route() string { return "orderbook" }
func (msg MsgMakeOrder) Type() string  { return "add_order" }

// Implements Msg.
func (msg MsgMakeOrder) ValidateBasic() sdk.Error {
	if msg.MakerAddr.Empty() {
		return sdk.ErrInvalidAddress(msg.MakerAddr.String())
	}

	if !msg.SellCoins.IsPositive() {
		return sdk.ErrInvalidCoins(msg.SellCoins.String())
	}

	if len(msg.BuyDenom) == 0 {
		return sdk.ErrInvalidCoins(msg.BuyDenom)
	}

	if ValidPriceRatio(msg.Price) {
		return sdk.ErrInternal(msg.Price.String())
	}

	return nil
}

// Implements Msg.
func (msg MsgMakeOrder) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// Implements Msg.
func (msg MsgMakeOrder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.MakerAddr}
}

type MsgRemoveOrder struct {
	MakerAddr sdk.AccAddress
	OrderID   int64
}

func NewMsgRemoveOrder(makerAddr sdk.AccAddress, orderID int64) MsgRemoveOrder {
	return MsgRemoveOrder{
		MakerAddr: makerAddr,
		OrderID:   orderID,
	}
}

// Implements Msg.
func (msg MsgRemoveOrder) Route() string { return "orderbook" }
func (msg MsgRemoveOrder) Type() string  { return "remove_order" }

// Implements Msg.
func (msg MsgRemoveOrder) ValidateBasic() sdk.Error {
	if msg.MakerAddr.Empty() {
		return sdk.ErrInvalidAddress(msg.MakerAddr.String())
	}

	if msg.OrderID < 0 {
		return sdk.ErrInternal(fmt.Sprintf("%d", msg.OrderID))
	}

	return nil
}

// Implements Msg.
func (msg MsgRemoveOrder) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// Implements Msg.
func (msg MsgRemoveOrder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.MakerAddr}
}