package consumer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	bolt "go.etcd.io/bbolt"
)

type mockedDataSource struct {
	mock.Mock
}

func (m *mockedDataSource) View(fn func(*bolt.Tx) error) error {
	args := m.Called(fn)
	return args.Error(0)
}

func (m *mockedDataSource) Update(fn func(*bolt.Tx) error) error {
	args := m.Called(fn)
	return args.Error(0)
}

func TestProcessAnimalsBolt(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		assert := assert.New(t)
		testDB := new(mockedDataSource)
		defer testDB.AssertExpectations(t)

		testDB.On("Update", mock.AnythingOfType("func(*bbolt.Tx) error")).Return(nil)
		testDB.On("View", mock.AnythingOfType("func(*bbolt.Tx) error")).Return(nil)

		animals := []Animal{
			GetAnimal("Dog"),
			GetAnimal("Cat"),
		}

		err := ProcessAnimalsBolt(testDB, animals)
		assert.NoError(err)
	})
	t.Run("error on reading the data", func(t *testing.T) {
		assert := assert.New(t)
		testDB := new(mockedDataSource)

		testDB.On("Update", mock.AnythingOfType("func(*bbolt.Tx) error")).Return(nil)
		testDB.On("View", mock.AnythingOfType("func(*bbolt.Tx) error")).Return(errors.New("puf"))

		animals := []Animal{
			GetAnimal("Dog"),
			GetAnimal("Cat"),
		}

		err := ProcessAnimalsBolt(testDB, animals)
		assert.Error(err)
		testDB.AssertExpectations(t)
	})
	t.Run("error on writing the data", func(t *testing.T) {
		assert := assert.New(t)
		testDB := new(mockedDataSource)

		testDB.On("Update", mock.AnythingOfType("func(*bbolt.Tx) error")).Return(errors.New("puf"))

		animals := []Animal{
			GetAnimal("Dog"),
			GetAnimal("Cat"),
		}

		err := ProcessAnimalsBolt(testDB, animals)
		assert.Error(err)
		testDB.AssertExpectations(t)
	})

}
