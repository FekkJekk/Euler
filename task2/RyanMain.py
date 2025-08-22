NextNum = 0
CurrentNum = 2
PrevNum = 1
Sum = 0 #sum of even numbers found
while CurrentNum < 4000000:
    if CurrentNum % 2 == 0:
        Sum += CurrentNum
    NextNum = CurrentNum + PrevNum
    PrevNum = CurrentNum
    CurrentNum = NextNum
print(Sum)
