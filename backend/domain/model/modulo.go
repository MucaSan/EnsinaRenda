package model

type Modulo struct {
	ID        int    `gorm:"primaryKey;column:id"`
	Nome      string `gorm:"column:nome"`
	Descricao string `gorm:"default:false;column:descricao"`
}

func (Modulo) TableName() string {
	return "modulo"
}
