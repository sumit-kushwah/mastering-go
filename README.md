## Mastering Go Concurrency

### Current Understanding
**Channels**

These are of two types buffered and unbuffered, buffered channel have size and unbuffered has no size. They are mainly used to communicate between go routines since go routine function can't return any value.

**WaitGroup**

Helps in maintaining the go routines.
