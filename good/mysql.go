package good

import (
	"database/sql"

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
	DB = MySQL{db}
	return DB, nil
}

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
