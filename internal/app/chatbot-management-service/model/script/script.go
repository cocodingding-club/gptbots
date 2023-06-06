package script

import "gorm.io/gorm"

type Script struct {
	gorm.Model
	Script string `gorm:"type:text"`
	Role   string `gorm:"size:50;comment:'system, assistant, user'"`
}
