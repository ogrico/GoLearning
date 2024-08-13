package domain

type Account struct {
    AccountNumber string  `json:"account_number"`
    Balance       float64 `json:"balance"`
}

type Credit struct {
    CreditID    string  `json:"credit_id"`
    CreditLimit float64 `json:"credit_limit"`
    Balance     float64 `json:"balance"`
}

type CreditCard struct {
    CardNumber   string  `json:"card_number"`
    CardLimit    float64 `json:"card_limit"`
    CurrentDebt  float64 `json:"current_debt"`
}