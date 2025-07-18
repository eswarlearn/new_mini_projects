// package controllers

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/eswar/mongo-golang/models"
// 	"github.com/julienschmidt/httprouter"
// 	"gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// )

// type UserController struct {
// 	session *mgo.Session
// }

// func NewUserController(s *mgo.Session) *UserController {
// 	return &UserController{s}
// }

// // no # function it is a sruct method
// func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	id := p.ByName("id")
// 	if !bson.IsObjectIdHex(id) {
// 		w.WriteHeader(http.StatusNotFound)
// 	}

// 	oid := bson.ObjectIdHex(id)

// 	u := models.User{}

// 	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(&u); err != nil {
// 		w.WriteHeader(404)
// 		return
// 	} //C - collections ; uc.session - from type

// 	uj, err := json.Marshal(u)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	w.Header().Set("Contenet-Type", "application/json")
// 	w.WriteHeader(http.StatusOK) //200
// 	fmt.Fprint(w, "%s\n", uj)
// }

// func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	u := models.User{}

// 	json.NewDecoder(r.Body).Decode(&u)

// 	u.Id = bson.NewObjectId() // generate new id

// 	uc.session.DB("mongo-golang").C("users").Insert(u)
// 	uj, err := json.Marshal(u)

// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated) //201
// 	fmt.Fprintf(w, "%s\n", uj)
// }

// func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	id := p.ByName("id")
// 	if !bson.IsObjectIdHex(id) {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	oid := bson.ObjectIdHex(id)

// 	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
// 		w.WriteHeader(404)
// 		// return
// 	}

// 	w.WriteHeader(http.StatusOK) //200
// 	fmt.Fprint(w, "Deleted user", oid, "\n")

// }

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/eswar/mongo-golang/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(client *mongo.Client) *UserController {
	return &UserController{client}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := p.ByName("id")
	collection := uc.client.Database("mongo-golang").Collection("users")

	var user models.User
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(uj)
}

// GetAllUsers lists all users from MongoDB
func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := uc.client.Database("mongo-golang").Collection("users")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err = cursor.All(ctx, &users); err != nil {
		http.Error(w, "Failed to parse users", http.StatusInternalServerError)
		return
	}

	resp, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

// UpdateUser updates a user by ID
func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := p.ByName("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedData models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	collection := uc.client.Database("mongo-golang").Collection("users")

	// Build the update document
	update := bson.M{
		"$set": bson.M{
			"name":   updatedData.Name,
			"age":    updatedData.Age,
			"gender": updatedData.Gender,
		},
	}

	result, err := collection.UpdateByID(ctx, objID, update)
	if err != nil || result.MatchedCount == 0 {
		http.Error(w, "User not found or update failed", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "✅ Updated user with ID: %s\n", id)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	u.Id = primitive.NewObjectID()

	collection := uc.client.Database("mongo-golang").Collection("users")
	_, err := collection.InsertOne(ctx, u)
	if err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(uj)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := p.ByName("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	collection := uc.client.Database("mongo-golang").Collection("users")
	res, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	if res.DeletedCount == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "✅ Deleted user with ID: %s\n", id)
}
