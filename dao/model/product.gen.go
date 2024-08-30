// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePRODUCT = "PRODUCT"

// PRODUCT mapped from table <PRODUCT>
type PRODUCT struct {
	ID                    int32     `gorm:"column:ID;primaryKey;autoIncrement:true" json:"ID"`
	PRODUCTORDERID        int32     `gorm:"column:PRODUCTORDER_ID" json:"PRODUCTORDER_ID"`
	DESCRIPTION           string    `gorm:"column:DESCRIPTION" json:"DESCRIPTION"`
	ISBUNDLE              bool      `gorm:"column:IS_BUNDLE" json:"IS_BUNDLE"`
	ISCUSTOMERVISIBLE     bool      `gorm:"column:IS_CUSTOMERVISIBLE" json:"IS_CUSTOMERVISIBLE"`
	NAME                  string    `gorm:"column:NAME" json:"NAME"`
	ORDERDATE             time.Time `gorm:"column:ORDERDATE" json:"ORDERDATE"`
	PRODUCTSERIALNUMBER   string    `gorm:"column:PRODUCTSERIALNUMBER" json:"PRODUCTSERIALNUMBER"`
	STARTDATE             time.Time `gorm:"column:STARTDATE" json:"STARTDATE"`
	TERMINATIONDATE       time.Time `gorm:"column:TERMINATIONDATE" json:"TERMINATIONDATE"`
	PRODUCTCHARACTERISTIC string    `gorm:"column:PRODUCTCHARACTERISTIC" json:"PRODUCTCHARACTERISTIC"`
	PRODUCTOFFERING       string    `gorm:"column:PRODUCTOFFERING" json:"PRODUCTOFFERING"`
	PRODUCTPRICE          float32   `gorm:"column:PRODUCTPRICE" json:"PRODUCTPRICE"`
	PRODUCTTERM           int32     `gorm:"column:PRODUCTTERM" json:"PRODUCTTERM"`
	PRODUCTSPECIFICATION  string    `gorm:"column:PRODUCTSPECIFICATION" json:"PRODUCTSPECIFICATION"`
	RELATEDPARTY          string    `gorm:"column:RELATEDPARTY" json:"RELATEDPARTY"`
}

// TableName PRODUCT's table name
func (*PRODUCT) TableName() string {
	return TableNamePRODUCT
}
