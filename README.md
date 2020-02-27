# setinterval
A golang package that helps you run a function intervalically.

It behaves in a similar way to Javascript's `setInterval` function.

It has two exported functions *Sync()* and *Async()*

Both accept a `time.Duration`, and a function to be called repeatedly. *Sync()* blocks, and keeps calling the function periodically, while *Async()* returns a pointer to an **Interval** which has methods *Pause()* and *Resume()*, and does not block.
