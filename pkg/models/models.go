package models

type Person struct {
	ID        uint   `json:”id”`
	FirstName string `json:”firstname”`
	LastName  string `json:”lastname”`
}



type Params struct {
	Type       string `json:"type"`
	DayDate    string `json:"dayDate"`
	WeekDate   string `json:"weekDate"`
	MonthDate  string `json:"monthDate"`
	Response   string `json:"response"`
	Controller string `json:"controller"`
	Action     string `json:"action"`
	Lang       string `json:"lang"`
	Timestamp  string `json:"_"`
}
