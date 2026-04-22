package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func TrimAllStrings(a any) {
	v := reflect.ValueOf(a)
	if !v.IsValid() {
		return
	}
	trim(v, nil)
}

func trim(v reflect.Value, visited map[uintptr]struct{}) {
	switch v.Kind() {
	case reflect.Pointer:
		if v.IsNil() {
			return
		}
		addr := v.Pointer()
		if visited == nil {
			visited = map[uintptr]struct{}{}
		} else if _, ok := visited[addr]; ok {
			return
		}
		visited[addr] = struct{}{}
		trim(v.Elem(), visited)
	case reflect.Interface:
		if !v.IsNil() {
			trim(v.Elem(), visited)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			trim(v.Field(i), visited)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			trim(v.Index(i), visited)
		}
	case reflect.String:
		if !v.CanSet() {
			return
		}
		s := v.String()
		if t := strings.TrimSpace(s); t != s {
			v.SetString(t)
		}
	}
}

func main() {
	type Person struct {
		Name string
		Age  int
		Next *Person
	}

	a := &Person{
		Name: " name ",
		Age:  20,
		Next: &Person{
			Name: " name2 ",
			Age:  21,
			Next: &Person{
				Name: " name3 ",
				Age:  22,
			},
		},
	}

	TrimAllStrings(&a)

	m, _ := json.Marshal(a)

	fmt.Println(string(m))

	a.Next = a

	TrimAllStrings(&a)

	fmt.Println(a.Next.Next.Name == "name")
}
