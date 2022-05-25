/*
	自定义包和类
	类/方法的名称首字母必须大写
*/
package mylib 

import "fmt"



type User struct {
        name string
        email string
}

func (this User) Notify() {
        fmt.Printf("Email is %s\n", this.email)
}

func (this *User) ChangeEmail(email string) {
        this.email = email
}

