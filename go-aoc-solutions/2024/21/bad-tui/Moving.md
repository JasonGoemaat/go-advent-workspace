# Moving

Here's the actions I see happening for the sample '0249'


`StartLine`
    - clears and resets all keypads
    - sets `[0].queue = "0249"`
    - returns `ProcessQueue(0)`
`ProcessQueue(0)`
    - calculate moves to get from current pos 'A' (2,3) to '0' (1,3) AND press it (`<A`)
    - returns `Enqueue(1, "<A")`
    `Enqueue(1, "<A")`
        - sets `[1].queue = "<A"`
        - returns `ProcessQueue(1)`
    `ProcessQueue(1)`
        - calculates moves to get from current pos 'A' (2,0) to '<' (0,1) AND press it (`v<<A`)
        - returns `Enqueue(2, "v<<A")`
        `Enqueue(2, "v<<A")`
            - sets `[2].queue = "v<<A"`
            - returns `ProcessQueue(2)`
        `ProcessQueue(2)`
            - calculates moves to get from current pos 'A' (2,0) to 'v' of `v<<A` (1,1) AND press it (`v<A`)
            - returns `Enqueue(3, "v<A")`
            `Enqueue(3, "v<A")`
                - sets `[3].queue = "v<A"`
                - returns `ProcessQueue(3)`
            `ProcessQueue(3)`
                - // HERE is where the work gets done, special processing as it is OUR controller
                - handle `v` in `v<A`
                - set Keypressed[3] = 'v'
                - move pos[2] from (2,0) to (2,1)
                - QueueIndex[3]++
                - return `ProcessQueue(3)`
            `ProcessQueue(3)`
                - handle `<` in `v<A`
                - set Keypressed[3] = '<'
                - move pos[2] from (2,1) to (1,1)
                - QueueIndex[3]++
                - return `ProcessQueue(3)`
            `ProcessQueue(3)`
                - handle `A` in `v<A`
                - set Keypressed[3] = '<'
                - move pos[2] from (2,1) to (1,1)
                - QueueIndex[3]++
                - return `ProcessQueue(3)`
            `ProcessQueue(3)`
                - we are past end of queue
                - move queue to history and clear queue
                - return `ContinueQueue(2)`
        `ContinueQueue(2)`
            - `QueueIndex[2]++`
            - if after end, return `EndQueue(2)`
            - otherwise, return `ProcessQueue(2)`
        `ProcessQueue(2)`
            - calculates moves to get from current pos 'v' (1,1) to '<' of `v<<A` (0,1) AND press it (`<A`)
            - returns `Enqueue(3, "<A")`
            `ProcessQueue(3)` ----- handle like above
        `ProcessQueue(2)`
            - calculates moves to get from current pos '<' (0,1) to '<' of `v<<A` (0,1) AND press it (`A`)
            - returns `Enqueue(3, "A")`
            `ProcessQueue(3)` ----- handle like above
        `ProcessQueue(2)`
        