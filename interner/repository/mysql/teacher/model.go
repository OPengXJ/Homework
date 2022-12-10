package teacher

type Teacher struct {
	Id        uint `gorm:"primarykey"`
	UserName  string
	UserPass  string
	TeaGender string
	TeaPhone  string
	TeaName   string
}
