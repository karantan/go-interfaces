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

func TestAddCity(t *testing.T) {
	assert := assert.New(t)
	db, _ := GetMySQL(dockerMySQL)
	got, err := AddCity(db, City{Name: "Ljubljana", Population: 200_000})
	assert.NoError(err)
	assert.Greater(got, int64(8))
}

func TestGetCityByID(t *testing.T) {
	assert := assert.New(t)
	db, _ := GetMySQL(dockerMySQL)
	got, err := GetCityByID(db, 8)
	want := City{Id: 8, Name: "Berlin", Population: 3_671_000}
	assert.NoError(err)
	assert.Equal(got, want)
}

func TestDeleteCityByIDs(t *testing.T) {
	assert := assert.New(t)
	db, _ := GetMySQL(dockerMySQL)

	type args struct {
		db  MySQL
		ids []int64
	}
	tests := []struct {
		name             string
		args             args
		wantAffectedRows int64
		wantErr          error
	}{
		{name: "0 rows", args: args{db: db, ids: []int64{}}, wantAffectedRows: 0, wantErr: nil},
		{name: "1 row", args: args{db: db, ids: []int64{1}}, wantAffectedRows: 1, wantErr: nil},
		{name: "2 rows", args: args{db: db, ids: []int64{2, 3}}, wantAffectedRows: 2, wantErr: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAffectedRows, err := DeleteCityByIDs(tt.args.db, tt.args.ids...)
			assert.Equal(gotAffectedRows, tt.wantAffectedRows)
			assert.Equal(err, tt.wantErr)
		})
	}
}

func Test_sliceToString(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		a     []int64
		delim string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0 elements", args: args{[]int64{}, ", "}, want: ""},
		{name: "1 element", args: args{[]int64{1}, ", "}, want: "1"},
		{name: "2 elements", args: args{[]int64{1, 2}, ", "}, want: "1, 2"},
		{name: "3 elements", args: args{[]int64{1, 2, 3}, ", "}, want: "1, 2, 3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(sliceToString(tt.args.a, tt.args.delim), tt.want)
		})
	}
}
