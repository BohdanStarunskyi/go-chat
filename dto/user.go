package dto

type UserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewUserResponse(
	id int64,
	name string,
) UserResponse {
	return UserResponse{
		ID:   id,
		Name: name,
	}
}
