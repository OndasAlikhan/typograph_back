package dto

type Room struct {
	Texts     map[uint][]Letter `json:"texts"`
	Users     []UserResponse    `json:"users"`
	UsersDone map[uint]bool     `json:"users_done"`
	Status    string            `json:"status"` // waiting starting running finished
}
