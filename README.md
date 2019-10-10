# syslog-server
server for golang package log/syslog and sink to file

## server

监听 UDP

## handler

写文件

## parser

基于 rfc3164 调用 github.com/jeromer/syslogparser，主要将时间格式设定为 log/syslog 中的 time.RFC3339