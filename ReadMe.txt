Explaination:

A future is a pattern where a computation is started and its result is returned as soon as it's ready. 
The Future interface is declared with methods to get the result, check if it's complete or cancelled, and to cancel the computation.

The implementation of the Future interface is provided by the futureImpl struct. 
It has fields to store the result, status of completion and cancellation, context, and a function to cancel the context.

The Run function takes a function fn as input and returns a Future. 
It creates a channel to store the result of the computation, and starts a goroutine that runs fn. 
The goroutine waits for either the context to be done or for fn to finish and then sends the result to the channel.

In main, a future is created by calling Run with a function that sleeps for some seconds and returns the string "Hello World!". 
The result of the future is obtained using the GetWithTimeout method with a timeout of 3 seconds. The result is printed if there's no error, otherwise, the error is printed.