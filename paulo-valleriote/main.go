package calculator

import "errors"

var (
	ErrInvalidNumbersQuantity = errors.New("at least two numbers are required to make this operation")
	ErrInvalidZeroDivision    = errors.New("division operation cannot be realized with the divisor equal to zero")
)

func Sum(nums ...int) (int, error) {
	if err := validateAtLeastTwoNumbers(nums...); err != nil {
		return 0, err
	}

	res, numsToSum := splitFirstOperatorFromSlice(nums...)
	for _, num := range numsToSum {
		res += num
	}
	return res, nil
}

func Sub(nums ...int) (int, error) {
	if err := validateAtLeastTwoNumbers(nums...); err != nil {
		return 0, err
	}

	res, numsToSub := splitFirstOperatorFromSlice(nums...)
	for _, num := range numsToSub {
		res -= num
	}
	return res, nil
}

func Multi(nums ...int) (int, error) {
	if err := validateAtLeastTwoNumbers(nums...); err != nil {
		return 0, err
	}

	res, numsToMultiplicate := splitFirstOperatorFromSlice(nums...)
	for _, num := range numsToMultiplicate {
		res *= num
	}
	return res, nil
}

func Div(nums ...int) (int, error) {
	if err := validateAtLeastTwoNumbers(nums...); err != nil {
		return 0, err
	}

	res, numsToDivide := splitFirstOperatorFromSlice(nums...)
	for _, num := range numsToDivide {
		if num == 0 {
			return 0, ErrInvalidZeroDivision
		}

		res /= num
	}
	return res, nil
}

func splitFirstOperatorFromSlice(nums ...int) (int, []int) {
	return nums[0], nums[1:]
}

func validateAtLeastTwoNumbers(nums ...int) error {
	if len(nums) < 2 {
		return ErrInvalidNumbersQuantity
	}

	return nil
}
