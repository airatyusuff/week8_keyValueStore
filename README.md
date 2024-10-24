## Key-Value Store Using Go Concurrency

> Actor Model: Don't communicate by sharing, Share by communicating

- Worked with using channels as a 'central data source'
> channels: the communication pipelines that hold the data to be shared

### Refactoring/Fixing Code Smells
- Switch statements can almost always be refactored (Almost?)

- In this exercise, it was refactored using Command pattern
- - i.e bunch of requests to be processed/executed but executed in different ways
- - commonalities: executing a Request; each uses similar params (operation)
