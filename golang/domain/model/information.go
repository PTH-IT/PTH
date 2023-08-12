package model

type Information struct {
	Id           int64  `gorm:"column:id"`
	User_id      int64  `gorm:"column:user_id"`
	Birth_day    string `gorm:"column:birth_day"`
	Number_phone string `gorm:"column:number_phone"`
	Link         string `gorm:"column:link"`
	Address      string `gorm:"column:address"`
	Education    string `gorm:"column:education"`
	Work         string `gorm:"column:work"`
	Project      string `gorm:"column:project"`
	Skill        string `gorm:"column:skill"`
}
type EducationInformation struct {
	Name   string
	From   string
	To     string
	Gpa    float64
	Majors string
}
type WorkInformation struct {
	Company  string
	From     string
	To       string
	Position string
	Details  string
}

type ProjectInformation struct {
	Company  string
	From     string
	To       string
	Position string
	Details  string
}

type SkillInformation struct {
	Details string
	scores  int64
}
