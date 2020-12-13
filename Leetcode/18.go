import "sort"
func fourSum(nums []int, target int) [][]int {
	ans :=[][]int{}
	counter :=map[int]int{}

	for _,value :=range nums{
		counter[value]++
	}

	uniqNums := []int{}
	for key:= range counter{
		uniqNums = append(uniqNums,key)
	}

	sort.Ints(uniqNums)


	for i:=0;i<len(uniqNums);i++{
		if(uniqNums[i]*4==target) && counter[uniqNums[i]]>=4{
			ans = append(ans,[]int{uniqNums[i],uniqNums[i],uniqNums[i],uniqNums[i]})
		}

		for j:=i+1;j<len(uniqNums);j++{
			if (uniqNums[i]*3+uniqNums[j]==target) && counter[uniqNums[i]]>=3 && counter[uniqNums[j]]>=1{
				ans = append(ans , []int{uniqNums[i],uniqNums[i],uniqNums[i],uniqNums[j]})
			}

			
			if (uniqNums[j]*3+uniqNums[i]==target) && counter[uniqNums[j]]>=3 && counter[uniqNums[i]]>=1{
				ans = append(ans , []int{uniqNums[i],uniqNums[j],uniqNums[j],uniqNums[j]})
			}
			//注意2+2也可以
			if (uniqNums[i]*2+uniqNums[j]*2==target) && counter[uniqNums[i]]>=2 && counter[uniqNums[j]]>=2{
				ans = append(ans , []int{uniqNums[i],uniqNums[i],uniqNums[j],uniqNums[j]})
			}
		
			for k:=j+1;k<len(nums);k++{
				if (uniqNums[i]*2+uniqNums[j]+uniqNums[k] == target) && counter[uniqNums[i]] > 1 {
					ans = append(ans, []int{uniqNums[i], uniqNums[i], uniqNums[j], uniqNums[k]})
				}


				
				if (uniqNums[i]+uniqNums[j]*2+uniqNums[k]==target) && counter[uniqNums[j]]>=2{
					ans = append(ans , []int{uniqNums[i],uniqNums[j],uniqNums[j],uniqNums[k]})
				}
				if (uniqNums[i]+uniqNums[j]+uniqNums[k]*2==target) && counter[uniqNums[k]]>=2{
					ans = append(ans , []int{uniqNums[i],uniqNums[j],uniqNums[k],uniqNums[k]})
				}

				c := target-uniqNums[i]-uniqNums[j]-uniqNums[k]
				if c > uniqNums[k] && counter[c] > 0{
					ans = append(ans, []int{uniqNums[i], uniqNums[j], uniqNums[k], c})
				}
	
			}   
		}
	}
    return ans
}