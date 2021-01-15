package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)
 
type Document struct {
    Id   int
    Name string
    File string
}
 
func main() {
	// Create DB
	
	// Database Connection
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database created successfully")
	}

	// Create Database
	_,err = db.Exec("CREATE DATABASE sig_api_db")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database..")
	}

	// Choose the Database
	_,err = db.Exec("USE sig_api_db")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}

	// Create Table
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS documents (id int NOT NULL AUTO_INCREMENT, name varchar(50), file varchar(225), PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table created successfully..")
    }
    
    // Create
    /*
    _, err = db.Exec("INSERT INTO documents (name, file) VALUES (?,?)", "Test File", "test.pdf")
    if err != nil {
        panic(err)
    }
    */

    // Udpate
    /*id := 7;
    file := "test.pdf";
    _, err = db.Exec("UPDATE documents SET file = ? WHERE id = ?", file, id)
    if err != nil {
        panic(err)
    }*/

    // Delete
    /*id := 7;
    _, err = db.Exec("DELETE FROM documents WHERE id = ?", id)
    fmt.Println("Error creating DB:", err)
    if err != nil {
        panic(err)
        fmt.Println("There was an error completing this request")
    }*/

	// Close Database
    // defer db.Close()
    
}
 
func ErrorCheck(err error) {
    if err != nil {
        panic(err.Error())
    }
}
 
func PingDB(db *sql.DB) {
    err := db.Ping()
    ErrorCheck(err)
}
