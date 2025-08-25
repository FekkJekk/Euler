n = 999 #starting num


def test_brute_force():
    global n
    formulae = [n**2,n**2-n,n**2-2*n+1,n**2-2*n,n**2-3*n+2,n**2-3*n,n**2-4*n+4,n**2-4*n+3,
                n**2-4*n,n**2-5*n+6,n**2-5*n+4,n**2-5*n,n**2-6*n+9,n**2-6*n+8,n**2-6*n+5,
                n**2-6*n,n**2-7*n+12,n**2-7*n+10,n**2-7*n+6,n**2-7*n,n**2-8*n+16,n**2-8*n+15,
                n**2-8*n+12,n**2-8*n+7,n**2-8*n
    ]

    for output in formulae:
        testx = n**2
        if output > testx:
            print('doesn\'t work here:',output)
        else:
            print(output)
        testx = output
        

#left off ------------------------------------------------------------------
#was gonna make test functions:
#print_pd which would eventually print the palindrome
#for reading if output was a palindrome
#for loops nested, would change the formula from n**2 to n**2+(x*n)+y, growing but always capping
#   after ylimit so the program would go from 99*99 to 98*99 instead of 98*100 
#   (n**2 is 99*99, n**2-1 is 98*100)

















#print('n^2:',n**2)
#print('n^2:',n**2-n)
#print('n^2:',n**2-2*n+1)
#print('n^2:',n**2-2*n)
#print('n^2:',n**2-3*n+2)
#print('n^2:',n**2-3*n)
#print('n^2:',n**2-4*n+4)
#print('n^2:',n**2-4*n+3)
#print('n^2:',n**2-4*n)
#print('n^2:',n**2-5*n+6)
#print('n^2:',n**2-5*n+4)
#print('n^2:',n**2-5*n)
#print('n^2:',n**2-6*n+9)
#print('n^2:',n**2-6*n+8)
#print('n^2:',n**2-6*n+5)
#print('n^2:',n**2-6*n)
#print('n^2:',n**2-7*n+12)
#print('n^2:',n**2-7*n+10)
#print('n^2:',n**2-7*n+6)
#print('n^2:',n**2-7*n)
#print('n^2:',n**2-8*n+16)
#print('n^2:',n**2-8*n+15)
#print('n^2:',n**2-8*n+12)
#print('n^2:',n**2-8*n+7)
#print('n^2:',n**2-8*n)






