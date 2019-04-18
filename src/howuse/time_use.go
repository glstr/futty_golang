package howuse

import (
	"log"
	"time"
)

//ParseTime show useage of time.Parse
func ParseTime() {
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	t, err := time.ParseInLocation("2006-01-02T15:04:05-0700", "2019-04-17T09:23:47+0800", beijing)
	if err != nil {
		log.Printf("err_msg:%s", err.Error())
		return
	}
	log.Printf("t:%v", t.Unix())
}

func LoadLocation() {

}
