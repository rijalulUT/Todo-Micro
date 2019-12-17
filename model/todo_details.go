package model

// TodoDetail .
type TodoDetail struct {
	ID         uint64 `json:"id" gorm:"primary_key"`
	Title      string `json:"title" gorm:"column:title;size:255"`
	Content    string `json:"description" gorm:"column:description;size:255"`
	CategoryID uint64 `json:"category_id" gorm:"column:category_id"`
}
