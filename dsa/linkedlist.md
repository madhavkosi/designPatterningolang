# Linked List Problems

## Contents
1. [Add Two Numbers](#add-two-numbers)
2. [Linked List Cycle](#linked-list-cycle)
3. [Merge Two Sorted Lists](#merge-two-sorted-lists)
4. [Copy List with Random Pointer](#copy-list-with-random-pointer)
5. [Reverse Linked List](#reverse-linked-list)
6. [Reverse Nodes in k-Group](#reverse-nodes-in-k-group)
7. [Reverse Linked List II](#reverse-linked-list-ii)
8. [Linked List Cycle II](#linked-list-cycle-ii)
9. [Remove Nth Node From End of List](#remove-nth-node-from-end-of-list)
10. [Remove Duplicates from Sorted List II](#remove-duplicates-from-sorted-list-ii)
11. [Rotate List](#rotate-list)
12. [LRU Cache](#lru-cache)
13. [Partition List](#partition-list)
14. [Intersection of Two Linked Lists](#intersection-of-two-linked-lists)
15. [Sort List](#sort-list)
16. [Merge k Sorted Lists](#merge-k-sorted-lists)

---

## Problem Details

###  Add Two Numbers

Given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

**Example:**

Input: `l1 = [2,4,3], l2 = [5,6,4]`  
Output: `[7,0,8]`  
Explanation: 342 + 465 = 807.

**Solution:**
```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode{}
    curr := dummyHead
    carry := 0

    for l1 != nil || l2 != nil || carry != 0 {
        x := 0
        if l1 != nil {
            x = l1.Val
            l1 = l1.Next
        }
        y := 0
        if l2 != nil {
            y = l2.Val
            l2 = l2.Next
        }

        sum := carry + x + y
        carry = sum / 10
        curr.Next = &ListNode{Val: sum % 10}
        curr = curr.Next
    }

    return dummyHead.Next
}
```

---

### Linked List Cycle

Given the head of a linked list, determine if the linked list has a cycle in it. Return true if there is a cycle in the linked list. Otherwise, return false.

**Example:**

Input: `head = [3,2,0,-4], pos = 1`  
Output: `true`  
Explanation: There is a cycle in the linked list, where the tail connects to the second node.

**Solution:**
```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}


func detectAndRemoveCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    slow, fast := head, head

    // Step 1: Detect the cycle using Floyd's Tortoise and Hare algorithm
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            break // Cycle detected
        }
    }

    if fast == nil || fast.Next == nil {
        return head // No cycle found, return the original list
    }

    // Step 2: Remove the cycle
    slow = head
    for slow != fast {
        slow = slow.Next
        fast = fast.Next
    }

    // Now, slow and fast are at the point where the cycle begins
    for fast.Next != slow {
        fast = fast.Next
    }

    fast.Next = nil // Remove the cycle by breaking the link

    return head
}
```

---

### Merge Two Sorted Lists

Given the heads of two sorted linked lists, merge the two lists in a sorted manner. Return the head of the merged linked list.

**Example:**

Input: `list1 = [1,2,4], list2 = [1,3,4]`  
Output: `[1,1,2,3,4,4]`

**Solution:**
```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy

    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            curr.Next = l1
            l1 = l1.Next
        } else {
            curr.Next = l2
            l2 = l2.Next
        }
        curr = curr.Next
    }

    if l1 != nil {
        curr.Next = l1
    } else {
        curr.Next = l2
    }

    return dummy.Next
}
```

---

### Copy List with Random Pointer

A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null. Construct a deep copy of the list.

**Example:**

Input: `head = [[7,null],[13,0],[11,4],[10,2],[1,0]]`  
Output: `[[7,null],[13,0],[11,4],[10,2],[1,0]]`

**Solution:**
```go
package main

type Node struct {
    Val    int
    Next   *Node
    Random *Node
}

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    // Creating a new weaved list of original and copied nodes.
    curr := head
    for curr != nil {
        newNode := &Node{Val: curr.Val}
        newNode.Next = curr.Next
        curr.Next = newNode
        curr = newNode.Next
    }

    // Assign the random pointers for the copy nodes.
    curr = head
    for curr != nil {
        if curr.Random != nil {
            curr.Next.Random = curr.Random.Next
        }
        curr = curr.Next.Next
    }

    // Unweave the linked list to get back the original linked list and the copied list.
    curr = head
    copyHead := head.Next
    for curr != nil {
        copy := curr.Next
        curr.Next = copy.Next
        if copy.Next != nil {
            copy.Next = copy.Next.Next
        }
        curr = curr.Next
    }

    return copyHead
}
```

---

### Reverse Linked List

Given the head of a singly linked list, reverse the list, and return the reversed list.

**Example:**

Input: `head = [1,2,3,4,5]`  
Output: `[5,4,3,2,1]`

**Solution:**
```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseListRecursive(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := reverseListRecursive(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}

// Iterative solution to reverse a linked list
func reverseListIterative(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head

    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    return prev
}
```

---

### Reverse Nodes in k-Group

Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.

**Example:**

Input: `head = [1,2,3,4,5], k = 2`  
Output: `[2,1,4,3,5]`

**Solution:**
```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseList(first, last *ListNode) *ListNode {
    prev := last
    for first != last {
        next := first.Next
        first.Next = prev
        prev = first
        first = next
    }
    return prev
}

// getKthNode returns the k-th node from the current node
func getKthNode(curr *ListNode, k int) *ListNode {
    for curr != nil && k > 0 {
        curr = curr.Next
        k--
    }
    return curr
}

// reverseKGroup reverses nodes of a linked list k at a time and returns the modified list.
func reverseKGroup(head *ListNode, k int) *ListNode {
    kthNode := getKthNode(head, k)
    if kthNode == nil {
        return head
    }   
    newHead := reverseList(head, kthNode)
    head.Next = reverseKGroup(kthNode, k)
    return newHead
}

```

---

### Reverse Linked List II

Given the head of a singly linked list and two integers left and right where left ≤ right, reverse the nodes of the list from position left to position right, and return the reversed list.

**Example:**

Input: `head = [1,2,3,4,5], left = 2, right = 4`  
Output: `[1,4,3,2,5]`

**Solution:**
```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if head == nil {
        return nil
    }

    var prev, curr *ListNode = nil, head
    for left > 1 {
        prev = curr
        curr = curr.Next
        left--
        right--
    }

    con, tail := prev, curr

    for right > 0 {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
        right--
    }

    if con != nil {
        con.Next = prev
    } else {
        head = prev
    }

    tail.Next = curr
    return head
}
```

---

### Remove Nth Node From End of List

Given the head of a linked list, remove the nth node from the end of the list and return its head.

**Example:**

Input: `head = [1,2,3,4,5], n = 2`  
Output: `[1,2,3,5]`

**Solution:**
```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    slow, fast := dummy, dummy

    for i := 0; i <= n; i++ {
        fast = fast.Next
    }

    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }

    slow.Next = slow.Next.Next
    return dummy.Next
}
```

---

### Remove Duplicates from Sorted List II

Given the head of a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct numbers from the original list. Return the linked list sorted as well.

**Example:**

Input: `head = [1,2,3,3,4,4,5]`  
Output: `[1,2,5]`

**Solution:**
```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
    dummy := &ListNode{Next: head}
    pred := dummy

    for head != nil {
        if head.Next != nil && head.Val == head.Next.Val {
            for head.Next != nil && head.Val == head.Next.Val {
                head = head.Next
            }
            pred.Next = head.Next
        } else {
            pred = pred.Next
        }
        head = head.Next
    }

    return dummy.Next
}
```

---

### Rotate List

Given the head of a linked list, rotate the list to the right by k places.

**Example:**

Input: `head = [1,2,3,4,5], k = 2`  
Output: `[4,5,1,2,3]`

**Solution:**
```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }

    length := 1
    lastElement := head
    for lastElement.Next != nil {
        lastElement = lastElement.Next
        length++
    }

    k = k % length
    if k == 0 {
        return head
    }

    lastElement.Next = head
    tempNode := head
    for i := 0; i < length-k-1; i++ {
        tempNode = tempNode.Next
    }

    newHead := tempNode.Next
    tempNode.Next = nil
    return newHead
}
```

---

### LRU Cache

Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

**Example:**

Input: `["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"] [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]`  
Output: `[null, null, null, 1, null, -1, null, -1, 3, 4]`

**Solution:**
```go
package main

type LRUCache struct {
    capacity int
    cache    map[int]*ListNode
    head     *ListNode
    tail     *ListNode
}

type ListNode struct {
    key  int
    val  int
    prev *ListNode
    next *ListNode
}

func Constructor(capacity int) LRUCache {
    head := &ListNode{}
    tail := &ListNode{}
    head.Next = tail
    tail.Prev = head
    return LRUCache{
        capacity: capacity,
        cache:    make(map[int]*ListNode),
        head:     head,
        tail:     tail,
    }
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok {
        this.moveToHead(node)
        return node.Val
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.cache[key]; ok {
        node.Val = value
        this.moveToHead(node)
    } else {
        if len(this.cache) == this.capacity {
            removed := this.removeTail()
            delete(this.cache, removed.Key)
        }
        node := &ListNode{Key: key, Val: value}
        this.cache[key] = node
        this.addToHead(node)
    }
}

func (this *LRUCache) moveToHead(node *ListNode) {
    this.removeNode(node)
    this.addToHead(node)
}

func (this *LRUCache) removeNode(node *ListNode) {
    node.Prev.Next = node.Next
    node.Next.Prev = node.Prev
}

func (this *LRUCache) addToHead(node *ListNode) {
    node.Prev = this.head
    node.Next = this.head.Next
    this.head.Next.Prev = node
    this.head.Next = node
}

func (this *LRUCache) removeTail() *ListNode {
    removed := this.tail.Prev
    this.removeNode(removed)
    return removed
}
```

---

### Partition List

Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x. You should preserve the original relative order of the nodes in each of the two partitions.

**Example:**

Input: `head = [1,4,3,2,5,2], x = 3`  
Output: `[1,2,2,4,3,5]`

**Solution:**
```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
    beforeHead := &ListNode{}
    before := beforeHead
    afterHead := &ListNode{}
    after := afterHead

    for head != nil {
        if head.Val < x {
            before.Next = head
            before = before.Next
        } else {
            after.Next = head
            after = after.Next
        }
        head = head.Next
    }

    after.Next = nil
    before.Next = afterHead.Next
    return beforeHead.Next
}
```

---

### Intersection of Two Linked Lists

Given the head of two linked lists, return the node where the two lists intersect. If there is no intersection, return null.

**Example:**

Input: `headA = [4,1,8,4,5], headB = [5,0,1,8,4,5]`  
Output: `8`

**Solution:**
```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }

    a, b := headA, headB
    for a != b {
        if a == nil {
            a = headB
        } else {
            a = a.Next
        }
        if b == nil {
            b = headA
        } else {
            b = b.Next
        }
    }

    return a
}
```

---

### Sort List

Given the head of a linked list, return the list after sorting it in ascending order.

**Example:**

Input: `head = [4,2,1,3]`  
Output: `[1,2,3,4]`

**Solution:**
```go
package main

type ListNode struct {
    Val  int
    Next *ListNode


}

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    mid := getMid(head)
    left := sortList(head)
    right := sortList(mid)
    return merge(left, right)
}

func getMid(head *ListNode) *ListNode {
    var midPrev *ListNode
    for head != nil && head.Next != nil {
        midPrev = head
        head = head.Next.Next
    }
    mid := midPrev.Next
    midPrev.Next = nil
    return mid
}

func merge(left, right *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy
    for left != nil && right != nil {
        if left.Val < right.Val {
            tail.Next = left
            left = left.Next
        } else {
            tail.Next = right
            right = right.Next
        }
        tail = tail.Next
    }
    if left != nil {
        tail.Next = left
    } else {
        tail.Next = right
    }
    return dummy.Next
}
```

---

### Merge k Sorted Lists

Given an array of k linked-lists, each linked-list is sorted in ascending order. Merge all the linked-lists into one sorted linked-list and return it.

**Example:**

Input: `lists = [[1,4,5],[1,3,4],[2,6]]`  
Output: `[1,1,2,3,4,4,5,6]`

**Solution:**
```go
package main

type ListNode struct {
    Val  int
    Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

    return mergeRange(lists, 0, len(lists)-1)
}

func mergeRange(lists []*ListNode, start, end int) *ListNode {
    if start == end {
        return lists[start]
    }
    mid := start + (end-start)/2
    left := mergeRange(lists, start, mid)
    right := mergeRange(lists, mid+1, end)
    return mergeTwoLists(left, right)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy

    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            curr.Next = l1
            l1 = l1.Next
        } else {
            curr.Next = l2
            l2 = l2.Next
        }
        curr = curr.Next
    }

    if l1 != nil {
        curr.Next = l1
    } else {
        curr.Next = l2
    }

    return dummy.Next
}



// PriorityQueue implements heap.Interface and holds the list nodes
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

// mergeKLists merges k sorted linked lists and returns the merged sorted list
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }

    pq := &PriorityQueue{}
    heap.Init(pq)

    // Push the head of each list onto the priority queue
    for _, list := range lists {
        if list != nil {
            heap.Push(pq, list)
        }
    }

    dummy := &ListNode{}
    current := dummy

    for pq.Len() > 0 {
        // Pop the smallest element from the heap
        smallest := heap.Pop(pq).(*ListNode)
        current.Next = smallest
        current = current.Next

        // Push the next element from the same list onto the heap
        if smallest.Next != nil {
            heap.Push(pq, smallest.Next)
        }
    }

    return dummy.Next
}
```
