export default class StayingIngrid {
    /**
     * Solve by straight forward and brute-force method.
     * @param n grid of n x n
     * @param startPos starting position of the robot
     * @param s instruction to take
     */
    static Solver(n, startPos, s) {
        // Returned variable:
        const answer = [];
        for (let i = 0; i < s.length; i += 1) {
            const pos = startPos;
            for (let j = i; j < s.length; j += 1) {
                switch (s[j]) {
                    case 'R': {
                        pos[1] += 1;
                        break;
                    }
                    case 'L': {
                        pos[1] -= 1;
                        break;
                    }
                    case 'D': {
                        pos[0] += 1;
                        break;
                    }
                    default: {
                        pos[0] -= 1;
                        break;
                    }
                }
                if (pos[0] > n - 1 || pos[0] < 0 || pos[1] > n - 1 || pos[1] < 0) {
                    break;
                }
                answer[i] += 1;
            }
        }
        return answer;
    }
    /**
     * m is the length of command.
     * First let us think, how do we know when will the robot
     * move out of the boundary if the robot start at current step?
     * We know the start position of robot, say, (x,y) means row x
     * and col y , this is quite different than Cartesian coordinate,
     * please get used to it. If the robot move out of boundary, one
     * of the four situation must happen:
     * 1, it move up for x+1 steps
     * 2, it move down for n-x steps
     * 3, it move left for y+1 steps
     * 4, it move right for n-y steps
     *
     * Then, let us trace back from last command to first command.
     * We assume that the if the robot operate all the commands to
     * the end (whenever it starts), the final position is (0,0),
     * notice that this is virtual position, not real, just for
     * easy understanding.
     *
     * Suppose we have the command "LURD", at time t=4, the location
     * will be (0,0). If the robot start at t=3, and the command will
     * be "D", what is the initial virtual position? (-1,0). In other
     * words, at time t=3, the position will be (-1,0). If the robot
     * start at t=2, and the command will be "RD", what is the initial
     * virtual position? (-1,-1). At time t=2, the position will be
     * (-1,-1)
     *
     * We can find each virtual position at time t when the final
     * position without bound is (0,0), we keep track of each initial
     * vertical and horizontal position and record them in separate
     * dictionary.
     *
     * Then we can judge if we can walk until the end? Previouly, I
     * declare, in the future, if the robot walk up (x+1) steps, down
     * (n-x) steps, left (y+1) steps, right (n-y) steps, it will stop.
     * In other words, suppose (row,col) is the position at time t, if
     * row-(x+1), row+(n-x) appears in the horizonal dictionary, or col-(y+1),
     * col+(n-y) appears in vertical dictionary, it will stop at that time.
     * If multiple answer exists, find the most recent time since it happens
     * first, and the answer will be difference between most recent dictionary
     * time and current time. If there is no stop time, congratulations, it
     * will walk till final. This is a one-pass solution, with time complexity
     * O(m), no matter how big n is. Time is 109ms, beat 100% currently.
     */
    static FastSolver(n, startPos, s) {
        const m = s.length;
        const directionObj = {
            U: [-1, 0],
            D: [1, 0],
            R: [0, 1],
            L: [0, -1],
        };
        // These variables defines the upper limit moves the robot can take
        // for each directions.
        //
        // For example:
        // startPos = [0,1].
        // Hence the robot is out of boundary if it moves up for startPos[0] + 1.
        const limitXMovesForUpDir = startPos[0] + 1;
        const limitXMovesForDownDir = n - startPos[0];
        const limitYMovesForRightDir = n - startPos[1];
        const limitYMovesForLeftDir = startPos[1] + 1;
        // This variable stores the information of the future current column
        // as the key and the time (the iterator in the loop) as the value.
        const futureHorizontalPosOnStepsObj = { 0: m };
        // This variable stores the information of the future current row
        // as the key and the time (the iterator in the loop) as the value.
        const futureVerticalPosOnStepsObj = { 0: m };
        // Assume that the startPosition is at 0,0 since we are doing it backward.
        let currColumn = 0;
        let currRow = 0;
        // Returned Variable:
        const answer = [];
        // Let t here denote the time since instructions are done one-at-a-time.
        for (let t = m - 1; t >= 0; t -= 1) {
            currRow -= directionObj[s[t]][0];
            currColumn -= directionObj[s[t]][1];
            // Holds the maximum instructions the robot can do without going outside
            // the boundary.
            let maxNoOfInstructions = m - t;
            if (futureHorizontalPosOnStepsObj[currRow + limitXMovesForDownDir] !== undefined) {
                maxNoOfInstructions = Math.min(maxNoOfInstructions, futureHorizontalPosOnStepsObj[currRow + limitXMovesForDownDir] - t - 1);
            }
            if (futureHorizontalPosOnStepsObj[currRow - limitXMovesForUpDir] !== undefined) {
                maxNoOfInstructions = Math.min(maxNoOfInstructions, futureHorizontalPosOnStepsObj[currRow - limitXMovesForUpDir] - t - 1);
            }
            if (futureVerticalPosOnStepsObj[currColumn + limitYMovesForRightDir] !== undefined) {
                maxNoOfInstructions = Math.min(maxNoOfInstructions, futureVerticalPosOnStepsObj[currRow + limitYMovesForRightDir] - t - 1);
            }
            if (futureVerticalPosOnStepsObj[currRow - limitYMovesForLeftDir] !== undefined) {
                maxNoOfInstructions = Math.min(maxNoOfInstructions, futureVerticalPosOnStepsObj[currRow - limitYMovesForLeftDir] - t - 1);
            }
            futureHorizontalPosOnStepsObj[currRow] = t;
            futureVerticalPosOnStepsObj[currColumn] = t;
            answer.push(maxNoOfInstructions);
        }
        return answer.reverse();
    }
}
//# sourceMappingURL=stayingingrid.js.map