package good

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Have a global DB variable and use it for all requests.
var DB MySQL

type MySQL struct {
	*sql.DB
}

type City struct {
	Id         int
	Name       string
	Population int
}

// GetMySQL returns a Database connection to the MySQL service with
// `dataSource` as the data source.
// Example of `dataSource` data source:
// 		"username:password@tcp(127.0.0.1:3306)/test"
// If the sql.DB should not have a lifetime beyond the scope of the function you should
// close its connection (i.e. `defer db.Close`)
func GetMySQL(dataSource string) (MySQL, error) {
	// return pooled connection
	if DB.DB != nil {
		return DB, nil
	}
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return MySQL{}, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(time.Hour)
	DB = MySQL{db}
	return DB, nil
}

// GetCities returns a list of all cities in the `db`.
func GetCities(db MySQL) (cities []City) {
	res, err := db.Query("SELECT * FROM cities")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var city City
		err := res.Scan(&city.Id, &city.Name, &city.Population)
		cities = append(cities, city)
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}

// AddCity inserts a City `city` to the `db` and returns it's new ID.
func AddCity(db MySQL, city City) (int64, error) {

	sql := fmt.Sprintf(
		"INSERT INTO cities(name, population) VALUES ('%s', %d)",
		city.Name, city.Population,
	)
	res, err := db.Exec(sql)

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// GetCityByID returns a `City` identified by ID `id`.
func GetCityByID(db MySQL, id int64) (city City, err error) {
	res, err := db.Query("SELECT * FROM cities WHERE id = ?", id)
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	if res.Next() {
		err := res.Scan(&city.Id, &city.Name, &city.Population)

		if err != nil {
			log.Fatal(err)
		}
	}
	return
}

// DeleteCityByIDs deletes `City`s which IDs are in `ids`
func DeleteCityByIDs(db MySQL, ids ...int64) (affectedRows int64, err error) {
	if len(ids) == 0 {
		return 0, nil
	}
	sql := fmt.Sprintf("DELETE FROM cities WHERE id IN (%s)", sliceToString(ids, ", "))
	res, err := db.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}

	return res.RowsAffected()
}

func sliceToString(a []int64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
