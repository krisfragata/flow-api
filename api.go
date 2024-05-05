package main

import (
	// "encoding/json"
	// "net/http"
	"time"
)

//flow data response
type flowResponse struct{
	//status code, usually 200
	Status int
	Date_Posted time.Time
	Date_String string
	Cfs string
	Time_Posted string
	Forecast string
	Expires string
	IsRelease bool
}

type Error struct{
	Code int
	Message string
}