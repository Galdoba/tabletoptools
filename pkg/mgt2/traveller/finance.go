package traveller

type FinanceData struct {
	Cash_On_Hand          int `json:"Cash On Hand"`
	Debt                  int `json:"Debt"`
	Pension               int `json:"Pension"`
	Living_Cost           int `json:"Living Cost"`
	Monthly_Ship_Payments int `json:"Monthly Ship Payments"`
}
