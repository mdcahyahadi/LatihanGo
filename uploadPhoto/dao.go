package uploadphoto

import "time"

type PhotoEntity struct {
	ID        int64     `gorm:"column:ID;primaryKey;autoIncrement;not null"`
	Title     string    `gorm:"column:TITLE;not null"`
	Caption   string    `gorm:"column:CAPTION"`
	URL       string    `gorm:"column:URL;not null;unique"`
	UserID    int64     `gorm:"column:USERID;not null"`
	CreatedAt time.Time `gorm:"column:CREATIONDATE"`
	UpdatedAt time.Time `gorm:"column:LASTUPDATE"`
}

func (*PhotoEntity) TableName() string {
	return "photosx"
}
