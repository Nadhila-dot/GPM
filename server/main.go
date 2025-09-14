package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"nadhi.dev/binaries/server/scrapper"
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using default port 3000")
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    // testing data
    result, err := scrapper.FetchSearchContent("Fiber", 25)
    if err != nil {
        log.Fatalf("Error fetching search content: %v", err)
    }
    log.Printf("Fetched %d packages\n", len(result))

    // Write result to a txt file
    f, err := os.Create("result.txt")
    if err != nil {
        log.Fatalf("Error creating file: %v", err)
    }
    defer f.Close()

    fmt.Fprintf(f, "%+v\n", result)

    pkgs, err := scrapper.ParseSearchContent(result)
    if err != nil {
        // handle error
    }
    jsonStr, _ := scrapper.ResultsToJSON(pkgs)
    fmt.Println(jsonStr)



    app := fiber.New()

    RegisterRoutes(app)

    log.Printf("Server running on port %s\n", port)
    log.Fatal(app.Listen(":" + port))
}