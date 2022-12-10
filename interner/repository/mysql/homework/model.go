package homework

import "gorm.io/gorm"

type Homework struct {
	gorm.Model
	Classname string
	Content   string
	StartTime string
	Deadline  string
	Workname  string
}
