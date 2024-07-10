package dto

type LobbyWSCreateRequest struct {
	UserID uint `json:"user_id"`
}

type RoomWSMessage struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}
