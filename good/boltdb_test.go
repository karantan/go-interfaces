package good

import (
	"go-interfaces/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDatabaseFunctional(t *testing.T) {
	assert := assert.New(t)
	want := utils.RootDir() + "/tmp.db"
	defer os.Remove(utils.RootDir() + "/tmp.db")

	_, err := GetDatabase(utils.RootDir()+"/tmp.db", false)
	assert.FileExists(want)
	assert.NoError(err)

	got, err := GetDatabase("", false)
	assert.Equal(got, &Database{})
	assert.Error(err)
}

func TestPutGetFunctional(t *testing.T) {
	assert := assert.New(t)
	db, _ := GetDatabase(utils.RootDir()+"/tmp.db", false)
	defer os.Remove(utils.RootDir() + "/tmp.db")

	empty, _ := Get(db, "myBucket", "key")
	assert.Equal(empty, "")

	err := Put(db, "myBucket", "key", "value")
	assert.NoError(err)
	got, _ := Get(db, "myBucket", "key")
	assert.Equal(got, "value")

	err = Put(db, "", "key", "value")
	assert.Error(err)

	err = Put(db, "myBucket", "", "value")
	assert.Error(err)
}
