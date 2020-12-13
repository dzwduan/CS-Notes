/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func detectCycle(head *ListNode) *ListNode {
    //O(1)放弃
    seen := map[*ListNode] struct {} {}

    for head!=nil{
        if _,ok :=seen[head] ; ok{
            return head
        }
        seen[head] = struct{}{}    //注意空类型
        head = head.Next
    }
    return nil
}