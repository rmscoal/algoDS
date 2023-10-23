from typing import List, Tuple


class Job:
    """
    Job class which stores profit and deadline.
    """

    def __init__(self, profit=0, deadline=0):
        self.profit = profit
        self.deadline = deadline
        self.id = 0


class Solution:
    def job_scheduling(self, Jobs: List[Job], n: int) -> Tuple[int, int]:
        """
        Note: Each job takes 1 unit time

        Problem Definition:
            https://www.youtube.com/watch?v=R6Skj4bT1HE
        """
        # Sort by profit
        Jobs.sort(key=lambda job: job.profit, reverse=True)

        # Job count and total profit
        count = 0
        max_profit = 0

        # Since we want to maximize profit, then we would want to find the
        # maximum deadlines so we can do many jobs.
        max_deadline = -1
        for i in range(n):
            if max_deadline < Jobs[i].deadline:
                max_deadline = Jobs[i].deadline

        # Create slots for jobs
        slots = [False] * max_deadline

        for job in Jobs:
            for i in range(job.deadline - 1, -1, -1):
                if slots[i] == False:
                    slots[i] = True
                    count += 1
                    max_profit += job.profit
                    break

        return count, max_profit


if __name__ == "__main__":
    jobs = [Job(20, 4), Job(10, 1), Job(40, 1), Job(30, 1)]
    count, profit = Solution().job_scheduling(jobs, 4)
    print("Count", count)
    print("Profit", profit)

    print()
