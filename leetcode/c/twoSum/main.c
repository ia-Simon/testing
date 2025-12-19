#include <stdio.h>
#include <stdlib.h>

int* twoSum(int* nums, int numsSize, int target, int* returnSize);

int main() {
    int nums[] = {3, 2, 4};
    int numsSize = sizeof(nums) / sizeof(int);

    int indexesSize = 0;
    int *indexes = twoSum(nums, numsSize, 6, &indexesSize);

    for (int i = 0; i < indexesSize; i++) {
        printf("> %d\n", indexes[i]);
    }

    return 0;
}

int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
    *returnSize = 2;
    int *indexes = malloc(sizeof(int) * *returnSize);

    for (int i = 0; i < numsSize; i++) {
        for (int j = 0; j < numsSize; j++) {
            if (i != j && (nums[i] + nums[j] == target)) {
                indexes[0] = i;
                indexes[1] = j;
                return indexes;
            }
        }
    }

    return NULL;
}