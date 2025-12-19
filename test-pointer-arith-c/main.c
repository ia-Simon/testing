#include <stdio.h>

typedef struct {
    int prop1;
    int prop2;
} Test;

int main() {
    Test arr[] = {{1, 2}, {3, 4}};

    printf("%d\n", (*(arr + 1)).prop1);
}