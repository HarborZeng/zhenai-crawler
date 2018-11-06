package model

type detailedInfo struct {
	Birthplace    string //籍贯
	Nationality   string //民族
	FigureType    string //体型
	Smoking       string //抽烟情况
	Drinking      string //饮酒情况
	Housing       string //居住情况
	Caring        string //车子情况
	Child         string //小孩情况
	WillGiveBirth string //是否想要孩子
	WhenToMarry   string //何时结婚
}

type basicInfo struct {
	Nickname      string
	Avatar        string
	Gender        string
	MaritalStatus string //婚姻状态
	Sigh          string //星座
	Age           int
	Height        int    //身高 cm
	Weight        int    //体重 kg
	WorkPlace     string //工作地
	Income        string
	Occupation    string //职位
	Education     string
	IsRealName    bool
	IsVIP         bool
}

type hobby struct {
	Dish             string //菜
	Celebrity        string //名人
	Song             string //歌曲
	Book             string //书籍
	OutstandingHobby string //与众不同的爱好
}

type Profile struct {
	Id           int64
	Introduction string
	Basis        basicInfo
	Detail       detailedInfo
	Hobby        hobby
}
