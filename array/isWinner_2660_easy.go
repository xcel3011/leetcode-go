package array

func isWinner(player1 []int, player2 []int) int {
	n := len(player1)
	diff := 0

	for i := 0; i < n; i++ {
		score1, score2 := player1[i], player2[i]
		if (i > 0 && player1[i-1] == 10) || (i > 1 && player1[i-1] == 10) {
			score1 = score1 * 2
		}
		if (i > 0 && player2[i-1] == 10) || (i > 1 && player2[i-1] == 10) {
			score2 = score2 * 2
		}
		diff += score1 - score2
	}
	if diff > 0 {
		return 1
	} else if diff < 0 {
		return 2
	} else {
		return 0
	}
}
