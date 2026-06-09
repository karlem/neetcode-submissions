func maxProfit(prices []int) int {
    bestBuy := prices[0]
    var profit int

    for i := 1; i < len(prices); i++ {
        profit = max(profit, prices[i]-bestBuy)
        bestBuy = min(bestBuy, prices[i])
    }

    return profit
}
