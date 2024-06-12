package leetcode_go

func accountBalanceAfterPurchase(purchaseAmount int) int {
	quotient := purchaseAmount / 10
	remainder := purchaseAmount % 10
	if remainder >= 5 {
		quotient++
	}
	return 100 - (quotient)*10
}
