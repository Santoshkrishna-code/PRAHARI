package holiday

import "time"

// Holiday represents a declared statutory or facility holiday.
type Holiday struct {
	ID          string    `json:"id"`
	CalendarID  string    `json:"calendar_id"`
	Name        string    `json:"name"`
	HolidayDate time.Time `json:"holiday_date"`
	IsRecurring bool      `json:"is_recurring"`
}
