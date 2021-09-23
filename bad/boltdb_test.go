package bad

import (
	"go-interfaces/utils"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDatabase(t *testing.T) {
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

func TestDatabase_PutGet(t *testing.T) {
	assert := assert.New(t)
	db, _ := GetDatabase(utils.RootDir()+"/tmp.db", false)
	defer os.Remove(utils.RootDir() + "/tmp.db")

	empty, _ := db.Get("myBucket", "key")
	assert.Equal(empty, "")

	db.Put("myBucket", "key", "value")
	got, _ := db.Get("myBucket", "key")
	assert.Equal(got, "value")
}
