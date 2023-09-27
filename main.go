package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"pie-fire-dire/meatcounter"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // Define the endpoint /beef/summary
    app.Get("/beef/summary", func(c *fiber.Ctx) error {
        // Fetch meat and filler text from the given API
        resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
        if err != nil {
            return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to fetch meat data",
            })
        }
        defer resp.Body.Close()
        // Read the response body
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to read response body",
            })
        }

        // Extract meat names and count them
        meatCounts := meatcounter.CountMeatNames(string(body))

        // Return the meat counts as JSON
        return c.JSON(fiber.Map{
            "beef": meatCounts,
        })
    })

    // Start the Fiber app
    port := 3000
    app.Listen(fmt.Sprintf(":%d", port))
}
