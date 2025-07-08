package dto

type TaskRequestDTO struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	DueDate     string `json:"due_date,omitempty"`
	UserId      string `json:"user_id,omitempty"`
	Status      string `json:"status,omitempty" validate:"oneof=pending in_progress completed cancelled"`
}
