def prison(n,m,h,v):
    largest_h_gap = 1
    largest_v_gap = 1

    v_cell = [1] * m
    h_cell = [1] * n

    if len(h) == 0 and len(v) == 0:
        return largest_h_gap * largest_v_gap

    for vi in v:
        v_cell[vi-1] = 0
    for hi in h:
        h_cell[hi-1] = 0

    temp_gap = 1
    for i in range(len(v_cell)):
        if v_cell[i] == 1:
            temp_gap = 1
        if v_cell[i] == 0:
            temp_gap = temp_gap + 1
        if temp_gap > largest_v_gap:
            largest_v_gap = temp_gap

    temp_gap = 1
    for i in range(len(h_cell)):
        if h_cell[i] == 1:
            temp_gap = 1
        if h_cell[i] == 0:
            temp_gap = temp_gap + 1
            if temp_gap > largest_h_gap:
                largest_h_gap = temp_gap
        
    return largest_h_gap * largest_v_gap

if __name__ == "__main__":
    print("Result is", prison(10, 10, [2,4,6,8,10], [1,3,5,7,9]))
    print("Result is", prison(5, 5, [2,4], [1,3,5]))
