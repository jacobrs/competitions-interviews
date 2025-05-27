func differenceOfSums(n int, m int) int {
func differenceOfSums(n int, m int) int {
	// nonDivisible := []int{};
	// divisible := []int{};

	nonDivisibleSum := 0;
	divisibleSum := 0;

	for i := 1; i <= n; i++ {
		if i % m == 0 {
			divisibleSum += i;
		} else {
			nonDivisibleSum += i;
		}
	}

	return nonDivisibleSum - divisibleSum;
}
	// nonDivisible := []int{};
	// divisible := []int{};

	nonDivisibleSum := 0;
	divisibleSum := 0;

	for i := 1; i <= n; i++ {
		if i % m == 0 {
			divisibleSum += i;
		} else {
			nonDivisibleSum += i;
		}
	}

	return nonDivisibleSum - divisibleSum;
}
