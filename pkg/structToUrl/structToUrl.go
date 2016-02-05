package structToUrl

import (
	"net/url"
	"reflect"
	"strings"

	"github.com/Sirupsen/logrus"
)

type TranslateFunc func(interface{}) (string, bool, error)

var (
	Funcs = make(map[string]TranslateFunc)
)

func Translate(obj interface{}) url.Values {
	if reflect.TypeOf(obj).Kind() != reflect.Struct &&
		reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return nil
	}
	values := url.Values{}
	e := reflect.TypeOf(obj).Elem()

	for i := 0; i < e.NumField(); i++ {
		field := e.Field(i)
		structFieldValue := reflect.ValueOf(obj).Elem().FieldByName(field.Name)
		if structFieldValue.IsValid() {
			tab := strings.Split(field.Tag.Get("url"), ",")
			if len(tab) > 1 {
				if validator, ok := Funcs[tab[1]]; ok {
					val, ok, err := validator(structFieldValue.Interface())
					if err != nil {
						logrus.Fatal(err)
					}
					if ok {
						values.Add(tab[0], val)
					}
				}
			}
		}
	}
	return values
}
