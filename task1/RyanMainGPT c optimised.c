#include <stdio.h>

int sum_of_multiples(int n, int limit) {
    int k = limit / n;
    return n * k * (k + 1) / 2;
}

int main() {
    int limit = 999;

    int sum3 = sum_of_multiples(3, limit);
    int sum5 = sum_of_multiples(5, limit);
    int sum15 = sum_of_multiples(15, limit);

    int total = sum3 + sum5 - sum15;

    printf("%d\n", total);
    return 0;
}
