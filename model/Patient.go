package model

import (
	"time"
)

type Patient struct {
	HosGuid  string `json:"fname" binding:"required"`
	Hn       string `json:"hn" gorm:"primaryKey"`
	Pname    string `json:"pname" binding:"required"`
	Fname    string `json:"fname" binding:"required"`
	Lname    string `json:"lname" binding:"required"`
	MobilePhoneNumber string `json:"phonnumber"`
	Birthday time.Time
	Cid string `json:"cid"`


}