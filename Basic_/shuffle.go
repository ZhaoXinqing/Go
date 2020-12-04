

func shuffle(a []int) []int {
	rand.Seed(time.Now().UnixNano()) //设置种子

	rand.Shuffle(len(a), func(i, j int) { //调用算法
		a[i], a[j] = a[j], a[i]
	})
}


func shuffle(arr) {
    let i = arr.length;
    while (i) {
        let j = Math.floor(Math.random() * i--);
        [arr[j], arr[i]] = [arr[i], arr[j]];
    }
    return arr;
}

func main() {
	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	shuffle1 := shuffle(i)

	fmt.Println(shuffle1)
}

// 高校洗牌
// https://studygolang.com/articles/24498?fr=sidebar
