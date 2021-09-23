package consumer

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockedRedisSource struct {
	mock.Mock
}

func (m *mockedRedisSource) Get(ctx context.Context, key string) *redis.StringCmd {
	args := m.Called(ctx, key)
	stringCmd := redis.NewStringCmd(ctx, "select", 1)
	stringCmd.SetErr(args.Error(0))
	return stringCmd
}

func (m *mockedRedisSource) Set(ctx context.Context, key string, val interface{}, t time.Duration) *redis.StatusCmd {
	args := m.Called(ctx, key, val, t)
	statusCmd := redis.NewStatusCmd(ctx, "select", 1)
	statusCmd.SetErr(args.Error(0))
	return statusCmd
}

func TestProcessAnimalsRedis(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		assert := assert.New(t)
		testDB := new(mockedRedisSource)
		defer testDB.AssertExpectations(t)
		var ctx = context.Background()

		animals := []Animal{GetAnimal("Dog")}

		testDB.On("Set", ctx, string(animals[0].hash), animals[0].name, 0*time.Second).Return(nil)
		testDB.On("Get", ctx, string(animals[0].hash)).Return(nil)

		err := ProcessAnimalsRedis(testDB, animals)
		assert.NoError(err)
	})

	t.Run("error on set", func(t *testing.T) {
		assert := assert.New(t)
		testDB := new(mockedRedisSource)
		defer testDB.AssertExpectations(t)
		var ctx = context.Background()

		animals := []Animal{GetAnimal("Dog")}

		testDB.On("Set", ctx, string(animals[0].hash), animals[0].name, 0*time.Second).Return(errors.New("puf"))

		err := ProcessAnimalsRedis(testDB, animals)
		assert.Error(err)
	})

	t.Run("error on get", func(t *testing.T) {
		assert := assert.New(t)
		testDB := new(mockedRedisSource)
		defer testDB.AssertExpectations(t)
		var ctx = context.Background()

		animals := []Animal{GetAnimal("Dog")}

		testDB.On("Set", ctx, string(animals[0].hash), animals[0].name, 0*time.Second).Return(nil)
		testDB.On("Get", ctx, string(animals[0].hash)).Return(errors.New("puf"))

		err := ProcessAnimalsRedis(testDB, animals)
		assert.Error(err)
	})
}
