package supertest

import "time"

func (ctx *supertest) Timeout(timeType string, value time.Duration) {
	switch timeType {
	case "second":
		time.Sleep(time.Second * value)
		return
	case "minute":
		time.Sleep(time.Minute * value)
		return
	case "hours":
		time.Sleep(time.Hour * value)
		return
	}
}
