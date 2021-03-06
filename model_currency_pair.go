/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 4.6.1
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type CurrencyPair struct {
	// Currency pair
	Id string `json:"id,omitempty"`
	// Base currency
	Base string `json:"base,omitempty"`
	// Quote currency
	Quote string `json:"quote,omitempty"`
	// Trading fee
	Fee string `json:"fee,omitempty"`
	// Minimum amount of base currency to trade, `null` means no limit
	MinBaseAmount string `json:"min_base_amount,omitempty"`
	// Minimum amount of quote currency to trade, `null` means no limit
	MinQuoteAmount string `json:"min_quote_amount,omitempty"`
	// Price scale
	Precision int32 `json:"precision,omitempty"`
}
