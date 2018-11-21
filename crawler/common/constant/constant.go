package constant

import "sync"

var DebugMode = false //是否开启debug模式

var ChildCondition = []string{
	"没有小孩",
	"有孩子且住在一起",
	"有孩子且偶尔会一起住",
	"有孩子但不在身边",
}

var WillGiveBirthCondition = []string{
	"视情况而定",
	"想要孩子",
	"不想要孩子",
	"以后再告诉你",
}

var MaritalStatusCondition = []string{
	"离异",
	"未婚",
	"丧偶",
}

var EducationCondition = []string{
	"不限",
	"高中及以下",
	"中专",
	"大专",
	"大学本科",
	"硕士",
	"博士",
}

var OccupationCondition = []string{
	"销售",
	"客户服务",
	"计算机/互联网",
	"通信/电子",
	"生产/制造",
	"物流/仓储",
	"商贸/采购",
	"人事/行政",
	"高级管理",
	"广告/市场",
	"传媒/艺术",
	"生物/制药",
	"医疗/护理",
	"金融/银行/保险",
	"建筑/房地产",
	"咨询/顾问",
	"法律",
	"财会/审计",
	"教育/科研",
	"服务业",
	"交通运输",
	"政府机构",
	"军人/警察",
	"农林牧渔",
	"自由职业",
	"在校学生",
	"待业",
	"其他行业",
	//TODO 还有好多没有弄进来的二级菜单里面，懒得弄了
}

var DeduplicationBoolMap = sync.Map{}
