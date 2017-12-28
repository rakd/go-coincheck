package coincheck

// BalanceResponse ..
/*{"success":true,"jpy":"338648.25501917","btc":"0.97009609","eth":"0","etc":"50.0","dao":"0","lsk":"1.0","fct":"1.0","xmr":"0.05","rep":"0.1","xrp":"310.0","zec":"0.05","xem":"10.0","ltc":"0","dash":"0.01","bch":"0.01","jpy_reserved":"0.0","btc_reserved":"0.0","eth_reserved":"0","etc_reserved":"0","dao_reserved":"0","lsk_reserved":"0","fct_reserved":"0","xmr_reserved":"0","rep_reserved":"0","xrp_reserved":"0","zec_reserved":"0","xem_reserved":"0","ltc_reserved":"0","dash_reserved":"0","bch_reserved":"0","jpy_lend_in_use":"0.0","btc_lend_in_use":"0.0","eth_lend_in_use":"0.0","etc_lend_in_use":"0.0","dao_lend_in_use":"0.0","lsk_lend_in_use":"0.0","fct_lend_in_use":"0.0","xmr_lend_in_use":"0.0","rep_lend_in_use":"0.0","xrp_lend_in_use":"0.0","zec_lend_in_use":"0.0","xem_lend_in_use":"0.0","ltc_lend_in_use":"0.0","dash_lend_in_use":"0.0","bch_lend_in_use":"0.0","jpy_lent":"0.0","btc_lent":"0.0","eth_lent":"0.0","etc_lent":"0.0","dao_lent":"0.0","lsk_lent":"0.0","fct_lent":"0.0","xmr_lent":"0.0","rep_lent":"0.0","xrp_lent":"0.0","zec_lent":"0.0","xem_lent":"0.0","ltc_lent":"0.0","dash_lent":"0.0","bch_lent":"0.0","jpy_debt":"0.0","btc_debt":"0.0","eth_debt":"0.0","etc_debt":"0.0","dao_debt":"0.0","lsk_debt":"0.0","fct_debt":"0.0","xmr_debt":"0.0","rep_debt":"0.0","xrp_debt":"0.0","zec_debt":"0.0","xem_debt":"0.0","ltc_debt":"0.0","dash_debt":"0.0","bch_debt":"0.0"}*/
type BalanceResponse struct {
	Success bool `json:"success"`

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
