package encurl

import "errors"

func ifStringIsNotEmpty(obj interface{}) (string, bool, error) {
	if val, ok := obj.(string); ok {
		if val != "" {
			return val, true, nil
		}
		return "", false, nil
	}
	return "", false, errors.New("this field should be a string")
}

func ifBoolIsFalse(obj interface{}) (string, bool, error) {
	if val, ok := obj.(bool); ok {
		if !val {
			return "False", true, nil
		}
		return "", false, nil
	}
	return "", false, errors.New("this field should be a boolean")
}

func ifBoolIsTrue(obj interface{}) (string, bool, error) {
	if val, ok := obj.(bool); ok {
		if val {
			return "True", true, nil
		}
		return "", false, nil
	}
	return "", false, errors.New("this field should be a boolean")
}
