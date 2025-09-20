class Solution(object):
    def distance(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        index_map = {}
        result = [0] * len(nums)

        for i, num in enumerate(nums):
            if index_map.get(num):
                index_map[num].append(i)
            else:
                index_map[num] = [i]

        for indices in index_map.values():
            prefix_sum = [0] * len(indices)
            prefix_sum[0] = indices[0]

            for i in range(1, len(indices)):
                prefix_sum[i] = prefix_sum[i - 1] + indices[i]

            total = len(indices)
            for i in range(total):
                left = i
                right = total - i - 1

                left_sum = prefix_sum[i - 1] if i > 0 else 0
                right_sum = prefix_sum[-1] - prefix_sum[i]

                result[indices[i]] = (indices[i] * left - left_sum) + (right_sum - indices[i] * right)

        return result


if __name__ == "__main__":
    examples = [[1,3,1,1,2], [0,5,3]]
    for index, example in enumerate(examples):
        print(f"Example {index}\nInput: {example} \nOutput: {Solution().distance(example)}")
