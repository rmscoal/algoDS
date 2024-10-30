def three_sum_time_limit_exceeded(nums: list[int]) -> list[list[int]]:
    '''
    Given an integer array nums, return all the triplets [nums[i], nums[j],
    nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

    Ex. [-1,0,1,2,-1,-4]

    Hint 1:
    So, we essentially need to find three numbers x, y, and z such that
    they add up to the given value. If we fix one of the numbers say x, we are left with the two-sum problem at hand!
    '''
    N = len(nums)
    results = []
    for i in range(N):
        required_sum = -nums[i]
        pairs = {}

        for j in range(i + 1, N):
            if nums[j] in pairs:
                triplets = [nums[i], nums[j], pairs.get(nums[j])]
                triplets.sort()
                found = False
                for triplet in results:
                    if triplet == triplets:
                        found = True
                if not found:
                    results.append(triplets)
                pairs.pop(nums[j])
            else:
                complement = required_sum - nums[j]
                pairs[complement] = nums[j]

    return results


def three_sum(nums: list[int]):
    '''
    Given an integer array nums, return all the triplets [nums[i], nums[j], 
    nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

    Ex. [-1,0,1,2,-1,-4]

    1. Sort the list in ascending order, [-1,0,1,2,-1,-4] -> [-4,-1,-1,0,1,2]
    2. Set the target to -nums[i]
    3. Explore about two-sum.
    4. Set the left and right pointer to be i + 1 and N - 1
    5. If the both the two-sum is equal to the target, add as result.
    6. Adjust the left and right pointer depending on the result of comparison with target.
    7. To avoid triplets duplication, skip the iteration where nums[i] == nums[i - 1] and
    during the iteration skip where nums[left] == nums[left-1] or nums[right] == nums[right + 1]
    '''
    N = len(nums)
    results = []
    nums.sort()
    for i in range(N):
        if i > 0 and nums[i] == nums[i-1]:
            continue

        target = -nums[i]
        left, right = i + 1, N - 1
        prev_left, prev_right = -1, -1
        while left < right:
            if prev_left > -1 and nums[left] == nums[prev_left]:
                prev_left = left
                left += 1
                continue
            if prev_right > -1 and nums[right] == nums[prev_right]:
                prev_right = right
                right -= 1
                continue

            two_sum = nums[left] + nums[right]
            if two_sum == target:
                results.append([nums[i], nums[left], nums[right]])
                prev_left = left
                prev_right = right
                left += 1
                right -= 1
            elif two_sum < target:
                prev_left = left
                left += 1
            else:
                prev_right = right
                right -= 1
    return results


if __name__ == '__main__':
    print(three_sum([-1, 0, 1, 2, -1, -4]))
