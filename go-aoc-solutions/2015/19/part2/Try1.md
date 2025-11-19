    C
    4
        5
        4
            2P73 Y 2P7
            4
                F
            1
            56F
        1
        25656PB225
        4
            5
            4
                773
            1
            PB2P3 Y P7
            4
                F
            1
            F
        1
        25
        4
            BP3
        1P
        4
        2P7
        4
            F
        1
        25622F
    1
    PB22P77
    4
        F
    1
    25
    4
        50 Y 56
        4
            F
        1
    1
    25
    4
        BF
    1
    225
    4
        56222F Y 2P7B256256P3
    1
    5
    4
        2PBF Y 22F
    1
    22225625
    4
        P
            4
                F
            1
            PB56P
            4
                F
            1
            5
            4
                3
            1
            2F Y F
        1
        25
        4
            50
        1
        7777777
        4
            P3
        1
        P777B5
        4
            50
        1
        77
        4
            P3
        1
        2F Y BPBP7
        4
            5
            4
                3
            1
            562F
        1
        256F
    1
    P
    4
        F
    1
    25
    4
        7B565
        4
            50 Y 2F
        1
        P
        4
            F
        1
        562F
    1
    22562225
    4
        P
        4
            2F
        1
        F Y P3
    1
    2PB2PB5
    4
        F Y PB2F
    1
    250

So I have indented it out so I can refer to line numbers.  Starting at an easy
one, line 119 is `2F` inside a 4-1.  So I have to start with the ones that
only have either a `2F` or something that will expand to `2F` in one step
inside.  Every single element expands into at least 2 elements, it never morphs
into a different single element.

Looking at the list of things with 1-2 characters between 4 and 1:

    H => C401

    2 => 5431
    H => N431
    O => C431

    0 => 64F1
    B => 74F1
    2 => P4F1
    H => O4F1
    N => C4F1
    O => N4F1
    P => 54F1

`0` and `3` don't expand into `2F`, but `F` does, so one of the ones with
`4F1` must be it. 

Hmmm...   Then nothing after the 1, so anything after the 1 must be expanded
prior to that expansion, right?

Well, there are just 11 top-level indents I believe, maybe that's possible
to semi-brute force?

Hmm...   I see `5` is special also.   Once I hit a `5`, it is stuck, no way
to get rid of it.   I can expand with as many `2`s as I want in front of it
though...  Likewise with a `6` there is no way to get rid of it, but you
can insert as many `2`s AFTER it as you want.