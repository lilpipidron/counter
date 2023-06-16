package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/status", Handler)
	e.GET("/time", Counter)
	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func Handler(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "test")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func Counter(ctx echo.Context) error {
	currentTime := time.Now()
	targetDate := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	timeDifference := targetDate.Sub(currentTime)

	days := int(timeDifference.Hours() / 24)
	hours := int(timeDifference.Hours()) % 24
	minutes := int(timeDifference.Minutes()) % 60
	seconds := int(timeDifference.Seconds()) % 60

	str := fmt.Sprintf("Дней: %d, Часов: %d, Минут: %d, Секунд: %d\n", days, hours, minutes, seconds)
	js := fmt.Sprintf(`
		<script>
			setTimeout(function() {
				window.location.href = '/time';
			}, 1000);
		</script>
		<div>%s</div>
	`, str)

	err := ctx.HTML(http.StatusOK, js)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
