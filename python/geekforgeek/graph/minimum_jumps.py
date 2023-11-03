from typing import DefaultDict


def minimizeJumps(arr):
    n = len(arr)

    # Initialize a map for mapping element
    # with indices of all similar value
    # occurrences in array
    unmap = {}

    # Mapping element with indices
    for i in range(n):
        if arr[i] in unmap:
            unmap.get(arr[i]).append(i)
        else:
            unmap.update({arr[i]: [i]})

    q = []
    visited = [False]*n

    # Push starting element into queue
    # and mark it visited
    q.append(0)
    visited[0] = True

    # Initialize a variable count for
    # counting the minimum number number
    # of valid jump to reach at last index
    count = 0

    # Do while queue size is
    # greater than 0
    while (len(q) > 0):
        size = len(q)

        # Iterate on all the
        # elements of queue
        for i in range(size):
            # Fetch the front element and
            # pop out from queue
            curr = q[0]
            q.pop(0)

            # Check if we reach at the
            # last index or not if true,
            # then return the count
            if (curr == n - 1):
                return count

            # Check if curr + 1 is valid
            # position to visit or not
            if (curr + 1 < n and visited[curr + 1] == False):
                # If true, push curr + 1
                # into queue and mark
                # it as visited
                q.append(curr + 1)
                visited[curr + 1] = True

            # Check if curr - 1 is valid
            # position to visit or not
            if (curr - 1 >= 0 and visited[curr - 1] == False):
                # If true, push curr - 1
                # into queue and mark
                # it as visited
                q.append(curr - 1)
                visited[curr - 1] = True

            # Now, Iterate over all the
            # element that are similar
            # to curr
            if arr[curr] in unmap:
                for j in range(len(unmap[arr[curr]])):
                    child = unmap.get(arr[curr])[j]
                    if (curr == child):
                        continue

                    # Check if child is valid
                    # position to visit or not
                    if (visited[child] == False):
                        # If true, push child
                        # into queue and mark
                        # it as visited
                        q.append(child)
                        visited[child] = True

            # Erase all the occurrences
            # of curr from map. Because
            # we already considered these
            # element for valid jump
            # in above step
            if arr[curr] in unmap:
                unmap.pop(arr[curr])

        # Increment the count of jump
        count = count + 1

    # Finally, return the count.
    return count


if __name__ == "__main__":
    print("Result", minimizeJumps(
        [100, -23, -23, 404, 100, 23, 23, 23, 3, 404]))
    print("Result", minimizeJumps(
        [7, 6, 9, 6, 9, 6, 9, 7]))
