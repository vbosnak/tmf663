// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSHOPPINGCART = "SHOPPING_CART"

// SHOPPINGCART mapped from table <SHOPPING_CART>
type SHOPPINGCART struct {
	ID            int32     `gorm:"column:ID;primaryKey" json:"ID"`
	STARTDATETIME time.Time `gorm:"column:STARTDATETIME" json:"STARTDATETIME"`
	ENDDATETIME   time.Time `gorm:"column:ENDDATETIME" json:"ENDDATETIME"`
}

// TableName SHOPPINGCART's table name
func (*SHOPPINGCART) TableName() string {
	return TableNameSHOPPINGCART
}
