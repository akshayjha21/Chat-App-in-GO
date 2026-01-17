package handler

import (
	"github.com/akshayjha21/Chat-App-in-GO/Backend/internal/storage/postgres"
	"github.com/akshayjha21/Chat-App-in-GO/Backend/internal/types"
	"github.com/gofiber/fiber/v2"
)

type MessageHandler struct {
	DB *postgres.Postgres
}
type MessageResponse struct{
	Message []types.Message `json:"Message"`
}
func (m *MessageHandler) GetRoomMessages(c *fiber.Ctx) error {
	
	var roomId uint
	if err := c.BodyParser(&roomId); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Response{
			Status:  false,
			Message: "Invalid Request Body",
		})
	}
	res:=getRoomMess(m.DB,roomId)
	if !res.Status {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}
	return c.Status(fiber.StatusCreated).JSON(res)
	// return c.JSON(messages)
}
func getRoomMess(db *postgres.Postgres,roomId uint)*Response{
	room_msg,err:=db.
}
func (m *MessageHandler) GetPrivateMessages(c *fiber.Ctx) error {
	userA := c.Query("user_id") // The current user
	userB := c.Params("otherId") // The person they are chatting with

	var messages []types.Message

	// Query: (from A to B AND room is null) OR (from B to A AND room is null)
	err := m.DB.Db.Preload("FromUser").
		Where("(from_id = ? AND to_id = ? AND room_id IS NULL) OR (from_id = ? AND to_id = ? AND room_id IS NULL)", 
			userA, userB, userB, userA).
		Order("created_at asc").
		Find(&messages).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(messages)
}