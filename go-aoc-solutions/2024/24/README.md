# 2024 Day 24

Christmas eve!  Doing it on Christmas though.  I wonder what part 2  will
be.  I thought it might be coming up with a different number to generate
certin input bits, but the bits as they are might produce only one output.

Oh, maybe there is a trick to it where the output can be made by other inputs.
For example if x00 and  y00 are 0,1 and only used as an input for an OR
gate, the actual possibilities are 0,1 + 1,0 + 1,1, just not 0,0.

OR!  Maybe it will give you a different answer and ask you what input bits
to change to make the new number.

OR!  Maybe give you a desired output and you have to find what operations
to change, like maybe an OR should be an XOR?

## Doing it

I thought maybe of tracking the bits for each level.   But you can combine
bits from different positions.  It may or bit 16 and bit 2 for example and
the result may go to bit 11.

## Part 2

Ah, the circuit SHOULD be doing addition, but there are four gates that need
to have their outputs swapped to make it work correctly.

Now before I go doing a bunch of programming, let's look at the actual size.

Inputs: 90 (45 x and 45 y)

Operations: 222 (46 z + 176)
There should be 45 operations with each x and y connected, that's not that many
more.  There should be an z00 = x00 XOR y00 for the first bit.
For the second bit we need the carry from the previous (x00 AND y00) included also.
I guess I don't know how to make that adding circuit.  But there shouldn't be too many
levels and there shouldn't be much cross-over.   You just need to carry the previous
operation to the next bit.

I'll write something to start with the zs and indent each level.

Ok, the output is large because each bit comparison is shifted to make the final
result.   `z00` is pretty easy, it's just XOR as expected:

    z00 = y00 XOR x00

`z01` has to take into account if both those bits are set with AND, then it does
the xor of wqc which is the xor of bit 1, and the AND of bit 0:

    z01 = wqc XOR qtf
        wqc = x01 XOR y01
        qtf = y00 AND x00

`z02` builds on the chain.   wmr is the xor of bit 2 in the inputs, then 
cpb is djn OR tnr which is the AND of bit 1 of othe inputs combined with the xor of the previous
bit.

    z02 = cpb XOR wmr // carry from bit 1 xor natural set of bit 2
        cpb = djn OR tnr // carry from bit 1
            djn = wqc AND qtf // carry bit 0 and bit 1 set naturally = carry bit 1
                wqc = x01 XOR y01 // bit1 would be set without carry
                qtf = y00 AND x00 // carry bit0
            tnr = y01 AND x01
        wmr = x02 XOR y02 // natural bit 2

So looking at the pattern it chains down like that and the nodes are re-used.
So I wrote something to go through each of the 45 input bits and try each
of the combinations and compare to actual.  Here's what I see:

    Bit 13: 1 + 1 = 00 got 32768 expected 16384
    Bit 14: 1 + 0 = 10 got 32768 expected 16384
    Bit 14: 0 + 1 = 10 got 32768 expected 16384
    Bit 14: 1 + 1 = 01 got 16384 expected 32768
    Bit 22: 1 + 0 = 10 got 8388608 expected 4194304
    Bit 22: 0 + 1 = 10 got 8388608 expected 4194304
    Bit 22: 1 + 1 = 01 got 4194304 expected 8388608
    Bit 30: 1 + 1 = 00 got 4294967296 expected 2147483648
    Bit 31: 1 + 0 = 10 got 4294967296 expected 2147483648
    Bit 31: 0 + 1 = 10 got 4294967296 expected 2147483648
    Bit 34: 1 + 1 = 00 got 68719476736 expected 34359738368
    Bit 35: 1 + 0 = 10 got 68719476736 expected 34359738368
    Bit 35: 0 + 1 = 10 got 68719476736 expected 34359738368
    Bit 35: 1 + 1 = 01 got 34359738368 expected 68719476736

So it looks like everything up to z13 is correct.  However
z14 gets set by carry of x14+y14, and z15 get set by carry of x13+x14.
So look for the carries for z14 and z15 and swap them?

z00 is too simple, just xor inputs

z01 is too simple for the pattern too, xor of bits01 and OR

z02 is where we get the pattern for the outputs starting.  The output 'z' is
an xor of A) the same bit XOR which is basically saying 1+0 or 0+1, and
the carry (0+0 or 1+1 is 0 and carry is set)

So output is XOR of XOR and OR.   The OR being the carry, which is the OR
of two AND.   And A) is xor of previous (setting to 1) and carry of prior.

I think z14 should be swapped with vss:

    vss = nhg XOR ttd
    z14 = tfw OR cgt

Cool, that solved one problem:

    Bit 22: 1 + 0 = 10 got 8388608 expected 4194304
    Bit 22: 0 + 1 = 10 got 8388608 expected 4194304
    Bit 22: 1 + 1 = 01 got 4194304 expected 8388608
    Bit 30: 1 + 1 = 00 got 4294967296 expected 2147483648
    Bit 31: 1 + 0 = 10 got 4294967296 expected 2147483648
    Bit 31: 0 + 1 = 10 got 4294967296 expected 2147483648
    Bit 34: 1 + 1 = 00 got 68719476736 expected 34359738368
    Bit 35: 1 + 0 = 10 got 68719476736 expected 34359738368
    Bit 35: 0 + 1 = 10 got 68719476736 expected 34359738368
    Bit 35: 1 + 1 = 01 got 34359738368 expected 68719476736
    297ms part2("input.aoc") = 0 (GOOD)

So output should be XOR of XOR and OR, let's check bits 22 and 23 which seem to be off.

    z22 = kdh XOR gjn
        kdh = y22 AND x22
        gjn = hbc OR wdr

    z23 = tks XOR sbg
        tks = x23 XOR y23
        sbg = bjb OR hjf

So need to look in z23 for an XOR where one component is x22 XOR y22

I think z23 is actually zupposed to be z22, or close...

Ok, re-ran to remove 'known good':

    z22 = kdh XOR gjn
    z30 = rbw XOR dfv
    z31 = nrr AND sms
    z34 = qfs XOR rfv
    z35 = y35 AND x35
    z45 = nmb OR mbk
        nmb = fsf AND fgs
        mbk = x44 AND y44

So z22 is bad for sure, but all it's children are good, so what gives?

Oh, z35 is missing `hjf = y22 XOR x22`, as is z22.
`hjf OR XXX` should be XOR into something to make z22.  Maybe kdh and sbg should be swapped?

    sbg = bjb OR hjf
