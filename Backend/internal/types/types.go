package types

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null" valiDate:"required"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	Rooms     []Room    `json:"rooms,omitempty" gorm:"many2many:room_members;"`
}
type Room struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" gorm:"unique;not null"`
	RoomCode  string `json:"roomcode" gorm:"unique;not null"`
	Isprivate bool   `json:"is_private" gorm:"default:false"`
	//many users can be in many room
	Members []User `json:"members" gorm:"many2many:room_members;"`
}
type RoomMember struct {
	RoomID uint   `json:"roomid" gorm:"primaryKey"`
	UserID uint   `json:"userid" gorm:"primaryKey"`
	Role   string `json:"role" gorm:"default:'member'"`
}

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	
	// Sender
	FromID    uint      `json:"from_id" gorm:"index"`
	FromUser  User      `json:"from_user" gorm:"foreignKey:FromID"`
	
	// Receiver (Used for 1-to-1)
	ToID      *uint     `json:"to_id,omitempty" gorm:"index"`
	ToUser    *User     `json:"to_user,omitempty" gorm:"foreignKey:ToID"`
	
	// Room (Used for Group Chat)
	RoomID    *uint     `json:"room_id,omitempty" gorm:"index"`
	
	CreatedAt time.Time `json:"created_at"`
}