package super

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

func(ctx *supertest) Send(payload interface{})  {

	var data interface{}

	encoded, _ := json.Marshal(payload)
	err := json.Unmarshal(encoded, &data)

	if err != nil {
		logrus.Error(err.Error())
		return
	}

	ctx.body.Data = data
}
