func canJump(nums []int) bool {
    // 思路：看最后一跳
    // 状态：f[i] 表示是否能从0跳到i
    // 推导：f[i] = OR(f[j],j<i&&j能跳到i) 判断之前所有的点最后一跳是否能跳到当前点
    // 初始化：f[0] = 0
    // 结果： f[n-1]
    if len(nums) == 0 {
        return true
    }
    f := make([]bool, len(nums))
    f[0] = true
    for i := 1; i < len(nums); i++ {
        for j := 0; j < i; j++ {
            if f[j] == true && nums[j]+j >= i {
                f[i] = true
            }
        }
    }
    return f[len(nums)-1]
}