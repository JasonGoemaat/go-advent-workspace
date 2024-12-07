# Misc

Things I was thinking through, taking out to clean up code:

```go
// Here's the tricky part...  Easy to write in code, but several things
// can be done.  What actions will comprise a 'Step'?
// 1. starting out, will have empty for everything but Queues[0] will add "029A" and set IsProcessingLine = true
// 2. state will show CurrentPad 0 has work to do (QueueIndex[0] is 0, but Queues[0] has len 4)
//		CurrentPad 0 will calculate moves required to get to 0 and add to next Pad and change CurrentPad to it
// 3. do #2 for CurrentPad 1, which has queue "^^^<<A" and will submit "<A" to Pad 2
// 4. do #2 for CurrentPad 2, which has queue "<A" and will submit "v<<A" to Pad 3
// 5. do #2 for CurrentPad 3, but since we are the final pad, we don't submit
//		CurrentLine: 0, IsProcessingLine: true
//		CurrentPad: 3
//		Queue: "v<<A", QueueIndex 0
//			Set KeyPressed to 'v', QueueIndex = 1
//			Set KeyPressed to '<', QueueIndex = 2
//			Set KeyPressed to '<', QueueIndex = 3
//			Set KeyPressed to 'A', QueueIndex = 4
//			QueueIndex = 4, so need to move up and handle Pad 3.  At this point Pad 3
//				as Queue '<A' and QueueIndex 0.  I set KeyPressed[2] to '<' and QueueIndex[2] to 1
// and set Queue[3] to say `A`, QueueIndex[3] to 0, and CurrentPad to 3 again
//		here is where I press a key.

// AT THE END OF A LINE
// when CurrentPad == 0 and I move
```
