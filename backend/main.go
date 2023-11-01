package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const (
	HOST = "localhost"
	PORT = 8000

	STATUS_AMAZON = "amazon-status"
	STATUS_GOOGLE = "google-status"
	STATUS_ALL    = "all-status"
)

// Email is used for the main contents of email
type Status struct {
	Url        string `json:"url"`
	StatusCode int    `json:"statusCode"`
	Duration   int64  `json:"duration"`
	Date       int64  `json:"date"`
}

func doRequest(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

// StatusHandler is parsing emails.zip folder and return slice of email threads to the response
func StatusHandler(c *fiber.Ctx) error {
	urls := map[string]string{
		STATUS_GOOGLE: "https://www.google.com",
		STATUS_AMAZON: "https://www.amazon.com",
	}
	start := time.Now()
	status := c.Params("status")

	statusArr := []string{}
	if status == STATUS_AMAZON {
		statusArr = append(statusArr, STATUS_AMAZON)
	} else if status == STATUS_GOOGLE {
		statusArr = append(statusArr, STATUS_GOOGLE)
	} else if status == STATUS_ALL {
		statusArr = append(statusArr, STATUS_AMAZON, STATUS_GOOGLE)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "status is invalid"})
	}

	resChan := make(chan Status, 2)
	errorChan := make(chan error, 2)
	var wg sync.WaitGroup

	wg.Add(len(statusArr))
	for _, st := range statusArr {
		go func(stat string) {
			defer wg.Done()
			statusCode, err := doRequest(urls[STATUS_AMAZON])
			if err != nil {
				errorChan <- err
			}
			resChan <- Status{
				Url:        urls[stat],
				StatusCode: statusCode,
				Duration:   time.Since(start).Milliseconds(),
				Date:       time.Now().UnixMilli(),
			}
		}(st)
	}
	wg.Wait()
	close(resChan)
	close(errorChan)

	for err := range errorChan {
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		}
	}

	res := make([]Status, 0, 2)
	for item := range resChan {
		res = append(res, item)
	}

	if len(res) == 1 {
		return c.JSON(res[0])
	}
	return c.JSON(res)

}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "POST",
		AllowHeaders: "*",
	}))

	app.Get("/v1/:status", StatusHandler)

	fmt.Printf("Server is running at: localhost:%d", PORT)
	app.Listen(fmt.Sprintf(":%d", PORT))
}
