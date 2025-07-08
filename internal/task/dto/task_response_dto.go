package dto

type TaskResponseDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	UserId      string `json:"user_id"`
	Status      string `json:"status" validate:"oneof=pending in_progress completed cancelled"`
}
