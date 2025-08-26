package lead

import (
	"github.com/eswar/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	var leads []Lead
	database.DBConn.Find(&leads)
	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	result := database.DBConn.First(&lead, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Lead not found"})
	}
	return c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	database.DBConn.Create(&lead)
	return c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	result := database.DBConn.First(&lead, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Lead not found"})
	}
	database.DBConn.Delete(&lead)
	return c.JSON(fiber.Map{"message": "Lead successfully deleted", "id": id})
}
