func longestPalindromeSubseq(s string) int {
    reversed := ""
    for i := len(s) - 1; i >= 0; i-- {
        reversed += string(s[i])
    }
	return longestCommonSubseq(s, reversed)
}

func longestCommonSubseq(s1, s2 string) int {
    dp := make([][]int, len(s1)+1)
    for i := range dp {
        dp[i] = make([]int, len(s2)+1)
    }

    for i := 0; i < len(s1); i++ {
        for j := 0; j < len(s2); j++ {
            if s1[i] == s2[j] {
                dp[i+1][j+1] = 1 + dp[i][j]
            } else {
                dp[i+1][j+1] = int(math.Max(float64(dp[i][j+1]), float64(dp[i+1][j])))
            }
        }
    }

    return dp[len(s1)][len(s2)]
}
