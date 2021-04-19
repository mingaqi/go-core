package practise

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

//请实现微信拼手气红包算法，输入红包总额，红包个数，输出红包抽取后的数值序列，算法需考虑公平性、随机性。
//题目答案要求：
//总和为金额数，个数为红包数，单个金额不小于0.01，取两位小数。

// 100元 10个红包  第一个人随机范围在 10*2
func TestHb(t *testing.T) {
	// 金额 红包个数
	var amount, count int64 = 100, 10
	// 验证数据
	var balance, sum int64 = amount, 0
	for i := int64(0); i < count; i++ {
		x := DoubleAvg(count-i, balance)
		fmt.Printf("第%d个红包:%.2f\n", i+1, float64(x)/float64(100))
		balance -= x
		sum += x

	}
	fmt.Println("总金额", float64(sum)/float64(100))
}

func DoubleAvg(count, amount int64) int64 {
	var min int64 = 1
	if count == 1 {
		return amount
	}
	// 计算最大可用金额
	max := amount - min*count
	// 最大可用平均值
	avg := max / count
	// 二倍均值加最小金额 防止随机
	maxAvg := avg*2 + min
	rand.Seed(time.Now().Unix())
	// 随机红包金额序列元素，把二倍均值作为随机的最大数
	x := rand.Int63n(maxAvg) + min
	return x
}

func lengthOfLongestSubstring(s string) int {
	start := 0
	end := 1
	for _, v := range s {
		if strings.Contains(s[start:end], string(v)) {
			start++
		} else {
			end++
		}
	}
	return end - start
}
