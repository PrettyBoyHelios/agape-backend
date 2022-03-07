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
	CreatorID     string       `firestore:"creator_id" json:"creator_id"`
	CoupleID      string       `firestore:"couple_id" json:"couple_id"`
	CreatedAt     int64        `firestore:"created_at" json:"created_at"`
	PairingCode   string       `firestore:"pairing_code" json:"pairing_code"`
	Status        CoupleStatus `firestore:"status" json:"status"`
	CreatorImages []Images     `firestore:"creator_images" json:"creator_images"`
	CoupleImages  []Images     `firestore:"couple_images" json:"couple_images"`
}

type CreateCoupleInput struct {
	CreatorID string `json:"creator_id"`
}
type JoinCoupleInput struct {
	ID       string `json:"id"`
	PairCode string `json:"pair_code"`
}

type Images struct {
	Name      string `firestore:"name" json:"name"`
	Note      string `firestore:"note" json:"note"`
	Timestamp string `firestore:"timestamp" json:"timestamp"`
}
