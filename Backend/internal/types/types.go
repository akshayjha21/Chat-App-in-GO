package types

type User struct{
	Id uint `json:"Id" gorm:"primaryKey"`
	Name string `json:"Name" validate:"required"`
	Email string `json:"Email"`
}

type Room struct{
	Id uint `json:"Id"`
	Name string `json:"Name"`
}


type RoomMember struct{
	RoomId uint32 `json:"RoomID"`
	UserId uint32 `json:"UserID"`
	Role string
}