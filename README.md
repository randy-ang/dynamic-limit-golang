# dynamic-limit-golang
uses channel and a timer goroutine to set a dynamic limit for a process (using timeout rather than hard limit)

We execute 2 goroutine functions:
1. a timer goroutine function that will be stop the whole process once the specified amount of time has passed.
2. a data processing goroutine function that will send the processed data to the channel. Once sent, 2 things will happen:
    - the data will be stored once the channel received the message
    - the data processing module will process another data (repeat no. 2)

By implementing this, we don't need to set a hard limit when trying to avoid timeout. This way, when saving to database, if the database is slightly slower than usual, the timeout will not be triggered because no hard limit is set.

go run example.go:
- case where saving data has no delays

go run example_2.go:
- case where saving data has a delay

go run example_3.go:
- example_2 but with maximum number of data to be processed rather than timeout
