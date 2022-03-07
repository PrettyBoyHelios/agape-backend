package models

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Uid   string `json:"uid"`
}

type CoupleStatus string

const (
	INCOMPLETE = "INCOMPLETE"
	READY      = "READY"
	DELETED    = "DELETED"
)

type Couple struct {
	CreatorID     string       `json:"creator_id"`
	CoupleID      string       `json:"couple_id"`
	CreatedAt     int64        `json:"created_at"`
	PairingCode   int64        `json:"pairing_code"`
	Status        CoupleStatus `json:"status"`
	CreatorImages []Images     `json:"creator_images"`
	CoupleImages  []Images     `json:"couple_images"`
}

type Images struct {
	Name      string `json:"name"`
	Note      string `json:"note"`
	Timestamp string `json:"timestamp"`
}
