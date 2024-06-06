package model

type Account struct {
	AccountID string `gorm:"primaryKey"`
	Name      string
	Username  string `gorm:"username"`
	Password  string
}

func (a *Account) TableName() string {
	return "account"
}
