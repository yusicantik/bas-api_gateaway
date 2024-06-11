package models

type Bank struct {
	BankCode string `gorm:"primaryKey"`
	Name     string
	Address  string
}

func (a *Bank) TableName() string {
	return "bank"
}
