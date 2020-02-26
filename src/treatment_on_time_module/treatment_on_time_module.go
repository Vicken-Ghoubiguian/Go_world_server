package treatment_on_time_module

import (
        //"strings"
        "time"
	"strconv"
)

func Time_shaping_function(current_time_for_tz time.Time) string {

	return strconv.Itoa(current_time_for_tz.Hour()) + ":" + strconv.Itoa(current_time_for_tz.Minute()) + ":" + strconv.Itoa(current_time_for_tz.Second()) + " " + strconv.Itoa(current_time_for_tz.Day()) + "/" + strconv.Itoa(int(current_time_for_tz.Month())) + "/" + strconv.Itoa(current_time_for_tz.Year())

}
