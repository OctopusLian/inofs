package routers

import (
	"inofs/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 不需要登录既可请求的url
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/main/user/log_out", &controllers.LoginController{}, "get:LogOut")
	beego.Router("/change_captcha", &controllers.LoginController{}, "get:ChangeCaptcha")

	// 必须登录才可请求的url

	// 后台首页
	beego.Router("/main/index", &controllers.HomeController{})
	beego.Router("/main/index/notify", &controllers.HomeController{}, "get:NotifyList")
	beego.Router("/main/index/read_notify", &controllers.HomeController{}, "get:ReadNotify")
	beego.Router("/main/welcome", &controllers.HomeController{}, "get:Welcome")

	// user模块
	beego.Router("/main/user/list", &controllers.UserController{}, "get:List")
	beego.Router("/main/user/to_add", &controllers.UserController{}, "get:ToAdd")
	beego.Router("/main/user/do_add", &controllers.UserController{}, "post:DoAdd")
	beego.Router("/main/user/is_active", &controllers.UserController{}, "post:IsActive")
	beego.Router("/main/user/delete", &controllers.UserController{}, "get:Delete")
	beego.Router("/main/user/reset_pwd", &controllers.UserController{}, "get:ResetPassword")
	beego.Router("/main/user/to_edit", &controllers.UserController{}, "get:ToUpdate")
	beego.Router("/main/user/do_edit", &controllers.UserController{}, "post:DoUpdate")
	beego.Router("/main/user/muli_delete", &controllers.UserController{}, "post:MuliDelete")

	// auth模块
	beego.Router("/main/auth/list", &controllers.AuthController{}, "get:List")
	beego.Router("/main/auth/to_add", &controllers.AuthController{}, "get:ToAdd")
	beego.Router("/main/auth/do_add", &controllers.AuthController{}, "post:DoAdd")

	// 角色模块
	beego.Router("/main/role/list", &controllers.RoleController{}, "get:List")
	beego.Router("/main/role/to_add", &controllers.RoleController{}, "get:ToAdd")
	beego.Router("/main/role/do_add", &controllers.RoleController{}, "post:DoAdd")
	// 角色--用户
	beego.Router("/main/role/to_role_user_add", &controllers.RoleController{}, "get:ToRoleUser")
	beego.Router("/main/role/do_role_user_add", &controllers.RoleController{}, "post:DoRoleUser")

	// 角色--权限
	beego.Router("/main/role/to_role_auth_add", &controllers.RoleController{}, "get:ToRoleAuth")
	beego.Router("/main/role/get_auth_json", &controllers.RoleController{}, "get:GetAuthJson")
	beego.Router("/main/role/do_role_auth_add", &controllers.RoleController{}, "post:DoRoleAuth")

	// 个人中心
	beego.Router("/main/user/my_center", &controllers.MyCenterController{})
	beego.Router("/main/user/salary_slip", &controllers.SalarySlipController{})
	beego.Router("/main/user/salary_slip_detail", &controllers.SalarySlipController{}, "get:Detail")

	// 财务中心
	beego.Router("/main/caiwu/salary_slip_list", &controllers.CaiWuSalarySlipController{})
	beego.Router("/main/caiwu/to_salary_slip_import", &controllers.CaiWuSalarySlipController{}, "get:ToImportExcel")
	beego.Router("/main/caiwu/do_salary_slip_import", &controllers.CaiWuSalarySlipController{}, "post:DoImportExcel")

	beego.Router("/main/caiwu/echart_data_list", &controllers.CaiWuEchartDataController{})
	beego.Router("/main/caiwu/to_echart_data_import", &controllers.CaiWuEchartDataController{}, "get:ToImportExcel")
	beego.Router("/main/caiwu/do_echart_data_import", &controllers.CaiWuEchartDataController{}, "post:DoImportExcel")

	// 内容管理
	beego.Router("/main/news/category_list", &controllers.CategoryController{})
	beego.Router("/main/news/to_add_category", &controllers.CategoryController{}, "get:ToAdd")
	beego.Router("/main/news/do_add_category", &controllers.CategoryController{}, "post:DoAdd")

	beego.Router("/main/news/news_list", &controllers.NewsController{})
	beego.Router("/main/news/to_news_addt", &controllers.NewsController{}, "get:ToAdd")
	beego.Router("/main/news/do_news_addt", &controllers.NewsController{}, "post:DoAdd")
	beego.Router("/main/news/upload_img", &controllers.NewsController{}, "post:UploadImg")
	beego.Router("/main/news/to_edit", &controllers.NewsController{}, "get:ToEdit")
	beego.Router("/main/news/do_edit", &controllers.NewsController{}, "post:DoEdit")

	// 车辆管理模块
	beego.Router("/main/cars/car_brand_list", &controllers.CarBrandController{})
	beego.Router("/main/cars/to_car_brand_add", &controllers.CarBrandController{}, "get:ToAdd")
	beego.Router("/main/cars/do_car_brand_add", &controllers.CarBrandController{}, "post:DoAdd")

	beego.Router("/main/cars/cars_list", &controllers.CarsController{})
	beego.Router("/main/cars/to_cars_add", &controllers.CarsController{}, "get:ToAdd")
	beego.Router("/main/cars/do_cars_add", &controllers.CarsController{}, "post:DoAdd")

	beego.Router("/main/cars/cars_apply_list", &controllers.CarsApplyController{})
	beego.Router("/main/cars/to_cars_apply", &ccontrollersars.CarsApplyController{}, "get:ToApply")
	beego.Router("/main/cars/do_cars_apply", &controllers.CarsApplyController{}, "post:DoApply")
	beego.Router("/main/cars/my_apply", &controllers.CarsApplyController{}, "get:MyApply")
	beego.Router("/main/cars/audit_apply", &ccontrollersrs.CarsApplyController{}, "get:AuditApply")
	beego.Router("/main/cars/to_audit_apply", &controllers.CarsApplyController{}, "get:ToAuditApply")
	beego.Router("/main/cars/do_audit_apply", &controllers.CarsApplyController{}, "post:DoAuditApply")
	beego.Router("/main/cars/do_return", &controllers.CarsApplyController{}, "get:DoReturn")

	// 报表管理模块
	beego.Router("/main/echarts/caiwu", &controllers.EchartsCaiwuController{})
	beego.Router("/main/echarts/get_caiwu_data", &controllers.EchartsCaiwuController{}, "get:GetCaiWuChart")

	beego.Router("/main/echarts/business", &controllers.EchartsBusinessController{})
	beego.Router("/main/echarts/get_business_data", &controllers.EchartsBusinessController{}, "get:GetBusinessChart")

	beego.Router("/main/echarts/course", &controllers.EchartsCourseController{})
	beego.Router("/main/echarts/get_course_data", &controllers.EchartsCourseController{}, "get:GetCourseChart")
}
