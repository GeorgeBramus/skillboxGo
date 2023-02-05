package model

type User struct {
	Id   uint64 `gorm:"primaryKey"`
	Name string `gorm:"index"`
	Age  uint8
}

type Friends struct {
	Id       uint64 `gorm:"primaryKey"`
	UserId   uint64 `json:"source_id,omitempty"`
	FriendId uint64 `json:"target_id,omitempty"`
}
