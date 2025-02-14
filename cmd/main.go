package main
//User-api
import (
	"log"
	"net/http"
	"os"
	"Task5/internal/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
    r := mux.NewRouter()
    godotenv.Load(".env")
    
    // Define routes
    r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
    r.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
    r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
    r.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")
    r.HandleFunc("/user/{id:[0-9a-fA-F-]{36}}", handlers.GetUserByIDHandler).Methods("GET")
    r.HandleFunc("/user/{id:[0-9a-fA-F-]{36}}", handlers.UpdateHandler).Methods("PUT")
    r.HandleFunc("/user/{id:[0-9a-fA-F-]{36}}", handlers.DeleteHandler).Methods("Delete")
    // Import Port from environment variables
    portString := os.Getenv("PORT")
    if portString == "" {
        log.Fatal("Port Not found in the environment")
    }
    log.Println("Server started on : " + portString)
    http.ListenAndServe(":"+portString, r)
}
