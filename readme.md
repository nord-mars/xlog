# Xlog - extend golang log

## Install

```
go get -u github.com/nord-mars/xlog
```

## Documentation
* [logger](https://github.com/nord-mars/xlog/tree/main/examples/logger) - example full log setup
* [logger.short](https://github.com/nord-mars/xlog/tree/main/examples/logger.short) - example short log setup
* [logger.speed](https://github.com/nord-mars/xlog/tree/main/examples/logger.speed) - example log speed

## Description
* Log level    - write log or not
* Log category - Constant: log field status
* Log fields   - FLAG: setup additional log fields
* Log filename - FLAG: setup log file name

### Log level
debug_level (global variable): write or skip record to log

### Constants describing the category of the log entry:
* INFO - the information record
* WARN - the warning record
* ERROR - the error record
* FATAL - the fatal record and call stuck

### FLAG for Log fields
* Standard:
  - Ldate         - DATE field. Example: 2020/12/15
  - Ltime         - TIME field. Example: 06:19:41
  - Lmicroseconds - MILLISECOND field. Example: .800297
  - Lmsgprefix    - [PREFIX-PLACE] - before / after

* Custom:
  - LINE_CALL - prefix: add [__FILE__:__LINE__] - debug
  - LINE_PID  - prefix: add [PID]
  - LINE_HOST - prefix: add [HOSTMAME]
  - LINE_APP  - prefix: add [APPNAME]

### FLAG for Log filename
* Field descriptions:
  - FILE_PID  - add filename.PID.log (split logs for same time application)
  - FILE_DATE - add filename.DATE.log
  - FILE_TIME - add filename.TIME.log

## License
Use of this source code is governed by a BSD-style
