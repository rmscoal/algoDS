class BinarySquare:
  @classmethod
  def Solver(n):
    i = 0
    odd = True
    while (i < n):
      j = 0
      while (j < n):
          if (odd):
              if (j % 2 == 0):
                  print(0, end='')
              else:
                  print(1, end='')
              j += 1
          else:
              j += 1
              if (j % 2 == 0):
                  print(0, end='')
              else:
                  print(1, end='')

        odd = not odd
        i += 1
        print()