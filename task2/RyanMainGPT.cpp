#include <iostream>

int main() {
    // Initialize variables
    int NextNum = 0;
    int CurrentNum = 2;
    int PrevNum = 1;
    int Sum = 0;  // Sum of even Fibonacci numbers found

    // Loop until CurrentNum exceeds 4,000,000
    while (CurrentNum < 4000000) {
        // If the current number is even, add it to the sum
        if (CurrentNum % 2 == 0) {
            Sum += CurrentNum;
        }
        
        // Update the Fibonacci sequence
        NextNum = CurrentNum + PrevNum;
        PrevNum = CurrentNum;
        CurrentNum = NextNum;
    }

    // Output the result
    std::cout << "Sum of even Fibonacci numbers: " << Sum << std::endl;

    return 0;
}
