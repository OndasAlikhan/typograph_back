package dto

type Room struct {
	Texts     map[uint][][]Letter `json:"texts"`
	Users     []UserResponse      `json:"users"`
	UsersDone map[uint]bool       `json:"users_done"`
	Status    string              `json:"status"` // waiting starting running finished
}

type Letter struct {
	Char  string `json:"char"`
	Color string `json:"color"`
}

type UpdateTextMsg struct {
	Type    string     `json:"type"`
	LobbyID uint       `json:"lobby_id"`
	UserID  uint       `json:"user_id"`
	Text    [][]Letter `json:"text"`
}

type FinishMsg struct {
	Type    string `json:"type"`
	LobbyID uint   `json:"lobby_id"`
	UserID  uint   `json:"user_id"`
}
