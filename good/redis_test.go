package good

import (
	"fmt"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	assert := assert.New(t)
	mr, err := miniredis.Run()
	assert.NoError(err)

	client, err := NewClient(fmt.Sprintf("redis://%s/0", mr.Addr()))
	assert.NoError(err)
	assert.IsType(MessageBroker{}, client)
}

func TestSetGetKey(t *testing.T) {
	assert := assert.New(t)
	mr, err := miniredis.Run()
	assert.NoError(err)
	client, err := NewClient(fmt.Sprintf("redis://%s/0", mr.Addr()))
	assert.NoError(err)

	err = SetKey(client, "foo", "bar")
	assert.NoError(err)

	got, err := GetKey(client, "foo")
	want := "bar"
	assert.NoError(err)
	assert.Equal(want, got)
}
