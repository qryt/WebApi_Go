package Model

type Test struct {
	Id int64 `gorm:"column:Id;primaryKey"`
	Name string `gorm:"column:Name;size:255"`
}
