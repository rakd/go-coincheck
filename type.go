package coincheck

import "github.com/shopspring/decimal"

// TickerResponse ...
type TickerResponse struct {
	/*
		Last   float64 `json:"last"`
		Bid    float64 `json:"bid"`
		Ask    float64 `json:"ask"`
		High   float64 `json:"high"`
		Low    float64 `json:"low"`
		Volume float64 `json:"volume"`
	*/
	Last   decimal.Decimal `json:"last"`
	Bid    decimal.Decimal `json:"bid"`
	Ask    decimal.Decimal `json:"ask"`
	High   decimal.Decimal `json:"high"`
	Low    decimal.Decimal `json:"low"`
	Volume decimal.Decimal `json:"volume"`

	Timestamp int64 `json:"timestamp"`
}

// BalanceResponse ..
/*{"success":true,"jpy":"338648.25501917","btc":"0.97009609","eth":"0","etc":"50.0","dao":"0","lsk":"1.0","fct":"1.0","xmr":"0.05","rep":"0.1","xrp":"310.0","zec":"0.05","xem":"10.0","ltc":"0","dash":"0.01","bch":"0.01","jpy_reserved":"0.0","btc_reserved":"0.0","eth_reserved":"0","etc_reserved":"0","dao_reserved":"0","lsk_reserved":"0","fct_reserved":"0","xmr_reserved":"0","rep_reserved":"0","xrp_reserved":"0","zec_reserved":"0","xem_reserved":"0","ltc_reserved":"0","dash_reserved":"0","bch_reserved":"0","jpy_lend_in_use":"0.0","btc_lend_in_use":"0.0","eth_lend_in_use":"0.0","etc_lend_in_use":"0.0","dao_lend_in_use":"0.0","lsk_lend_in_use":"0.0","fct_lend_in_use":"0.0","xmr_lend_in_use":"0.0","rep_lend_in_use":"0.0","xrp_lend_in_use":"0.0","zec_lend_in_use":"0.0","xem_lend_in_use":"0.0","ltc_lend_in_use":"0.0","dash_lend_in_use":"0.0","bch_lend_in_use":"0.0","jpy_lent":"0.0","btc_lent":"0.0","eth_lent":"0.0","etc_lent":"0.0","dao_lent":"0.0","lsk_lent":"0.0","fct_lent":"0.0","xmr_lent":"0.0","rep_lent":"0.0","xrp_lent":"0.0","zec_lent":"0.0","xem_lent":"0.0","ltc_lent":"0.0","dash_lent":"0.0","bch_lent":"0.0","jpy_debt":"0.0","btc_debt":"0.0","eth_debt":"0.0","etc_debt":"0.0","dao_debt":"0.0","lsk_debt":"0.0","fct_debt":"0.0","xmr_debt":"0.0","rep_debt":"0.0","xrp_debt":"0.0","zec_debt":"0.0","xem_debt":"0.0","ltc_debt":"0.0","dash_debt":"0.0","bch_debt":"0.0"}*/
type BalanceResponse struct {
	Success bool `json:"success"`

	Jpy  decimal.Decimal `json:"jpy"`
	Eth  decimal.Decimal `json:"eth"`
	Etc  decimal.Decimal `json:"etc"`
	Btc  decimal.Decimal `json:"btc"`
	Dao  decimal.Decimal `json:"dao"`
	Fct  decimal.Decimal `json:"fct"`
	Lsk  decimal.Decimal `json:"lsk"`
	Xmr  decimal.Decimal `json:"xmr"`
	Rep  decimal.Decimal `json:"rep"`
	Xrp  decimal.Decimal `json:"xrp"`
	Xem  decimal.Decimal `json:"xem"`
	Zec  decimal.Decimal `json:"zec"`
	Bch  decimal.Decimal `json:"bch"`
	Dash decimal.Decimal `json:"dash"`

	JpyReserved decimal.Decimal `json:"jpy_reserved"`
	EthReserved decimal.Decimal `json:"eth_reserved"`
	EtcReserved decimal.Decimal `json:"etc_reserved"`
	BtcReserved decimal.Decimal `json:"btc_reserved"`
	DaoReserved decimal.Decimal `json:"dao_reserved"`

	LskReserved  decimal.Decimal `json:"lsk_reserved"`
	FctReserved  decimal.Decimal `json:"fct_reserved"`
	XmrReserved  decimal.Decimal `json:"xmr_reserved"`
	RepReserved  decimal.Decimal `json:"rep_reserved"`
	XrpReserved  decimal.Decimal `json:"xrp_reserved"`
	ZecReserved  decimal.Decimal `json:"zec_reserved"`
	XemReserved  decimal.Decimal `json:"xem_reserved"`
	LtcReserved  decimal.Decimal `json:"ltc_reserved"`
	DashReserved decimal.Decimal `json:"dash_reserved"`
	BchReserved  decimal.Decimal `json:"bch_reserved"`

	JpyLendInUse  decimal.Decimal `json:"jpy_lend_in_use"`
	BtcLendInUse  decimal.Decimal `json:"btc_lend_in_use"`
	EthLendInUse  decimal.Decimal `json:"eth_lend_in_use"`
	EtcLendInUse  decimal.Decimal `json:"etc_lend_in_use"`
	DaoLendInUse  decimal.Decimal `json:"dao_lend_in_use"`
	LskLendInUse  decimal.Decimal `json:"lsk_lend_in_use"`
	FctLendInUse  decimal.Decimal `json:"fct_lend_in_use"`
	XmrLendInUse  decimal.Decimal `json:"xmr_lend_in_use"`
	RepLendInUse  decimal.Decimal `json:"rep_lend_in_use"`
	XrpLendInUse  decimal.Decimal `json:"xrp_lend_in_use"`
	ZecLendInUse  decimal.Decimal `json:"zec_lend_in_use"`
	XemLendInUse  decimal.Decimal `json:"xem_lend_in_use"`
	LtcLendInUse  decimal.Decimal `json:"ltc_lend_in_use"`
	DashLendInUse decimal.Decimal `json:"dash_lend_in_use"`
	BchLendInUse  decimal.Decimal `json:"bch_lend_in_use"`

	JpyLent  decimal.Decimal `json:"jpy_lent"`
	BtcLent  decimal.Decimal `json:"btc_lent"`
	EthLent  decimal.Decimal `json:"eth_lent"`
	EtcLent  decimal.Decimal `json:"etc_lent"`
	DaoLent  decimal.Decimal `json:"dao_lent"`
	LskLent  decimal.Decimal `json:"lsk_lent"`
	FctLent  decimal.Decimal `json:"fct_lent"`
	XmrLent  decimal.Decimal `json:"xmr_lent"`
	RepLent  decimal.Decimal `json:"rep_lent"`
	XrpLent  decimal.Decimal `json:"xrp_lent"`
	ZecLent  decimal.Decimal `json:"zec_lent"`
	XemLent  decimal.Decimal `json:"xem_lent"`
	LtcLent  decimal.Decimal `json:"ltc_lent"`
	DashLent decimal.Decimal `json:"dash_lent"`
	BchLent  decimal.Decimal `json:"bch_lent"`

	JpyDebt  decimal.Decimal `json:"jpy_debt"`
	BtcDebt  decimal.Decimal `json:"btc_debt"`
	EthDebt  decimal.Decimal `json:"eth_debt"`
	EtcDebt  decimal.Decimal `json:"etc_debt"`
	DaoDebt  decimal.Decimal `json:"dao_debt"`
	LskDebt  decimal.Decimal `json:"lsk_debt"`
	FctDebt  decimal.Decimal `json:"fct_debt"`
	XmrDebt  decimal.Decimal `json:"xmr_debt"`
	RepDebt  decimal.Decimal `json:"rep_debt"`
	XrpDebt  decimal.Decimal `json:"xrp_debt"`
	ZecDebt  decimal.Decimal `json:"zec_debt"`
	XemDebt  decimal.Decimal `json:"xem_debt"`
	LtcDebt  decimal.Decimal `json:"btc_debt"`
	DashDebt decimal.Decimal `json:"dash_debt"`
	BchDebt  decimal.Decimal `json:"bch_debt"`

	/*
		Jpy  string `json:"jpy"`
		Eth  string `json:"eth"`
		Btc  string `json:"btc"`
		Dao  string `json:"dao"`
		Fct  string `json:"fct"`
		Lsk  string `json:"lsk"`
		Xmr  string `json:"xmr"`
		Rep  string `json:"rep"`
		Xrp  string `json:"xrp"`
		Xem  string `json:"xem"`
		Zec  string `json:"zec"`
		Bch  string `json:"bch"`
		Dash string `json:"dash"`
	*/

	/*
		Jpy          decimal.Decimal `json:"jpy"`
		Eth          decimal.Decimal `json:"eth"`
		Btc          decimal.Decimal `json:"btc"`
		JpyReserved  decimal.Decimal `json:"jpy_reserved"`
		BtcReserved  decimal.Decimal `json:"btc_reserved"`
		JpyLendInUse decimal.Decimal `json:"jpy_lend_in_use"`
		BtcLendInUse decimal.Decimal `json:"btc_lend_in_use"`

		JpyLent decimal.Decimal `json:"jpy_lent"`
		BtcLent decimal.Decimal `json:"btc_lent"`
		JpyDebt decimal.Decimal `json:"jpy_debt"`
		BtcDebt decimal.Decimal `json:"btc_debt"`
	*/
}

// ErrorResponse ..
/*{"success":false,"error":"Nonce must be incremented"}
 */
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
