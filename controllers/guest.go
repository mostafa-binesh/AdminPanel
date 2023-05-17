package controllers

import (
	D "docker/database"
	"fmt"
	"strconv"

	M "docker/models"
	U "docker/utils"

	"github.com/gofiber/fiber/v2"
)

func GuestMessages(c *fiber.Ctx) error {
	guestID, err := strconv.Atoi(c.Cookies("guest"))
	fmt.Printf("guest id: %d\n", guestID)
	if err != nil {
		panic("guest messages atoi panic")
	}
	messages := []M.GuestMessage{}
	if result := D.DB().Find(&messages, "guest_id = ?", guestID); result.Error != nil {
		return U.DBError(c, result.Error)
	}
	return c.JSON(fiber.Map{
		"data": messages,
	})
}
func GuestSendMessage(c *fiber.Ctx) error {
	if c.Cookies("guest") == "" {
		return U.ResErr(c, "لطفا ابتدا چتی ایجاد کنید")
	}
	guestID, err := strconv.Atoi(c.Cookies("guest"))
	if err != nil {
		panic("guest messages atoi panic")
	}
	payload := new(M.GuestMessageCreate)
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	if errs := U.Validate(payload); errs != nil {
		return c.Status(400).JSON(fiber.Map{"errors": errs})
	}
	fmt.Println("till here")
	if result := D.DB().Create(&M.GuestMessage{Sender: 1, Body: payload.Body, GuestChatID: guestID}); result.Error != nil {
		return U.DBError(c, result.Error)
	}
	return U.ResErr(c, "پیام جدید ثبت شد")
}

func GuestChats(c *fiber.Ctx) error {
	if c.Cookies("guest") == "" {
		return U.ResErr(c, "لطفا ابتدا چتی ایجاد کنید")
	}
	guestID, err := strconv.Atoi(c.Cookies("guest"))
	if err != nil {
		panic("guest messages atoi panic")
	}
	guestChats := M.GuestChat{}
	if result := D.DB().Preload("Messages").First(&guestChats, "id = ?", guestID); result.Error != nil {
		return U.DBError(c, result.Error)
	}
	return c.JSON(fiber.Map{
		"data": guestChats,
	})
}
func CreateGuestChat(c *fiber.Ctx) error {
	// ! parse payload
	payload := new(M.GuestChatCreate)
	if err := c.BodyParser(payload); err != nil {
		U.ResErr(c, err.Error())
	}
	// ! validate payload
	if errs := U.Validate(payload); errs != nil {
		return c.Status(400).JSON(fiber.Map{"errors": errs})
	}
	// ! create guestChat in database
	guestChat := M.GuestChat{
		// GuestID:  guestID,
		Category: payload.Category,
		Title:    payload.Title}
	if result := D.DB().Create(&guestChat); result.Error != nil {
		return U.DBError(c, result.Error)
	}
	// ! create chat and link it to the previous created chat
	// ! TODO if creating message failed, we need to delete the chat as well
	if result := D.DB().Create(&M.GuestMessage{
		// GuestID:     guestID,
		GuestChatID: guestChat.ID,
		Sender:      1,
		Body:        payload.Body,
	}); result.Error != nil {
		return U.DBError(c, result.Error)
	}
	// ! set guestID (chatID value) cookie
	c.Cookie(&fiber.Cookie{
		Name:  "guest",
		Value: strconv.Itoa(guestChat.ID), // int to string
	})
	return U.ResMessage(c, "چت ایجاد شد")
}
