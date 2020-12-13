func commonChars(a []string) (ans []string) {
    minFreq := [26]int{}  //数组
    for i := range minFreq {
        minFreq[i] = math.MaxInt64  //数组每一位设为最大
    }
    for _, word := range a {  //string拆成单个的单词
        freq := [26]int{}     //数组，每次刷新
        for _, b := range word {
            freq[b-'a']++     //对单词里面的单个字符计数
        }
        for i, f := range freq[:] { //i代表-a的下标，f代表计数值
            if f < minFreq[i] {
                minFreq[i] = f
            }
        }
    }
    for i := byte(0); i < 26; i++ {   
        for j := 0; j < minFreq[i]; j++ { //j代表要打印几次，大部分是0次
            ans = append(ans, string('a'+i))
        }
    }
    return
}

