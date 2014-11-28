@echo off

set GOPATH=D:\work\00.LoadGen\loadgengo

go install loadgen/system
go install loadgen/http

go install loadgen

go install loadgen/loadgen
go install loadgen/http/httprecorder
go install loadgen/http/httpparser