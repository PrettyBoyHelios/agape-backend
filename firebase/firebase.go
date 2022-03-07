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
	return couplesCreated
}
