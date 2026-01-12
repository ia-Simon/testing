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

/** O(n^2) */
// int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
//     *returnSize = 2;
//     int *indexes = malloc(sizeof(int) * *returnSize);

//     for (int i = 0; i < numsSize; i++) {
//         for (int j = 0; j < numsSize; j++) {
//             if (i != j && (nums[i] + nums[j] == target)) {
//                 indexes[0] = i;
//                 indexes[1] = j;
//                 return indexes;
//             }
//         }
//     }

//     return NULL;
// }

/** O(?) */
int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
    for (int i = 0; i < numsSize; i++) {
        for (int j = i+1; j < numsSize; j++) {
            if (nums[i] + nums[j] == target) {
                int *indexes = malloc(2 * sizeof(int));
                indexes[0] = i;
                indexes[1] = j;
                *returnSize = 2;
                return indexes;
            }
        }
    }

    *returnSize = 0;
    return NULL;
}

/** O(n log n) */
// typedef struct {
//     int value;
//     int idx;
// } Pair;

// int cmp_pair(const void *a, const void *b) {
//     const Pair *pa = (Pair *)a;
//     const Pair *pb = (Pair *)b;
//     if (pa->value < pb->value) return -1;
//     if (pa->value > pb->value) return 1;
//     return 0;
// }

// /**
//  * Note: The returned array must be malloced, assume caller calls free().
//  */
// int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
//     Pair *searchArr = malloc(numsSize * sizeof(Pair));
//     for (int i = 0; i < numsSize; i++) {
//         searchArr[i].value = nums[i];
//         searchArr[i].idx = i;
//     }

//     qsort(searchArr, numsSize, sizeof(Pair), cmp_pair);

//     int l = 0;
//     int r = numsSize - 1;
//     while (l < r) {
//         int sum = searchArr[l].value + searchArr[r].value;
//         if (sum == target) {
//             int *result = malloc(2 * sizeof(int));
//             result[0] = searchArr[l].idx;
//             result[1] = searchArr[r].idx;

//             free(searchArr);
//             *returnSize = 2;
//             return result;
//         }

//         if (sum > target) r--;
//         else l++;
//     }

//     free(searchArr);
//     *returnSize = 0;
//     return NULL;
// }