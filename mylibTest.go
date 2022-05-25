/*
调用自定义的包和类
*/
package main

import "./mylib"

func main() {
		var u mylib.User
		u.ChangeEmail("1012083552@qq.com")
		u.Notify()
}
