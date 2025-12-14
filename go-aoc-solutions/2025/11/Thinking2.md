# Method 1:

Making parent go below children, moving after the last one
of it's children in the list and moving everything else
up to where it is.  Not great because there are lots of
elements moving for every step

    ggg: out
    hhh: out
    fff: ggg hhh
    hub: fff
    ddd: hub
    dac: fff
    eee: dac
    ccc: ddd eee
    tty: ccc
    bbb: tty
    fft: ccc
    aaa: fft
    svr: aaa bbb

# Method 2:

Here I can build the list, adding any device that doesn't have
parents to the new list and removing them from the old and removing
them from the parents list of any children.

    
    svr: aaa bbb
    aaa: fft
    bbb: tty
    fft: ccc
    tty: ccc
    ccc: ddd eee
    ddd: hub
    eee: dac
    hub: fff
    dac: fff
    fff: ggg hhh
    ggg: out
    hhh: out
    out:
