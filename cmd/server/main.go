package main

import (
	"context"
	"fmt"

	"github.com/Slava1989/Go_REST_API/internal/comment"
	"github.com/Slava1989/Go_REST_API/internal/db"
	// "go/doc/comment"
)

// MARK: responsible for instantiation and startup of app
func Run() error {
	fmt.Println("startin up our app")

	db, err := db.NewDataBase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)
	fmt.Println(cmtService.GetComment(context.Background(),
		"a32cb3c6-0a24-4484-8b6a-2a0e2beae999",
	))

	return nil
}

func main() {
	fmt.Println("Go Rest API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
