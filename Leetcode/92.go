

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        // 保存当前head.Next节点，防止重新赋值后被覆盖
        // 一轮之后状态：nil<-1 2->3->4
        //              prev   head
        temp := head.Next
        head.Next = prev
        // pre 移动
        prev = head
        // head 移动
        head = temp
    }
    return prev
}

	head -> next
	
	[]  -> []  -> []





class Solution {
	public:
		//思路：head表示需要反转的头节点，pre表示需要反转头节点的前驱节点
		//我们需要反转n-m次，我们将head的next节点移动到需要反转链表部分的首部，需要反转链表部分剩余节点依旧保持相对顺序即可
		//比如1->2->3->4->5,m=1,n=5
		//第一次反转：1(head) 2(next) 3 4 5 反转为 2 1 3 4 5
		//第二次反转：2 1(head) 3(next) 4 5 反转为 3 2 1 4 5
		//第三次发转：3 2 1(head) 4(next) 5 反转为 4 3 2 1 5
		//第四次反转：4 3 2 1(head) 5(next) 反转为 5 4 3 2 1
		ListNode* reverseBetween(ListNode* head, int m, int n) {
			ListNode *dummy=new ListNode(-1);
			dummy->next=head;
			ListNode *pre=dummy;
			for(int i=1;i<m;++i)pre=pre->next;
			head=pre->next;
			for(int i=m;i<n;++i){
				ListNode *nxt=head->next;
				//head节点连接nxt节点之后链表部分，也就是向后移动一位
				head->next=nxt->next;
				//nxt节点移动到需要反转链表部分的首部
				nxt->next=pre->next;
				pre->next=nxt;//pre继续为需要反转头节点的前驱节点
			}
			return dummy->next;
		}
	};


	func reverseBetween(head *ListNode, m int, n int) *ListNode {
		// 思路：先遍历到m处，翻转，再拼接后续，注意指针处理
		// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
		if head == nil {
			return head
		}
		// 头部变化所以使用dummy node
		dummy := &ListNode{Val: 0}
		dummy.Next = head
		head = dummy
		// 最开始：0->1->2->3->4->5->nil
		var pre *ListNode
		var i = 0
		for i < m {
			pre = head
			head = head.Next
			i++
		}
		// 遍历之后： 1(pre)->2(head)->3->4->5->NULL
		// i = 1
		var j = i
		var next *ListNode
		// 用于中间节点连接
		var mid = head
		for head != nil && j <= n {
			// 第一次循环： 1 nil<-2 3->4->5->nil
			temp := head.Next
			head.Next = next
			next = head
			head = temp
			j++
		}
		// 循环需要执行四次
		// 循环结束：1 nil<-2<-3<-4 5(head)->nil
		pre.Next = next
		mid.Next = head
		return dummy.Next
	}



	class Solution {
		public ListNode reverseBetween(ListNode head, int m, int n) {
	
			// Empty list
			if (head == null) {
				return null;
			}
	
			// Move the two pointers until they reach the proper starting point
			// in the list.
			ListNode cur = head, prev = null;
			while (m > 1) {
				prev = cur;
				cur = cur.next;
				m--;
				n--;
			}
	
			// The two pointers that will fix the final connections.
			ListNode con = prev, tail = cur;
	
			// Iteratively reverse the nodes until n becomes 0.
			ListNode third = null;
			while (n > 0) {
				third = cur.next;
				cur.next = prev;
				prev = cur;
				cur = third;
				n--;
			}
	
			// Adjust the final connections as explained in the algorithm
			if (con != null) {
				con.next = prev;
			} else {
				head = prev;
			}
	
			tail.next = cur;
			return head;
		}
	}
	