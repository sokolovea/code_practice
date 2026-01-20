from typing import List


class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        nums1_current: int = 0
        nums2_current: int = 0

        nums1_copy = nums1[0:m + 1]

        for i in range(n + m):
            if nums2_current >= n or (nums1_current < m and nums1_copy[nums1_current] <= nums2[nums2_current]):
                nums1[i] = nums1_copy[nums1_current]
                nums1_current += 1
            else:
                nums1[i] = nums2[nums2_current]
                nums2_current += 1
