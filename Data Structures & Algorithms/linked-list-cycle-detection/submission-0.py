# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        seen = {head: True}

        while head:
            if head.next in seen:
                return True
            
            head = head.next
            seen[head] = True

        return False