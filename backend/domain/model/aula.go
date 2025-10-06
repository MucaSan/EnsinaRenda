package model

type Aula struct {
	ID        int    `gorm:"primaryKey;column:id"`
	Nome      string `gorm:"column:nome"`
	Descricao string `gorm:"default:false;column:descricao"`
}

func (Aula) TableName() string {
	return "Aula"
}
