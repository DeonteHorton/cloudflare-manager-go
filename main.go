package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// url := os.Getenv("CLOUDFLARE_URL")
	email := os.Getenv("CLOUDFLARE_EMAIL")
	key := os.Getenv("CLOUDFLARE_KEY")
	// keyid := os.Getenv("CLOUDFLARE_KEYID")
	// secret := os.Getenv("CLOUDFLARE_SECRET")
	//	bucket := os.Getenv("CLOUDFLARE_BUCKET")

	api, err := cloudflare.New(key, email)
	if err != nil {
		log.Fatal(err)
	}

	// Most API calls require a Context
	ctx := context.Background()

	// Fetch user details on the account
	u, err := api.UserDetails(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print user details
	fmt.Println(u)
}
