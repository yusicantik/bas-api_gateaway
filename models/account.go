package models

type Account struct {
	AccountID string `gorm:"primaryKey"`
	Name      string
	Username  string
	Password  string
}

func (a *Account) TableName() string {
	return "account"
}
