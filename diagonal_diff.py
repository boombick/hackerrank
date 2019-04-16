#!/usr/local/bin/python3

# https://www.hackerrank.com/challenges/diagonal-difference/problem

import os


# Complete the diagonalDifference function below.
def diagonal_difference(arr):
    return True


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    arr = []

    for _ in range(n):
        arr.append(list(map(int, input().rstrip().split())))

    result = diagonal_difference(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
