package printdata

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type School struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Phone int `json:"phone"`
}

func RecByte(q []byte) {
	var school School
	json.Unmarshal(q, &school)
	fmt.Println(school.ID)


}
func Rec(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)

	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k)

	//fmt.Println(x)
	//field, ok := x.(*School)
	//
	//fmt.Println(field)
	//fmt.Println(ok)

}