package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	fb "firebase.google.com/go"
	"fmt"
	"github.com/PrettyBoyHelios/agape-backend/models"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"time"
)

type FirebaseAdmin struct {
	firestore *firestore.Client
}

func NewFirebaseAdmin() *FirebaseAdmin {
	f := new(FirebaseAdmin)
	opt := option.WithCredentialsFile("agape.json")
	app, err := fb.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
		return nil
	}
	f.firestore, err = app.Firestore(context.Background())
	return f
}

func (f *FirebaseAdmin) GetUsers() []models.User {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	ref := f.firestore.Collection("users")
	docIterator := ref.Documents(ctx)
	docSnap, err := docIterator.GetAll()
	if err != nil {
		panic(err.Error())
	}
	var users []models.User
	for _, doc := range docSnap {
		var user models.User
		_ = doc.DataTo(&user)
		users = append(users, user)
	}
	return users
}

func (f *FirebaseAdmin) CreateCouple(couple models.Couple) string {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	ref := f.firestore.Collection("couples")
	id, _, _ := ref.Add(ctx, couple)
	fmt.Println(id.ID)
	return id.ID
}

func (f *FirebaseAdmin) Update(couple models.Couple, doc string) models.Couple {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	ref := f.firestore.Collection("couples")
	ref.Doc(doc).Set(ctx, couple)
	return couple
}

// GetUserCouples This retrieves all user couples as both a creator and couple into a single array, to restrict couple creation when still in a relationship.
func (f *FirebaseAdmin) GetUserCouples(userID string) []models.Couple {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	var couplesCreated []models.Couple
	docSnap := f.firestore.Collection("couples").Where("creator_id", "==", userID).Documents(ctx)
	for {
		var couple models.Couple
		doc, err := docSnap.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err.Error())
			return couplesCreated
		}
		fmt.Println(doc.Data())
		_ = doc.DataTo(&couple)
		couplesCreated = append(couplesCreated, couple)
	}

	couplesAsCouple := f.firestore.Collection("couples").Where("couple_id", "==", userID).Documents(ctx)

	for {
		var couple models.Couple
		doc, err := couplesAsCouple.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err.Error())
			return couplesCreated
		}
		fmt.Println(doc.Data())
		_ = doc.DataTo(&couple)
		couplesCreated = append(couplesCreated, couple)
	}
	var res []models.Couple

	for _, couple := range couplesCreated {
		if couple.CreatorImages == nil {
			couple.CreatorImages = []models.Images{}
		}
		if couple.CoupleImages == nil {
			couple.CoupleImages = []models.Images{}
		}
		res = append(res, couple)
	}
	return res
}

// GetCoupleByPairingCode This retrieves a created couple by pairing code.
func (f *FirebaseAdmin) GetCoupleByPairingCode(pairingCode string) (models.Couple, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	var couple models.Couple
	var docId string
	docSnap := f.firestore.Collection("couples").Where("pairing_code", "==", pairingCode).Documents(ctx)
	for {
		doc, err := docSnap.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err.Error())
			return couple, ""
		}
		fmt.Println(doc.Data())
		_ = doc.DataTo(&couple)
		docId = doc.Ref.ID

	}
	return couple, docId
}
