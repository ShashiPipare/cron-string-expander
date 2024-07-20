# cron-string-expander
Receives a cron string in standard format with command line and explains each field of cron, when the command will be executed.


##
Pass the cron string in the given format:

go run main.go minutes hours days-of-month months days-of-week command

Example:
go run main.go */15 0 1,15 * 1-5 /usr/bin/find
go run main.go */20 0,6,12 1,15 * 1-5 http://localhost:4200/api/v1/users/add