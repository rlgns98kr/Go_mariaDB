package main

import (
	"github.com/joho/godotenv"
	"github.com/rlgns98kr/dbpractice/mariadb"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	m, err := mariadb.NewDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer m.CloseDatabase()

	//example query
	/*rows, err := m.Worker.Query("select * from book")
	if err != nil {
		log.Fatal(err.Error())
	}


	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(strconv.Itoa(id) + ":" + name)
	}*/
}
