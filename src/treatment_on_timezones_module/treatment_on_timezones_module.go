package treatment_on_timezones_module

import (
        "strings"
	"treatment_on_time_module"
        "time"
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

//Function to pilot the treatment of timezones passed in argument to return a global array with all datas
func Master_function(bruts_timezones_array []string) []string {

	var master_array []string

        timezones_array := tz_extraction_function(bruts_timezones_array)

        for incrementer := range timezones_array {

                tz_loc, result_error := time.LoadLocation(timezones_array[incrementer])

		if result_error == nil {

			tz_now := time.Now().In(tz_loc)

			master_array = append(master_array, timezones_array[incrementer] + " : " + treatment_on_time_module.Time_shaping_function(tz_now))

		} else {

			master_array = append(master_array, "Une erreur est survenue: la timezone " + timezones_array[incrementer] + " n'est pas valide...")

		}

        }

        //return master_array
	return master_array
}
