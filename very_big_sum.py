#!/usr/local/bin/python3

# https://www.hackerrank.com/challenges/a-very-big-sum/problem

import os


# Complete the breakingRecords function below.
def big_sum(scores):


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    scores = list(map(int, input().rstrip().split()))

    result = big_sum(scores)

    fptr.write(' '.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
