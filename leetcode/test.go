package main

import (
	"fmt"
	"strconv"
	// "strings"
)

func main() {
	ans := maxA2(3)
	fmt.Println(ans)
}

func maxA(n int) int{
	memo := map[string]int{}
	return dp(n, 0, 0, memo)
}

func dp(n, printNum, cacheNum int, memo map[string]int) int {
	if n <= 0 {
		return printNum
	}
	theKey := strconv.Itoa(n) + "-" + strconv.Itoa(printNum) + "-" + strconv.Itoa(cacheNum)
	if value, ok := memo[theKey]; ok {
		return value
	}
	res := max(dp(n-1, printNum+1, cacheNum, memo), dp(n-1, printNum+cacheNum, cacheNum, memo), dp(n-2, printNum, printNum, memo))
	memo[theKey] = res
	return res
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v > res {
			res = v
		}
	}
	return res
}

func maxA2(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = i
		for j := 2; j < i; j++ {
			dp[i] = max(dp[i], dp[j-2]*(i-j+1))
		}
	}
	return dp[n]
}