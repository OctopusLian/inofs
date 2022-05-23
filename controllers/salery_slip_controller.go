package controllers

import (
	"inofs/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/beego/beego/orm"
)

type SalarySlipController struct {
	beego.Controller
}

func (s *SalarySlipController) Get() {

	month := s.GetString("month")
	if month == "" {
		month = time.Now().Format("2006-01")
	}
	id := s.GetSession("id")
	o := orm.NewOrm()
	user := models.User{}
	o.QueryTable("sys_user").Filter("id", id).One(&user)
	card_id := user.CardId
	salary_slip := models.SalarySlip{}

	o.QueryTable("sys_salary_slip").Filter("card_id", card_id).Filter("pay_date", month).One(&salary_slip)
	s.Data["salary_slip"] = salary_slip
	s.TplName = "user/salary_slip_list.html"
}

func (s *SalarySlipController) Detail() {
	id, _ := s.GetInt("id")
	o := orm.NewOrm()
	salary_slip := models.SalarySlip{}
	o.QueryTable("sys_salary_slip").Filter("id", id).One(&salary_slip)
	s.Data["salary_slip"] = salary_slip
	s.TplName = "user/salary_slip_detail.html"

}
