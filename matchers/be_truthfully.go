package matchers

import (
	"fmt"
	"strconv"

	"github.com/onsi/gomega/types"
)

func BeTruthfully(expected bool) types.GomegaMatcher {
	return &beTruthfully{
		expected: expected,
	}
}

func BeTruthfullyTrue() types.GomegaMatcher {
	return &beTruthfully{
		expected: true,
	}
}

func BeTruthfullyFalse() types.GomegaMatcher {
	return &beTruthfully{
		expected: false,
	}
}

type beTruthfully struct {
	expected bool
}

func (matcher *beTruthfully) Match(actual interface{}) (success bool, err error) {
	actualBool, err := convertActualToBool(actual)

	if err != nil {
		return false, err
	}

	return matcher.expected == actualBool, nil
}

func (matcher *beTruthfully) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto be equivalent to\n\t%#v", actual, matcher.expected)
}

func (matcher *beTruthfully) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nnot to be equivalent to\n\t%#v", actual, matcher.expected)
}

func convertActualToBool(actual interface{}) (bool, error) {
	switch actual.(type) {
	case string:
		actualBool, err := strconv.ParseBool(actual.(string))
		if err != nil {
			return false, fmt.Errorf("ERROR: Failed to parse %s to bool: %s", actual, err.Error())
		}

		return actualBool, nil
	case int:
		actualValue := actual.(int)
		var actualBool bool = false
		switch actualValue {
		case 1:
			actualBool = true
		case 0:
			actualBool = false
		default:
			return false, fmt.Errorf("ERROR: invalid actual value: %v .  Value must be 1 or 0", actual)
		}

		return actualBool, nil
	case int8:
		actualValue := actual.(int8)
		var actualBool bool = false
		switch actualValue {
		case 1:
			actualBool = true
		case 0:
			actualBool = false
		default:
			return false, fmt.Errorf("ERROR: invalid actual value: %v .  Value must be 1 or 0", actual)
		}

		return actualBool, nil
	case int16:
		actualValue := actual.(int16)
		var actualBool bool = false
		switch actualValue {
		case 1:
			actualBool = true
		case 0:
			actualBool = false
		default:
			return false, fmt.Errorf("ERROR: invalid actual value: %v .  Value must be 1 or 0", actual)
		}

		return actualBool, nil
	case int32:
		actualValue := actual.(int32)
		var actualBool bool = false
		switch actualValue {
		case 1:
			actualBool = true
		case 0:
			actualBool = false
		default:
			return false, fmt.Errorf("ERROR: invalid actual value: %v .  Value must be 1 or 0", actual)
		}

		return actualBool, nil
	case int64:
		actualValue := actual.(int64)
		var actualBool bool = false
		switch actualValue {
		case 1:
			actualBool = true
		case 0:
			actualBool = false
		default:
			return false, fmt.Errorf("ERROR: invalid actual value: %v .  Value must be 1 or 0", actual)
		}

		return actualBool, nil
	case bool:
		actualBool := actual.(bool)
		return actualBool, nil
	case uint:
		actualValue := actual.(uint)
		var actualBool bool = false
		switch actualValue {
		case 1:
			actualBool = true
		case 0:
			actualBool = false
		default:
			return false, fmt.Errorf("ERROR: invalid actual value: %v .  Value must be 1 or 0", actual)
		}

		return actualBool, nil
	case uint8:
		actualValue := actual.(uint8)
		var actualBool bool = false
		switch actualValue {
		case 1:
			actualBool = true
		case 0:
			actualBool = false
		default:
			return false, fmt.Errorf("ERROR: invalid actual value: %v .  Value must be 1 or 0", actual)
		}

		return actualBool, nil
	case uint16:
		actualValue := actual.(uint16)
		var actualBool bool = false
		switch actualValue {
		case 1:
			actualBool = true
		case 0:
			actualBool = false
		default:
			return false, fmt.Errorf("ERROR: invalid actual value: %v .  Value must be 1 or 0", actual)
		}

		return actualBool, nil
	case uint32:
		actualValue := actual.(uint32)
		var actualBool bool = false
		switch actualValue {
		case 1:
			actualBool = true
		case 0:
			actualBool = false
		default:
			return false, fmt.Errorf("ERROR: invalid actual value: %v .  Value must be 1 or 0", actual)
		}

		return actualBool, nil
	case uint64:
		actualValue := actual.(uint64)
		var actualBool bool = false
		switch actualValue {
		case 1:
			actualBool = true
		case 0:
			actualBool = false
		default:
			return false, fmt.Errorf("ERROR: invalid actual value: %v .  Value must be 1 or 0", actual)
		}

		return actualBool, nil
	default:
		return false, fmt.Errorf("ERROR: Unsupported actual type: %t.  Expected to string, int, uint", actual)
	}
}
