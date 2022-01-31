package entity

type CustomerTable struct {
	ID string `gorm:"column:id;size:3;primaryKey"`
}

func (c *CustomerTable) TableName() string {
	return "customer_table"
}
