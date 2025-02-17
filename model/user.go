package model

type User struct {
	ID         int    `gorm:"column:id;primary_key"`
	FacebookID int    `gorm:"column:facebook_id;unique"`
	Name       string `gorm:"column:name"`
	Age        int    `gorm:"column:age"`
	Friends    string `gorm:"column:friends"`
}

type UserRequest struct {
	FacebookID int    `json:"facebook_id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Friends    string `json:"friends"`
}

type UserResponse struct {
	ID         int    `json:"id"`
	FacebookID int    `json:"facebook_id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Friends    string `json:"friends"`
}

type MatchUserResponse struct {
	IsMatch bool `json:"is_match"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:         u.ID,
		FacebookID: u.FacebookID,
		Name:       u.Name,
		Age:        u.Age,
		Friends:    u.Friends,
	}
}

func UsersToResponses(users []User) []UserResponse {
	var responses []UserResponse
	for _, user := range users {
		responses = append(responses, user.ToResponse())
	}
	return responses
}
