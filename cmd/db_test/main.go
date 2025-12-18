package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/faisal-990/age/db/sqlc/generated" // Import the package you just generated
	_ "github.com/lib/pq"                         // Don't forget the driver!
)

func main() {
	// 1. Connect to Postgres
	conn, err := sql.Open("postgres", "postgres://task_user:task_user@123@localhost:5432/usermgmt?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 2. Initialize SQLC (The "New" function you asked about)
	// We pass 'conn' because it satisfies the DBTX interface
	queries := generated.New(conn)

	// 3. Test it out! (Create a user)
	ctx := context.Background()
	newUser, err := queries.CreateUser(ctx, generated.CreateUserParams{
		Name: "Gemini User",
		Dob:  time.Now(),
	})

	if err != nil {
		log.Fatal("cannot create user:", err)
	}

	fmt.Println("User created with ID:", newUser.ID)
}
