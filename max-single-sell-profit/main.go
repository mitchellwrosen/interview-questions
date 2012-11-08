package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func minIndex(arr ...int) (index int) {
	min := math.MaxInt32
	for i, val := range arr {
		if val < min {
			min = val
			index = i
		}
	}
	return
}

func maxIndex(arr ...int) (index int) {
	max := -1
	for i, val := range arr {
		if val > max {
			max = val
			index = i
		}
	}
	return
}

/*
Simple divide-and-conquer.
	prices: slice of ints representing stock prices
	buy: index into |prices| to buy stock
	sell: index into |prices| to sell stock
	profit: profit earned per share if stock is bought on |buy| and sold on
		|sell|
Solve each half of |prices| recursively. The optimal buy/sell is then either
located in the first half, second half, or |buy| is in the first half and
|sell| is in the second half. If this is the case, then |buy| is simply the
index of the minimum of the first half of |prices|, and |sell| is the index of
the maximum of the second half of |prices|.
*/
func GetOptimalBuySell1(prices []int) (int, int) {
	// If there is only one price, buy and sell on the same day is optimal.
	if len(prices) == 1 {
		return 0, 0
	}

	halfLen := len(prices) / 2
	firstHalf := prices[:halfLen]
	secondHalf := prices[halfLen:]

	buy1, sell1 := GetOptimalBuySell1(firstHalf)

	buy2, sell2 := GetOptimalBuySell1(secondHalf)
	buy2, sell2 = buy2+halfLen, sell2+halfLen

	buy3, sell3 := minIndex(firstHalf...), halfLen+maxIndex(secondHalf...)

	profit1 := prices[sell1] - prices[buy1]
	profit2 := prices[sell2] - prices[buy2]
	profit3 := prices[sell3] - prices[buy3]

	if profit1 > profit2 {
		if profit1 > profit3 {
			return buy1, sell1
		}
		return buy3, sell3
	}
	if profit2 > profit3 {
		return buy2, sell2
	}
	return buy3, sell3
}

/*
Slightly smarter divide-and-conquer.
	prices: slice of ints representing stock prices
	buy: index into |prices| to buy stock
	sell: index into |prices| to sell stock
	min: index into |prices| of the minimum price
	max: index into |prices| of the maximum price
	profit: profit earned per share if stock is bought on |buy| and sold on
		|sell|
Solve each half of |prices| recursively. The optimal buy/sell is then either
located in the first half, second half, or |buy| is in the first half and
|sell| is in the second half. If this is the case, then |buy| is simply the
index of the minimum of the first half of |prices|, and |sell| is the index of
the maximum of the second half of |prices|.

Since we're recursively calculating |buy| and |sell|, we might as well also
return |min| and |max|, so that the merging function doesn't need to re-iterate
over each half and determine either the min or the max.
*/
func GetOptimalBuySell2(prices []int) (buy, sell, min, max int) {
	// If there is only one price, buy and sell on the same day is optimal.
	if len(prices) == 1 {
		return 0, 0, 0, 0
	}

	halfLen := len(prices) / 2
	firstHalf := prices[:halfLen]
	secondHalf := prices[halfLen:]

	buy1, sell1, min1, max1 := GetOptimalBuySell2(firstHalf)

	buy2, sell2, min2, max2 := GetOptimalBuySell2(secondHalf)
	buy2, sell2, min2, max2 = buy2+halfLen, sell2+halfLen, min2+halfLen,
		max2+halfLen

	buy3, sell3 := min1, max2
	min, max = min1, max1 // Default to first half, possibly overwrite
	if prices[min2] < prices[min1] {
		min = min2
	}
	if prices[max2] > prices[max1] {
		max = max2
	}

	profit1 := prices[sell1] - prices[buy1]
	profit2 := prices[sell2] - prices[buy2]
	profit3 := prices[sell3] - prices[buy3]

	if profit1 > profit2 {
		if profit1 > profit3 {
			return buy1, sell1, min, max
		}
		return buy3, sell3, min, max
	}
	if profit2 > profit3 {
		return buy2, sell2, min, max
	}
	return buy3, sell3, min, max
}

func GetOptimalBuySell3(prices []int) (buy, sell int) {
	min := 0    // Index of min price in [0..k]
	profit := 0 // Best profit so far

	for i, val := range prices[1:] {
		if val-prices[min] > profit {
			profit = val - prices[min]
			buy = min
			sell = i + 1
		} else if val < prices[min] {
			min = i + 1
		}
	}
	return
}

func main() {
	// Initialize the array.
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Initializing prices.")
	prices := make([]int, 100000000)
	for i := range prices {
		prices[i] = rand.Int() % 100000000
	}

	//fmt.Println(prices)

	fmt.Println("Method 1: Divide and conquer")
	buy, sell := GetOptimalBuySell1(prices)
	fmt.Printf("Buy on %d ($%d), sell on %d ($%d) (profit $%d)\n", buy,
		prices[buy], sell, prices[sell], prices[sell]-prices[buy])

	fmt.Println("Method 2: Divide and conquer++")
	buy, sell, _, _ = GetOptimalBuySell2(prices)
	fmt.Printf("Buy on %d ($%d), sell on %d ($%d) (profit $%d)\n", buy,
		prices[buy], sell, prices[sell], prices[sell]-prices[buy])

	fmt.Println("Method 3: Dynamic programming")
	buy, sell = GetOptimalBuySell3(prices)
	fmt.Printf("Buy on %d ($%d), sell on %d ($%d) (profit $%d)\n", buy,
		prices[buy], sell, prices[sell], prices[sell]-prices[buy])

}
