package radosAPI

import (
	"errors"
	"fmt"
	"time"

	"github.com/QuentinPerez/go-encodeUrl"
)

func init() {
	encurl.AddEncodeFunc(ifTimeIsNotNilCeph)
}

func ifTimeIsNotNilCeph(obj interface{}) (string, bool, error) {
	if val, ok := obj.(*time.Time); ok {
		if val != nil {
			return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
				val.Year(), val.Month(), val.Day(),
				val.Hour(), val.Minute(), val.Second()), true, nil
		}
		return "", false, nil
	}
	return "", false, errors.New("this field should be a *time.Time")
}
