package user

type CreateUserRequest struct {
	Name    string `json:"name" binding:"required"`
	Age     int    `json:"age" binding:"required"`
	PetName string `json:"pet_name,omitempty"`
}
