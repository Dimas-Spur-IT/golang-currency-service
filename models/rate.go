package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Rate struct {
	Id              int     `orm:"auto"`
	CurID           int     `orm:"column(cur_id)" json:"Cur_ID"`
	Date            string  `orm:"column(date);type(date)" json:"Date"`
	CurAbbreviation string  `orm:"size(10)" json:"Cur_Abbreviation"`
	CurScale        int     `json:"Cur_Scale"`
	CurName         string  `orm:"size(100)" json:"Cur_Name"`
	CurOfficialRate float64 `orm:"digits(12);decimals(4)" json:"Cur_OfficialRate"`
}

func init() {
	orm.RegisterModel(new(Rate))
}
