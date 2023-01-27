class StayingInGrid:

  @classmethod
  def FastSolver(self, n, startPos, s):
    m = len(s)
    direc = {'U':[-1,0],'D':[1,0],'L':[0,-1],'R':[0,1]}
    upmost = startPos[0] + 1
    downmost = n - startPos[0]
    leftmost = startPos[1] + 1
    rightmost = n - startPos[1]
    curr_row,curr_col = 0,0    
    next_row,next_col = {0:m}, {0:m}
    ans = []

    print("startPos %a", startPos)
    print("Upmost %s", upmost)
    print("Downmost %s", downmost)
    print("Leftmost %s", leftmost)
    print("Rightmost %s", rightmost)
    
    for i in range(m-1,-1,-1):
        print("previous next_row ", next_row)
        print("previous next_col ", next_col)

        curr_row -= direc[s[i]][0]
        curr_col -= direc[s[i]][1]
        maxstep = m-i
        if curr_row - upmost in next_row:  
            maxstep = min(maxstep,  next_row[curr_row - upmost] - i - 1 )
        if curr_row + downmost in next_row:  
            maxstep = min(maxstep,  next_row[curr_row + downmost] - i - 1 )
        if curr_col - leftmost in next_col:  
            maxstep = min(maxstep,  next_col[curr_col - leftmost] - i - 1 )
        if curr_col + rightmost in next_col: 
            maxstep = min(maxstep,  next_col[curr_col + rightmost] - i - 1 )
        next_row[curr_row] = i
        next_col[curr_col] = i
        ans.append(maxstep)

        print(next_row)
        print(next_col)
        
    return ans[::-1]