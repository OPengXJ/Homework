package student

type Student struct {
	Id        uint `gorm:"primarykey"`
	UserName  string
	UserPass  string
	StuGender string
	StuClass  string
	StuMajor  string
	StuPhone  string
	StuName   string
}
