package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(
		new(User),
		new(Auth),
		new(Role),
		new(CarBrand),
		new(Cars),
		new(CarsApply),
		new(MessageNotify),
		new(CaiwuData),
		new(Category),
		new(News),
		new(SalarySlip),
	)
}
