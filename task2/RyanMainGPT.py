def euler2(limit=4000000):
    a, b = 1, 2
    total = 0

    while a <= limit:
        if a % 2 == 0:
            total += a
        a, b = b, a + b

    return total

print(euler2())
