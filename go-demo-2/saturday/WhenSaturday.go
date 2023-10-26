package saturday

import (
	"time"
)

// whenisaturday returns when saturday is.
func WhenSaturday() string {
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		{
			return "Today"
		}
	case today + 1:
		{
			return "Tomorrow"
		}
	case today + 2:
		{
			return "In two days"
		}
	default:
		{
			return "saturday is too far away"
		}

	}
}
