#simple logger

## Document

* http://godoc.org/github.com/tenntenn/log

## Usage

```
// Write to stderr as error log
log.Error("error")

// Write to stdout as info log
log.Info("info")

// Set log level to debug
log.SetLogLevel(LogLevelDebug)

// Write to stdout as debug log
log.Debug("debug")
```
