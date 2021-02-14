//						猜数字
/* 1. 生成1-100的随机数，存储为目标数，让玩家猜测
 * 2. 玩家给出猜测值，存储其猜测值。
 * 3. 若猜测值<目标数，打印“猜低了”，若大于，猜高了
 * 4. 允许猜10次，且每次猜之前让玩家知道还剩多少次
 * 5. 若猜对了， 告诉：你猜对了;结束
 * 6. 若用完所有论次没猜对，告诉：没猜对，它是[目标数]
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() // 获取当前日期和时间的整数形式
	rand.Seed(seconds)           // 播种随机数生成器
	target := rand.Intn(100) + 1 // 若传递100, 你会得到0-99的随机数，但我们需要1-100,所以+1
	fmt.Println("I've chosen a random num between 1 and 100.")
	fmt.Println("Can u guess it?")

	reader := bufio.NewReader(os.Stdin) // 创建一个bufio.Reader，允许我们读取键盘输入

	success := false
	for i := 0; i < 10; i++ {
		fmt.Println("You have", 10-i, "guesses left.")

		//  读取输入
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// 剔除空白字符，并转换成int
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops.Your guess was LOW.")
		} else if guess > target{
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			fmt.Println("Good job! You guessed it!")
			success = true
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my num, it was:", target)
	}
}
