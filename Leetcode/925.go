func isLongPressedName(name string, typed string) bool {
    i, j := 0, 0
    for j < len(typed) {
        if i < len(name) && name[i] == typed[j] {
            i++
            j++
        } else if j > 0 && typed[j] == typed[j-1] {
            j++
        } else {
            return false
        }
    }
    return i == len(name)
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/long-pressed-name/solution/chang-an-jian-ru-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。