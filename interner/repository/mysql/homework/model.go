package homework

import (
	"time"

	"gorm.io/gorm"
)

type Homework struct {
	gorm.Model
	College   string    `json:"college"`
	TeaName   string    `json:"teaname"`
	ClassName string    `json:"classname"`
	Content   string    `json:"contetn"`
	StartTime time.Time `json:"starttime"`
	Deadline  time.Time `json:"deadline"`
	WorkName  string    `json:"workname"`
	Session   int       `json:"session"`
	TeaId     int       `json:"teaid"`
}
