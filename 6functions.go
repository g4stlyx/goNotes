package main

func functions1(num1 int, num2 int) int {
	return (num1 + num2)
}

func functions2(num1 int, num2 int) (int, int, int, float32) {
	addition := num1 + num2
	subtraction := num1 - num2
	multiplication := num1 * num2
	division := float32(num1 / num2)

	return addition, subtraction, multiplication, division
}

func functions3(nums ...int) int { // variadic functions, when we dont know how many params'll be given
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}
