# 2015 Day 19 Part 2 is BS

So you really HAVE to examine the input to figure this one out.   You might
be able to check your input and find things like this, but there will always
be inputs that could screw you over.  So there seem to be certain things
that are special.   For example, these replacements end in 'Ar' and there is
no other way to get an 'Ar'.   So I think we can split the main string
by 'Ar' and reduce those.

    Al => ThRnFAr
    B => TiRnFAr
    Ca => PRnFAr
    Ca => SiRnFYFAr
    Ca => SiRnMgAr
    H => CRnAlAr
    H => CRnFYFYFAr
    H => CRnFYMgAr
    H => CRnMgYFAr
    H => NRnFYFAr
    H => NRnMgAr
    H => ORnFAr
    N => CRnFAr
    O => CRnFYFAr
    O => CRnMgAr
    O => NRnFAr
    P => SiRnFAr

So we are going backwards here, I don't think we need to worry about for instance
`Al => ThRnFAr`, that is one whole string.   There won't be a way to expand the
`F` for instance.

    Al => ThF
    B => BCa
    Ca => CaCa
    B => TiB
    Ca => CaCa
    Ca => PB
    Ca => SiTh
    F => CaF
    F => PMg
    F => SiAl
    H => HCa
    H => NTh
    H => OB
    Mg => BF
    Mg => TiMg
    N => HSi
    O => HP
    O => OTi
    P => CaP
    P => PTi
    Si => CaSi
    Th => ThCa
    Ti => BP
    Ti => TiTi
    e => HF
    e => NAl
    e => OMg

Ok, another thing.    The double-letters are making it more difficult.   Let's
try to condense the multiple letter elements into a single letter.  `Al`, `Mg`,
`Si`, and `Ca` are the only elements starting with those letters.  There are
two double-letter elements starting with `T` (`Th` and `Ti`), so let's use
`h` and `i` for those.  Then we have some elements appearing only in the results
and that cannot be expanded (`Rn` and `Ar`), so we can use `n` and `r` for
those, or maybe I should use special characters like `.` and `-`?  Oops, I see
there's another `C` that only appears in the output like `CRn` but there are
other `Rn`, so I'll make `Ca` into `a`

    Al -> A
    Ar -> r
    Ca -> a
    Mg -> M
    Rn -> n
    Si -> S
    Th -> h
    Ti -> i

Let's use digits to avoid any ambiguity:

    Al -> 0
    Ar -> 1
    Ca -> 2
    Mg -> 3
    Rn -> 4
    Si -> 5
    Th -> 6
    Ti -> 7

Ah, I notice now there's a `Y` that is not on any inputs and is only on a some
distinctive outputs, might be worth a look.   But now let's look at the end of
the string, noting the appearances of `1` (`Ar`), where once you reach that, 
there is no way to modify the right side:

    FYP312PB2PB54FYPB2F1250
        ^              ^

So we have an ending of `250` and then need to use one of the values ending in
'1', which we couldn't use before or we would never get the 250.  Here are
things we can use starting at 'e' to make the `1250`:

    e => HF
    e => N0
    e => O3
    5 => 25
    F => 50

It seems like we could start with e, HF, H50, H250.   Wait, there's too many options...

So looking at it more, The only ways to get `Ar` (`1`) and `Rn` (`4`) are to have
them placed together in a string, and they cannot then deviate.  And the `4` is
always first, and the string always ends in a `1`.  So we can divide that way somehow.

    0 => 6 4 F 1
    B => 7 4 F 1
    2 => P 4 F 1
    2 => 5 4 F Y F 1
    2 => 5 4 3 1
    H => C 4 0 1
    H => C 4 F Y F Y F 1
    H => C 4 F Y 3 1
    H => C 4 3 Y F 1
    H => N 4 F Y F 1
    H => N 4 3 1
    H => O 4 F 1
    N => C 4 F 1
    O => C 4 F Y F 1
    O => C 4 3 1
    O => N 4 F 1
    P => 5 4 F 1

Ah, NEAT!  The only way to get a `Y` is to have it between those 4s and 1s, and
so if there are two `Y`s between those, it has to come from that one expansion of
`H`.   And for a single `Y` there are only 5 options.   There are many more options
for zero `Y`.  Let's take our final molecule and use these rules:

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

Ok, check out [Try1.md](Try1.md) where I try to parse manually... The main
thing I think I'll have problems with is repeated 4..1 at the same level.
Let me create a helper page...
