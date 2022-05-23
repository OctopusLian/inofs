package models

import "github.com/beego/beego/orm"

func init() {
	orm.RegisterModel(
		new(User),
		new(Auth),
		new(Role),
		new(User),
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
