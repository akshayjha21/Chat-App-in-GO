package user

import (
	// "fmt"
	// "log/slog"

	"errors"

	"github.com/akshayjha21/Chat-App-in-GO/Backend/internal/storage/postgres"
	"github.com/akshayjha21/Chat-App-in-GO/Backend/internal/types"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Handler struct {
	DB *postgres.Postgres
}

type data struct {
	Id   uint   `json:"Id"`
	Name string `json:"Username"`
}
type response struct {
	Status  bool   `json:"Status"`
	Message string `json:"Message"`
	Data    *data  `json:"Data,omitempty"`
}

func (h *Handler) Registerhandler(c *fiber.Ctx) error {
	user := new(types.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  false,
			Message: "Invalid request body",
		})
	}
	// fmt.Println(user)

	// Make sure these field names match your types.User struct!
	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  false,
			Message: "Name and Password are required",
		})
	}

	res := register(h.DB, user)

	if !res.Status {
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	return c.Status(fiber.StatusCreated).JSON(res)
}
func (h *Handler) LoginHandler(c *fiber.Ctx) error {
	user := &types.User{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  false,
			Message: "Invalid request body",
		})
	}
	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response{
			Status:  false,
			Message: "Name and Password are required",
		})
	}
	res := login(h.DB, user)
	if !res.Status {
		// Use 401 for bad credentials, 500 for actual DB crashes
		if res.Message == "Invalid credentials" {
			return c.Status(fiber.StatusUnauthorized).JSON(res)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
func login(db *postgres.Postgres, user *types.User) *response {
	u, err := db.GetUser(&types.User{Username: user.Username})
	if err != nil {
		// Specifically check if the record just wasn't found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &response{Status: false, Message: "Invalid credentials"}
		}
		// Handle actual database connection errors
		return &response{Status: false, Message: "Internal server error"}
	}
	if !checkPassword(user.Password, u.Password) {
		return &response{Status: false, Message: "Invalid password"}
	}
	return &response{
		Status:  true,
		Message: "User login succesfully",
		// Using map[string]any if your response Data field supports it
		Data: &data{
			Id:   u.ID,
			Name: u.Username},
	}
}
func register(db *postgres.Postgres, user *types.User) *response {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return &response{Status: false, Message: "Error processing password"}
	}
	user.Password = string(hashPassword)

	dbUser, err := db.RegisterUser(user)
	if err != nil {
		return &response{Status: false, Message: "Error registering user in database"}
	}

	return &response{
		Status:  true,
		Message: "User registered succesfully",
		Data: &data{
			Id:   dbUser.ID,
			Name: dbUser.Username,
		},
	}
}

func checkPassword(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
