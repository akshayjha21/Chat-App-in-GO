package storage

import "github.com/akshayjha21/Chat-App-in-GO/Backend/internal/types"

// import "os/user"
type Storage interface {
	RegisterUser(user *types.User)(*types.User,error)
	CreateConnection(user *types.User,room *types.Room) (*types.RoomMember,error)
}