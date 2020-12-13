func min(a int,b int) int{
	if a<b{
		return a
	}
	return b
}

//自底向上递推
func minimumTotal(triangle [][]int) int {
	row,col := len(triangle)-2,0
	for ;row>=0;row-- {
		for col=0;col<len(triangle[row]);col++{	 
		triangle[row][col] += min(triangle[row+1][col],triangle[row+1][col+1])
		}
	}
	return triangle[0][0]
}