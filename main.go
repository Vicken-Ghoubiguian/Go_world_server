package main

import (
	"fmt"
	"strings"
	"time"
	"net/http"
)

//Function to shaping the timezone to replace '&' by '/'
func tz_shaping_function(brut_timezone string) string {

	brut_tz_as_array := strings.Split(brut_timezone, "&")

	shaped_timezone := ""

	for incrementer := range brut_tz_as_array {

		shaped_timezone += brut_tz_as_array[incrementer]

		if incrementer != len(brut_tz_as_array) - 1 {

			shaped_timezone += "/"

		}
	}

	return shaped_timezone

}

//Function to treat all timezones in the array passed in parameter to return an array with all olsen timezones
func tz_extraction_function(bruts_timezones_array []string) []string {

	treated_timezones_array := []string{}

	for incrementer := range bruts_timezones_array {

		treated_timezone := tz_shaping_function(bruts_timezones_array[incrementer])

		treated_timezones_array = append(treated_timezones_array, treated_timezone)

	}

	return treated_timezones_array

}

//Function to pilot the treatment of timezones passed in argument to return global string with all datas
func master_function(bruts_timezones_array []string) string {

	master_string := "\n"

	timezones_array := tz_extraction_function(bruts_timezones_array)

	for incrementer := range timezones_array {

		tz_loc, _ := time.LoadLocation(timezones_array[incrementer])

		tz_now := time.Now().In(tz_loc)

		master_string += timezones_array[incrementer] + " : " + tz_now.String() + "\n"

	}

	return master_string
}

func main() {

	http.HandleFunc("/", handlerFunction)

	http.ListenAndServe(":8080", nil)

}

func handlerFunction(w http.ResponseWriter, r *http.Request) {

	bruts_timezones_array := strings.Split(r.URL.Path[1:], "/")

	master_string := master_function(bruts_timezones_array)

	fmt.Fprintf(w, master_string)

}
