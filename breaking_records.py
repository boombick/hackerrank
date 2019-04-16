#!/usr/local/bin/python3

import os


# Complete the breakingRecords function below.
def breaking_records(scores):
    best = [0, 0]   # best_val, count
    worst = [0, 0]  # worst_val, count
    prev = 0
    for k, x in enumerate(scores):
        print("====== STEP: {} ========" . format(k))
        print("Current val is {} Prev val is {}" . format(x, prev))
        if x > best[0] and k != 0:
            best[1] += 1
            best[0] = x
            print("Best increased")
        elif prev == 0:
            best[0] = worst[0] = x
        if x < worst[0] and k != 0:
            worst[1] += 1
            worst[0] = x
            print("Worst increased")

        prev = x
        print("======")
    print("---------------------------------")
    print("Best: {} Worst: {}" . format(best[1], worst[1]))
    return best[1], worst[1]


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    scores = list(map(int, input().rstrip().split()))

    result = breaking_records(scores)

    fptr.write(' '.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
