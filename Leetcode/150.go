func evalRPN(tokens []string) int {
    if len(tokens)==0{
        return 0
    }
    stack:=make([]int,0)
    for i:=0;i<len(tokens);i++{
        switch tokens[i]{
        case "+","-","*","/":
            if len(stack)<2{       //注意！！！！！！！！
                return -1
            }
            // 注意：a为被除数，b为除数
            b:=stack[len(stack)-1]
            a:=stack[len(stack)-2]
            stack=stack[:len(stack)-2]         //注意！！！！！！！！！！
            var result int
            switch tokens[i]{
            case "+":
                result=a+b
            case "-":
                result=a-b
            case "*":
                result=a*b
            case "/":
                result=a/b
            }
            stack=append(stack,result)
        default:
            // 转为数字
            val,_:=strconv.Atoi(tokens[i])
            stack=append(stack,val)
        }
    }
    return stack[0]
}

//本质是switch套switch,原因是为了在外层判断合法性