# Conway's Game of Life ^_^
import numpy as np
from time import sleep

alive = '\u25A0'
dead = '\u25A1'

def board_to_string(board: np.array) -> str:
    l = '\n'.join([''.join(x) for x in board]) + '\n'
    return l

def rand_board(size, prob) -> np.array:
    a = np.random.choice(a=[0,1],size=(size, size), p = [1-prob, prob])
    a = np.where(a==0,dead,a)
    a = np.where(a=='1',alive,a)
    return a

def deter_board(size, *places):
    r = np.full((size,size),dead)
    for i,j in places:
        r[i,j] = alive
    return r

def get(board: np.array, i: int, j: int) -> int:
    n,m = board.shape
    if i>=0 and j>=0 and i<n and j<m:
        return board[i,j]
    return None

def count_around(board: np.array, i: int, j: int) -> int:
    c=0
    for k in [i-1,i,i+1]:
        for t in [j-1,j,j+1]:
            if k==i and t==j:
                continue
            if get(board,k,t)==alive:
                c+=1
    return c

def decide(board: np.array, i: int, j: int) -> str:
    n = count_around(board,i,j)
    if board[i,j]==alive:
        return alive if n in [2,3] else dead
    else:
        return alive if n==3 else dead

def step(board: np.array) -> np.array:
    size = len(board)
    res = np.copy(board)
    for i in range(size):
        for j in range(size):
            res[i,j] = decide(board,i,j)
    return res

if __name__ == '__main__':
    size = int(input("Enter size for the board: "))
    prob = float(input("Enter probability for a cell to be alive at the beggining: "))
    interval = int(input("Enter time in miliseconds between updates: "))
    
    a = rand_board(size, prob)
    while(1):
        print(board_to_string(a))
        a = step(a)
        sleep(interval/1000)
