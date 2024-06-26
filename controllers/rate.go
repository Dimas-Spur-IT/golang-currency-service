package controllers

import (
	"golang-currency-service/models"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type RateController struct {
	web.Controller
}

func (c *RateController) GetAll() {
	o := orm.NewOrm()
	var rates []models.Rate
	_, err := o.QueryTable(new(models.Rate)).All(&rates)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = rates
	}
	c.ServeJSON()
}

func (c *RateController) GetByDate() {
	dateStr := c.GetString("date")
	_, err := time.Parse("2006-01-02", dateStr)

	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = "Invalid date format"
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	var rates []models.Rate
	_, err = o.QueryTable(new(models.Rate)).Filter("Date", dateStr).All(&rates)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = rates
	}
	c.ServeJSON()
}
