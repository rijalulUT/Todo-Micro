package model

// TodoCategory .
type TodoCategory struct {
	ID          uint64 `json:"id" gorm:"primary_key"`
	Title       string `json:"title" gorm:"column:title;size:255"`
	Description string `json:"description" gorm:"column:description;size:255"`
	UserID      uint64 `json:"user_id" gorm:"column:user_id"`
	TodoDetail  []TodoDetail
}
