//参考 go in action 5.4.4 多态
package main

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.Lshortfile)

//接口定义。 接口是用来定义行为的类型。
//这些被定义的行为不由接口直接实现，而是通过方法由用户定义的类型实现
//
//
type notifier interface {
	notify()
}

//user
type user struct {
	name  string
	email string
}

//user method 实现了notifier接口
func (u user) notify() {
	logger.Printf("User name : %s, email : %s",
		u.name, u.email)
}

//admin
type admin struct {
	name  string
	email string
}

//admin method 实现了notifier接口
func (a *admin) notify() {
	logger.Printf("Admin name : %s, email : %s",
		a.name, a.email)
}

//notifaction是一个普通函数，入参为interface。
//接受一个实现了notifier接口的值
func notification(n notifier) {
	n.notify()
}

func main() {
	usr := user{
		name:  "tom",
		email: "tom@qq.com",
	}
	notification(&usr)

	adm := admin{"jobs", "jobs@google.com"}
	notification(&adm)
}
