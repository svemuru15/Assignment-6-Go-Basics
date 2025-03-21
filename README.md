# Assignment-6-Go-Basics

Do you find your code showing race conditions? Try it with larger p values to get many concurrent processes. What can you do about that? Why or why not are you seeing race conditions?

Yes, my code shows race conditions because the second program has a goroutine that updates the variable in a proper way using synchronization. Without synchronization, we would not know which goroutine is being processed and what number routine res is storing the variable from. In order to avoid this we use mutex because this makes sure that only one goroutine can modify res at a time. When I try it with larger p values the code works and I am not seeing race conditions because of the mutex and synchronizing the wait groups. If this was not in place we would be getting different answers everytime depending on when the goroutines finish and there would be no order. 
