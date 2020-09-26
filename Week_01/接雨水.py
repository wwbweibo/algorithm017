class Solution:
    def trap(self, height: List[int]) -> int:
        stack = []
        idx = 0
        ans = 0
        for i in height:
            while len(stack) > 0 and height[stack[-1]] < i :
                value = stack.pop()
                if len(stack) == 0:
                    break
                dis = idx - stack[-1] - 1
                h = min(i, height[stack[-1]]) - height[value] 
                ans = ans + (dis * h)
            stack.append(idx)
            idx = 1 + idx
        return ans