package noteHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gofiber-gorm-mysql/database"
	"gofiber-gorm-mysql/internal/model"
)

func GetNotes(ctx *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	db.Find(&notes)

	if len(notes) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "Error",
			"message": "No notes present",
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Notes found",
		"data":    notes,
	})
}

func CreateNotes(ctx *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	// Store body into the note and return error if encountered
	err := ctx.BodyParser(note)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Error",
			"message": "Review your input",
			"data":    err,
		})
	}

	//	Add a uuid to the note
	note.ID = uuid.New()
	//	Create the note and return error if encountered
	err = db.Create(&note).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Error",
			"message": "Could not create note",
			"data":    err,
		})
	}
	// Return the created note
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Created note",
		"data":    note,
	})
}

func GetNote(ctx *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	paramId := ctx.Params("noteId")

	db.Find(&note, "id = ?", paramId)

	if note.ID == uuid.Nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "Error",
			"message": "Note not found",
			"data":    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Notes found",
		"data":    note,
	})
}

func UpdateNote(ctx *fiber.Ctx) error {
	type Request struct {
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
		Text     string `json:"text"`
	}

	db := database.DB
	var note model.Note

	paramId := ctx.Params("noteId")
	db.Find(&note, "id = ?", paramId)
	if note.ID != uuid.Nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "Error",
			"message": "Note not found",
			"data":    nil,
		})
	}
	var updateNote Request
	err := ctx.BodyParser(&updateNote)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "Error",
			"message": "Review your input",
			"data":    err,
		})
	}

	note.Title = updateNote.Title
	note.Subtitle = updateNote.Subtitle
	note.Text = updateNote.Text

	db.Save(&note)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Your note updated",
		"data":    note,
	})
}

func DeleteNote(ctx *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	paramId := ctx.Params("noteId")
	db.Find(&note, "id = ?", paramId)
	if note.ID == uuid.Nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "Error",
			"message": "Note not found",
			"data":    nil,
		})
	}

	err := db.Delete(&note, "id = ?", paramId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Error",
			"message": "Failed to delete note",
			"data":    err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "Success",
		"message": "Note deleted",
		"data":    note,
	})

}
