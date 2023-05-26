package config

type Jwt struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`   // jwt加密值 mapstrcture 是 gorm.io/gorm 的一个库，用于将结构体与数据库进行映射
	JwtTtl int64  `mapstructure:"jwt_ttl" json:"jwtTtl" yaml:"jwt_ttl"` // jwt过期时间
}
