package forPersonalTraining

import "errors"

func Divide(x, y int) (res int, er error) {
	if y == 0 {
		return 0, errors.New("cannot divide on zero")
	}

	return x / y, nil
}
