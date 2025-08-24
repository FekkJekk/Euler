test = 10
for lowest in range(2,test):
    if test % lowest == 0:
        break
print(lowest)
print('test / lowest factor:',test/lowest)
for possible_highest in range(int(test/lowest),1,-1):
    for test in range(possible_highest,1,-1):
        print(test)
        prime = 0
        if possible_highest != test and possible_highest % test == 0:
            prime = 0
            print('tested')
            print('highest prime is',possible_highest)
            break
        if prime == 0:
            print('test2')
            break
