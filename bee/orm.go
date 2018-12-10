package main

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
    Userid   int `orm:"auto"`
    Username string `orm:"size(45)"`
    Userage int
    Usersex int
}

func init() {
    // set default database
    orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/golang?charset=utf8", 30)

    // register model
    orm.RegisterModel(new(User))

    // create table
    orm.RunSyncdb("default", false, true)
}

func main() {
    o := orm.NewOrm()

    user := User{Username: "slene"}

    // insert
    id, err := o.Insert(&user)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    // update
    user.Username = "astaxiea"
    user.Userage = 1;
    user.Usersex =12345
    num, err := o.Update(&user)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)

    // read one
    u := User{Userid: user.Userid}
    err = o.Read(&u)
    fmt.Printf("ERR: %v\n", err)
    fmt.Printf("id: %d name:%s\n", u.Userid,u.Username)
}