// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/aakritigkmit/my-go-crud/internal/handlers"
// 	"github.com/aakritigkmit/my-go-crud/internal/repository"
// 	"github.com/aakritigkmit/my-go-crud/internal/routes"
// 	"github.com/aakritigkmit/my-go-crud/internal/services"
// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// // mongodb+srv://khannaaakriti206:<db_password>@cluster0.27u43.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0

// func init() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("error whilereading .env")
// 	}
// }
// func mongoConnection() *mongo.Client {
// 	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
// 	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)
// 	// Create a new client and connect to the server
// 	client, err := mongo.Connect(context.TODO(), opts)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// defer func() {
// 	// 	if err = client.Disconnect(context.TODO()); err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// }()
// 	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Connected to MongoDB!")
// 	return client
// }

// func main() {
// 	mongoClient := mongoConnection()
// 	defer mongoClient.Disconnect(context.Background())

// 	collection := mongoClient.Database(os.Getenv("MONGO_DBNAME")).Collection(os.Getenv("MONGO_COLLECTION_NAME"))

// 	userRepo := repository.NewUserRepo(collection)
// 	authRepo := repository.NewAuthRepo(collection)
// 	userService := services.NewUserService(userRepo)
// 	userHandler := handlers.NewUserHandler(userService)

// 	authService := services.NewAuthService(authRepo, userRepo)
// 	authHandler := handlers.NewAuthHandler(authService)

// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)
// 	r.Route("/api", func(subRouter chi.Router) {
// 		subRouter.Get("/health", func(w http.ResponseWriter, r *http.Request) {
// 			w.Write([]byte(" server is healthy"))

// 		})

// 		subRouter.Mount("/users", routes.UserRoutes(userHandler))
// 		subRouter.Mount("/auth", routes.AuthRoutes(authHandler))

// 	})

// 	fmt.Println("Server started on port :4444")
// 	http.ListenAndServe(":4444", r)
// }



package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mongoClient, router := initializeApp()
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Fatal("Error disconnecting MongoDB:", err)
		}
		fmt.Println("Disconnected from MongoDB")
	}()

	port := ":4444"
	fmt.Println("Server started on port", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
