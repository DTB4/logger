# logger
DTB4's educationlal logger project

## !Requirements:
<code>/logs/</code>  directory must be in your project

## Features:
- 4 levels of loging: Info, Warning, Error, Fatal.
- FatalLog and FFatalLog finish application with <code>os.Exit(1)</code>
- Error and Fatal provide info about caller function in format  <code>path/file.go:lineNumber</code>
- Files of logger created every day of app working witch allow to quick search and also simplifies deletind of old logs
- Super Fast and intuitive func calls logger.Info for console output, logger.FInfo for file output
- Minimalistic parameters: <code>logger.InfoLog(message string, data interface{})</code> just call it like your standart logger with message
- Pretty colors in console output (thx :) Igor!)