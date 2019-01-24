package builder

import (
	"reflect"
)

type MenuItem struct {
	Title string
	Link  string
}

type Menu struct {
	Items []MenuItem
}

func newMenu(menu interface{}) *Menu {
	v := reflect.ValueOf(menu)
	var items []MenuItem
	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		res := make([]MenuItem, 0)
		for i := 0; i < v.Len(); i++ {
			vv, ok := v.Index(i).Interface().(map[string]interface{})
			if ok {
				var title, link string
				if k, ok := vv["title"]; ok {
					title = k.(string)
				}

				if k, ok := vv["link"]; ok {
					link = k.(string)
				}

				res = append(res, MenuItem{title, link})
			}
		}
		items = res
	}
	return &Menu{items}
}
