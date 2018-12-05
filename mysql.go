package main
import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
func main() {
	db, err := sql.Open("mysql", "jm_test:csfSyQieQELqte9H@(39.108.113.217)/test_jm_admin")
	fmt.Println(err,db);
}
