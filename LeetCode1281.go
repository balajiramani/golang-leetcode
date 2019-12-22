package main

func subtractProductAndSum(n int) int {
	var product int = 1
	var sum int
	var digit int
	for n > 0 {
		digit = n % 10
		n = n / 10
		product = product * digit
		sum = sum + digit
	}
	return product - sum
}
