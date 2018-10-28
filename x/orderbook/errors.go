//nolint
package orderbook

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	DefaultCodespace sdk.CodespaceType = 431

	CodeInvalidPriceRange  sdk.CodeType = 1
	CodeInvalidPriceFormat sdk.CodeType = 2
)

//----------------------------------------
// Error constructors

// Error for when a price is not in the sortable range (and thus cannot go in the orderbook)
func ErrInvalidPriceRange(codespace sdk.CodespaceType, priceRatio sdk.Dec) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidPriceRange, fmt.Sprintf("Invalid Price %d. Must be between 10^10 & 10^-10.", priceRatio))
}

// Error for when the Price units aren't in the right for an order
func ErrInvalidPriceFormat(codespace sdk.CodespaceType, price Price) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidPriceFormat, fmt.Sprintf("Invalid Price %v", price))
}
