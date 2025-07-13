// package main

// import (
// 	"net/http"

// 	"github.com/eswar/mongo-golang/controllers"
// 	"github.com/julienschmidt/httprouter"
// 	"gopkg.in/mgo.v2"
// )

// func main() {

// 	r := httprouter.New()
// 	uc := controllers.NewUserController(getSession()) // for session
// 	r.GET("/user/:id", uc.GetUser)
// 	r.POST("/user", uc.CreateUser)
// 	r.DELETE("/user/:id", uc.DeleteUser)
// 	http.ListenAndServe("localhost:9000", r) // Start the server on port 9000
// }

// func getSession() *mgo.Session {
// 	s, err := mgo.Dial("127.0.0.1:27017") // safest option
// 	if err != nil {
// 		panic("MongoDB connection failed: " + err.Error())
// 	}
// 	return s
// }

package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/eswar/mongo-golang/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client := getClient()
	uc := controllers.NewUserController(client)

	r := httprouter.New()
	r.GET("/user/:id", uc.GetUser)
	// TODO: Add CreateUser, DeleteUser using same style
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)

	log.Println("🚀 Server starting at http://localhost:9000")
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}

func getClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
	return client
}
