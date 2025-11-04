package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mikehquan19/connect/routes"

	//"github.com/mikehquan19/connect/seed"
	"github.com/mikehquan19/connect/setup"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env vars")
	}
	//						From ENV 		 DB Name
	setup.ConnectDB(os.Getenv("MONGO_URI"), "Workshop")
	//seed.SeedUsers()
	r := gin.Default()

	routes.RegisterUserRoutes(r)

	r.Run(":8080")
}

/*
On first run, this will seed the database with initial users.


 --Hint
 	Uncomment the seed.SeedUsers() line above to seed the database on first run.
	and Uncomment the import for seed

	MAKE SURE TO comment the seed.SeedUsers() line after seeding is complete, so that it doesnt keep running on server start.
*/
