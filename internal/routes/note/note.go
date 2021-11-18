package noteRoutes

import (
	"github.com/gofiber/fiber/v2"
	noteHandler "gofiber-gorm-mysql/internal/handlers/note"
)

func SetupNoteRouter(router fiber.Router) {
	note := router.Group("/note")
	note.Post("/", noteHandler.CreateNotes)
	note.Get("/", noteHandler.GetNotes)
	note.Get("/:noteId", noteHandler.GetNote)
	note.Put("/:noteId", noteHandler.UpdateNote)
	note.Delete("/:noteId", noteHandler.DeleteNote)
}
