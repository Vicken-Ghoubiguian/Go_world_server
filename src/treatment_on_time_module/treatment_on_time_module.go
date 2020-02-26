package treatment_on_time_module

import (
        "time"
	"strconv"
)

func int_shaping_function(int_to_treat int) string {

	if int_to_treat < 10 {

		return "0" + strconv.Itoa(int_to_treat)

	} else {

		return strconv.Itoa(int_to_treat)

	}
}

func Time_shaping_function(current_time_for_tz time.Time) string {

	return int_shaping_function(current_time_for_tz.Hour()) + ":" + int_shaping_function(current_time_for_tz.Minute()) + ":" + int_shaping_function(current_time_for_tz.Second()) + " " + int_shaping_function(current_time_for_tz.Day()) + "/" + int_shaping_function(int(current_time_for_tz.Month())) + "/" + int_shaping_function(current_time_for_tz.Year())

}
