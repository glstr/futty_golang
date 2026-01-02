package models

type TraceData struct {
	IP       string
	Port     int
	Protocol string
	Data     []byte
}
