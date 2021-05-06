package supertest

import (
	"fmt"

	"github.com/stretchr/testify/assert"
)

func (ctx *supertest) Expect(options Options) {
	if options.Key != "" {
		getHeader := ctx.httpRequest.Header.Get(fmt.Sprintf("%v", options.Key))
		assert.Equal(ctx.test, options.Value, getHeader)
	} else {
		assert.Equal(ctx.test, options.Value,  ctx.httpResponse.Code)
	}
}
