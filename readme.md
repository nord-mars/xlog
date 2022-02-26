# Xlog - golang library with additional log features


## Install

```
go get -u github.com/nord-mars/xlog
```

## Documentation
* [logger](https://github.com/nord-mars/xlog/tree/main/examples/logger)
* [logger.short](https://github.com/nord-mars/xlog/tree/main/examples/logger.short)
* [logger.speed](https://github.com/nord-mars/xlog/tree/main/examples/logger.speed)

## Description
* Summary:
  - Log level
  - Log category
  - Log fields
  - Log filename

### Log level
debug_level (global variable): write or skip record to log

### Log category
* Log record category:
  - INFO
  - WARN
  - ERROR
  - FATAL - add call stuck to log

### Log fields
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

### Log filename
* Field descriptions:
  - FILE_PID  - add filename.PID.log (split logs for same time application)
  - FILE_DATE - add filename.DATE.log
  - FILE_TIME - add filename.TIME.log

## License
Use of this source code is governed by a BSD-style
