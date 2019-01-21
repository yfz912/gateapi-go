/*
 * Gate API v4
 *
 * APIv4 futures provides all sorts of futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 1.2.0
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type FundingRateRecord struct {
	// Unix timestamp in seconds
	T int64 `json:"t,omitempty"`
	// Funding rate
	R string `json:"r,omitempty"`
}
