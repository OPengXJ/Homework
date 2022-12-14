package homework

import (
	"time"

	"gorm.io/gorm"
)

type Homework struct {
	gorm.Model
	College   string
	TeaName string
	ClassName string
	Content   string
	StartTime time.Time
	Deadline  time.Time
	WorkName  string
	Session int
}
