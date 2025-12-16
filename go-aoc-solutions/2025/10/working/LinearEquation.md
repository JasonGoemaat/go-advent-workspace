Sample

    [.###.#]
    
    a(0,1,2,3,4)
    b(0,3,4)
    c(0,1,2,4,5)
    d(1,2)
    
    {10,11,11,5,10,5}


    1a +    1b +    1c +    0d  = 10    (joltage 0)
    1a +    0b +    1c +    1d  = 11    (joltage 1)
    1a +    0b +    1c +    1d  = 11    (joltage 2)
    1a +    1b +    0c +    0d  = 5     (joltage 3)
    1a +    1b +    1c +    0d  = 10    (joltage 4)
    0a +    0b +    1c +    0d  = 5     (joltage 5)

Simplified:

    1a +    1b +    1c +    0d  = 10    (joltages 0,4)
    1a +    0b +    1c +    1d  = 11    (joltages 1,2)
    1a +    1b +    0c +    0d  = 5     (joltage 3)
    0a +    0b +    1c +    0d  = 5     (joltage 5)

[Row echelon form]()

    1a +    1b +    1c +    0d  = 10    (joltages 0,4)
    1a +    1b +    0c +    0d  = 5     (joltage 3)
    1a +    0b +    1c +    1d  = 11    (joltages 1,2)
    0a +    0b +    1c +    0d  = 5     (joltage 5)


Maximum times I can press the buttons:

* a = 5 (joltage 3)
* b = 5 (joltage 3)
* c = 5 (joltage 5)
* d = 11 (joltages 1,2)

LOOKING at it, only button c affects joltage 5, so it MUST
be the case that `c=5`, I have to press button c EXACTLY
5 times.  That lowers the required joltages for 0,4 from 10
to 5 if I get rid of the c column.  It also loweres the required
joltages for 1,2 from 11 to 6.   This alters the maximum for 
'd' now, changing it from 11 to 6

* a, b, c, d are how many times I press the button, that equates to x1..xn
* a11..amn is the matrix of 0s and 1s defined by whether a button(n) affects a voltage (m)
* b is the required voltages
* puzzle is to find a solution that MINIMIZES a+b+c+d or x1+x2+x3+x4

Wikipedia says a linear system can behave in one of three ways:

1. The system has infinitely many solutions.  
2. The system has a unique solution.  
3. The system has no solution  

But what about this case?

    a(0)
    b(1)
    c(0,1)
    {1,1}

In this case button a can affect joltage 0, button b can affect
joltage 1, and button c can affect both joltages.  There are 
two possible solutions.   You could press a,b or you could just press c.   There are no more solutions.

## Sample from reddit

Guy was having trouble with this, I was able to figure it out manually...

```
  1    2    3    4    5    6    7    8    9
 1.0  0.0  0.0  0.0  0.0  0.0  0.0  0.0  0.5 | 20.5
 0.0  1.0  0.0  0.0  0.0  0.0  0.0  0.0  0.0 |  5.0
 0.0  0.0  1.0  0.0  0.0  0.0  0.0  0.0 -1.0 |  2.0
 0.0  0.0  0.0  1.0  0.0  0.0  0.0  0.0 -0.5 | -7.5
 0.0  0.0  0.0  0.0  1.0  0.0  0.0  0.0  1.0 | 20.0
 0.0  0.0  0.0  0.0  0.0  1.0  0.0  0.0  0.5 |  7.5
 0.0  0.0  0.0  0.0  0.0  0.0  1.0  0.0  1.0 | 35.0
 0.0  0.0  0.0  0.0  0.0  0.0  0.0  1.0  0.0 | 19.0
                                         -15

 13    5   17    0    5    0   20   19   15
  1    2    3    4    5    6    7    8    9
 1.0  0.0  0.0  0.0  0.0  0.0  0.0  0.0  0.5 | 13.0
 0.0  1.0  0.0  0.0  0.0  0.0  0.0  0.0  0.0 |  5.0
 0.0  0.0  1.0  0.0  0.0  0.0  0.0  0.0 -1.0 | 17.0
 0.0  0.0  0.0  1.0  0.0  0.0  0.0  0.0 -0.5 |  0.0
 0.0  0.0  0.0  0.0  1.0  0.0  0.0  0.0  1.0 |  5.0
 0.0  0.0  0.0  0.0  0.0  1.0  0.0  0.0  0.5 |  0.0
 0.0  0.0  0.0  0.0  0.0  0.0  1.0  0.0  1.0 | 20.0
 0.0  0.0  0.0  0.0  0.0  0.0  0.0  1.0  0.0 | 19.0
                                         -15

 13, 5, 17, 0, 5, 0, 20, 19, 15
 ```
 