package storage

import "github.com/akshayjha21/Chat-App-in-GO/Backend/internal/types"

// import "os/user"
type Storage interface {
	RegisterUser(user *types.User)(*types.User,error)
	CreateConnection(user *types.User,room *types.Room) (*types.RoomMember,error)
	GetUser(user  *types.User)(*types.User,error)
	RegisterRoom(chatRoom *types.Room)(*types.Room,error)
	GetRoom(code string) (*types.Room, error)
	GetUserRooms(userID uint) ([]types.Room, error) 
	CheckExistingMembers(userid uint, roomid uint) (*types.RoomMember, error) 

}
