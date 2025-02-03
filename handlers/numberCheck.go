package handlers

import (
	"math"
	"strconv"
)

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsPerfect(n int) bool {
	if n <= 1 {
		return false
	}
	sum := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func IsEven(n int) bool {
	isEven := n%2 == 0
	return isEven
}

func IsArmStrong(n int) bool {
	if n < 0 {
		return false
	}
	numberStr := strconv.Itoa(n)
	numberOfDigits := len(numberStr)
	sum := 0
	for _, digit := range numberStr {
		digitInt, _ := strconv.Atoi(string(digit))
		sum += int(math.Pow(float64(digitInt), float64(numberOfDigits)))
	}
	return n == sum
}

func SumOfDigits(n int) int {
	numberString := strconv.Itoa(n)
	sum := 0
	for _, digit := range numberString {
		digitInt, _ := strconv.Atoi(string(digit))
		sum += digitInt
	}
	return sum
}
