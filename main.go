package main

import (
  "fmt"
  "net/http"
  "os"
  "time"
)

func dayOfTheYear(w http.ResponseWriter, r *http.Request) {
  pacific, err := time.LoadLocation("America/Los_Angeles")
  if err != nil {
    panic(err)
  }

  now := time.Now().In(pacific)
  firstDayOfThisYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, pacific)
  firstDayOfNextYear := time.Date(now.Year() + 1, 1, 1, 0, 0, 0, 0, pacific)
  totalDaysOfThisYear := daysBetween(firstDayOfThisYear, firstDayOfNextYear)

  dayOfThisYear := now.YearDay()
  percentage := float64(dayOfThisYear) / float64(totalDaysOfThisYear)
  daysUntilNextYear := totalDaysOfThisYear - dayOfThisYear

  message := fmt.Sprintf("Today is Day: %d.\n", dayOfThisYear)
  message += fmt.Sprintf("We are %.1f%% through the year.\n", percentage * 100)
  message += fmt.Sprintf("There are %d days left in the year.", daysUntilNextYear)
  w.Write([]byte(message))
}

func daysBetween(earlier, later time.Time) int {
  return int(later.Sub(earlier).Hours() / 24)
}

func main() {
  http.HandleFunc("/", dayOfTheYear)
  port := os.Getenv("PORT")
  if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
    panic(err)
  }
}
