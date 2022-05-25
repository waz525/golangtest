package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

/* 打印类的所有属性 */
func GetClassAttribute(body interface{}) {
		var prop []string
		refType := reflect.TypeOf(body)
	if refType.Kind() != reflect.Struct {
		fmt.Println("Not a structure type.")
	}
	for i:=0;i<refType.NumField();i++ {
		field := refType.Field(i)
		if field.Anonymous {
			prop = append(prop, field.Name)
			for j := 0;j<field.Type.NumField();j++ {
				prop = append(prop, field.Type.Field(j).Name)
			}
			continue
		}
		prop = append(prop, field.Name)
	}
	fmt.Printf("%v\n", prop)
}


func sayhelloName(w http.ResponseWriter, r *http.Request) {
   fmt.Println("打印Header参数列表：")
   if len(r.Header) > 0 {
      for k,v := range r.Header {
         fmt.Printf("%s = %s\n", k, v[0])
      }
   }
   fmt.Println("打印Form参数列表：")
   r.ParseForm()
   if len(r.Form) > 0 {
      for k,v := range r.Form {
         fmt.Printf("%s = %s\n", k, v[0])
      }
   }

   fmt.Println("打印r.URL信息：")
   fmt.Println("r.URL.Path: ",r.URL.Path)
   fmt.Println("r.URL.RawPath: ",r.URL.RawPath)
	GetClassAttribute(*r.URL)
   fmt.Println("===================================================\n\n")

	fmt.Fprintln(w, "hello world!")
}

func main() {
		pwd, _ := os.Getwd()
		fmt.Println("pwd: ",pwd)
	http.HandleFunc("/", sayhelloName)

	err := http.ListenAndServe(":12802", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)

	}
}
