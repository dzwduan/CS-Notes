func climbStairs(n int) int {
    if n==0 || n==1{
        return 1
    }
    p,q,r :=0,0,1
    for i:=1;i<=n;i++{
        p=q
        q=r
        r=p+q
    }
    return r
}