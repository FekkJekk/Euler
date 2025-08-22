z = 0
x = 2
y = 1
a = 0
while x < 4000000:
    if x % 2 == 0:
        a += x
    z = x + y
    y = x
    x = z
print(a)
