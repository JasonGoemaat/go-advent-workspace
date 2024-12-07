# Naming

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

