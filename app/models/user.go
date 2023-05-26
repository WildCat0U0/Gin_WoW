package models

//定义 User 模型

type User struct {
	ID
	Name     string `gorm:"not null;comment:用户名称" json:"name"`
	Password string `gorm:"not null;comment:用户密码" json:"password"`
	Mobile   string `gorm:"not null;index;comment:用户手机号" json:"mobile"`
	Email    string `gorm:"not null;index;comment:用户邮箱" json:"email"` // index 表示 Email 字段在数据库中创建索引
	Timestamps
	SoftDeletes
}
