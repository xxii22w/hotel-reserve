package api

import (
	"hotel-reservation/types"

	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	// 携带user
	user,ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return ErrUnAuthorized()
	}
	if !user.IsAdmin {
		return ErrUnAuthorized()
	}
	return c.Next()
}