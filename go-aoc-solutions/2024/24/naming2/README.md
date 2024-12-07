# Naming2

TLDR; `hjf,kdh,kpp,sgj,vss,z14,z31,z35`

My goal is to start with bare operations, then combine them
and name them something representing what they actually
do along with setting the largest bit

Ok, I start with simple "ADD" and "BOTH" when an operation
works with initial values for one bit and either XOR which
gives the result bit (ADD) or AND which results in a CARRY.

Now I want to name new things.  cpb is the CARRY of bit 01.
I know this because it's an ADD of 01 (xor) and BOTH of 00.
cpb = CARRY.  an OR is a CARRY of the largest bit in it's
children

    z02 = cpb XOR wmr
        cpb = djn OR tnr
            djn = ADD(01) AND BOTH(00)
            tnr = BOTH(01)
        wmr = ADD(02)

Converts to:

    z02 = CARRY(01) XOR ADD(02)

Now z03:

    z03 = tdw XOR vrk
        tdw = fwt OR vtm // CARRY(02)
            fwt = BOTH(02)
            vtm = ADD(02) AND CARRY(01)
        vrk = ADD(03)

Turns into:

    z03 = CARRY(02) XOR ADD(03)

Now z04:

    z04 = vjr XOR vbb
        vjr = ADD(04)
        vbb = hsf OR bjw // CARRY03
            hsf = ADD(03) AND CARRY(02)
            bjw = BOTH(03)

Becomes:

    z04 = ADD(04) XOR CARRY(03)

Now z05:

    z05 = ctv XOR wht
        ctv = ADD(05)
        wht = khw OR djm
            khw = BOTH(04)
            djm = ADD(04) AND CARRY(03)

    z05 = ADD(05) XOR CARRY(04)

...

Now down to z10:

    z10 = bgd XOR msb
        bgd = ADD(10)
        msb = bmn OR hbj
            bmn = BOTH(09)
            hbj = CARRY(08) AND ADD(09)

So the pattern is we have XOR of two things:

    ADD(b)
    CARRY(b-1)

And CARRY(b) is OR of:
    BOTH(b)
    CARRY(b-1) AND ADD(b)

## OK

Looking good, z22 is not findable:

    z22 = BOTH(22) XOR CARRY(21) (kdh XOR gjn)

It should look like this:

    z22 = ADD(22) XOR CARRY(21)

So I need to swap ADD(22) and BOTH(22) (kdh)

    hjf = ADD(22)
    
That worked!

### z31 is bad

    z31 = CARRY(31)

It should look like this:

    z31 = ADD(31) XOR CARRY(31)

Now looking at z32, it has my answer:

    ghr = ANSWER(31) OR BOTH(31) (kpp OR pwg)

That makes sense, swap z31 and kpp

### z35 is borked:

    z35 = BOTH(35)

Again, looking at z36 I see the answer:

    rqf = ANSWER(35) OR CARRY(35) (sgj OR ptb)

So swap z35 and sgj

## Done!

Cool, all of them show 'ANSWER(#)' except for the last bit
which should just be the carry of the previous operation:

    z45 = CARRY(44)

Swaps:

    swap(operations, "z14", "vss")
    swap(operations, "hjf", "kdh")
    swap(operations, "z31", "kpp")
    swap(operations, "z35", "sgj")

    z14
    z31
    z35
    sgj
    vss
    hjf
    kdh
    kpp
