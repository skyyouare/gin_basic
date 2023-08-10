package model

type Auth struct {
	ID       uint   `gorm:"column:id;primary_key"` //primary_key:设置主键
	Username string `gorm:"column:username;type:varchar(100)"`
	Password string `gorm:"column:password;type:varchar(100)"`
	// CreatedAt time.Time `gorm:"column:created_at"`
	// UpdatedAt time.Time `gorm:"column:updated_at"`
}
