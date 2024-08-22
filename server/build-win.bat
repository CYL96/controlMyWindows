@echo off

chcp 65001 > nul
echo 当前时间：%time:~0,8%

echo 开始编译 control-server.exe
go build -ldflags -H=windowsgui -o control-server.exe

if %errorlevel%==0 (
    echo 编译成功
) else (
    echo 编译失败
)

echo 开始编译 control-server-cli.exe
go build  -o control-server-cli.exe

if %errorlevel%==0 (
    echo 编译成功
) else (
    echo 编译失败
)

echo 当前时间：%time:~0,8%
