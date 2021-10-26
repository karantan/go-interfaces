package good

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

var dockerMySQL = "root@tcp(localhost:3306)/testdb"

func TestGetMySQL(t *testing.T) {
	assert := assert.New(t)
	db, err := GetMySQL(dockerMySQL)
	assert.NoError(err)
	assert.NoError(db.Ping())
}

func TestGetCities(t *testing.T) {
	assert := assert.New(t)
	db, _ := GetMySQL(dockerMySQL)
	got := GetCities(db)
	want := []City{
		{Id: 1, Name: "Bratislava", Population: 432000},
		{Id: 2, Name: "Budapest", Population: 1759000},
		{Id: 3, Name: "Prague", Population: 1280000},
		{Id: 4, Name: "Warsaw", Population: 1748000},
		{Id: 5, Name: "Los Angeles", Population: 3971000},
		{Id: 6, Name: "New York", Population: 8550000},
		{Id: 7, Name: "Edinburgh", Population: 464000},
		{Id: 8, Name: "Berlin", Population: 3671000},
	}
	assert.Equal(got, want)
}
