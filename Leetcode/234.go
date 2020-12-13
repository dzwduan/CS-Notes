func isPalindrome(head *ListNode) bool {
    // 1 2 nil
    // 1 2 1 nil
    // 1 2 2 1 nil
    if head==nil{
        return true
    }
    slow:=head
    // fast如果初始化为head.Next则中点在slow.Next
    // fast初始化为head,则中点在slow
    fast:=head.Next
    for fast!=nil&&fast.Next!=nil{
        fast=fast.Next.Next
        slow=slow.Next
    }

    tail:=reverse(slow.Next)
    // 断开两个链表(需要用到中点前一个节点)
    slow.Next=nil
    for head!=nil&&tail!=nil{
        if head.Val!=tail.Val{
            return false
        }
        head=head.Next
        tail=tail.Next
    }
    return true

}

func reverse(head *ListNode)*ListNode{
    // 1->2->3
    if head==nil{
        return head
    }
    var prev *ListNode
    for head!=nil{
        t:=head.Next
        head.Next=prev
        prev=head
        head=t
    }
    return prev
}
