@echo off
@chcp 65001 >nul
REM 开发模式启动脚本 - 同时启动 Vite 开发服务器和 Go 应用

echo ===================================
echo Vue + Vite 桌面应用 - 开发模式
echo ===================================
echo.


echo.
REM 检查 VueAndVite.go 中的 isDebug 是否为 false
findstr /C:"isDebug = true" VueAndVite.go >nul
if errorlevel 1 (
    echo 警告: VueAndVite.go 中的 isDebug 仍为 false
    echo 请手动将 VueAndVite.go 中的 var isDebug = false 改为 var isDebug = true
    echo 然后重新运行此脚本
    pause
    exit /b 1
)


REM 检查 node_modules 是否存在
if not exist "node_modules" (
    echo [1/3] 首次运行，正在安装依赖...
    call npm install
    if errorlevel 1 (
        echo 安装依赖失败！
        pause
        exit /b 1
    )
) else (
    echo [1/3] 依赖已安装
)


setlocal
:: 检查 dist 目录内是否有文件，如果没有则创建 keep 文件
:: 设置目标目录
set "targetDir=dist"

:: 检查目录是否存在，如果不存在则创建
if not exist "%targetDir%" (
    mkdir "%targetDir%"
    echo 已创建目录: %targetDir%
)

:: 检查目录中是否有文件
set "hasFiles=false"
for /f "delims=" %%i in ('dir /a-d /b "%targetDir%" 2^>nul') do (
    set "hasFiles=true"
)

:: 如果没有文件，则创建 keep 文件
if "%hasFiles%"=="false" (
    type nul > "%targetDir%\keep"
    echo 已创建文件: %targetDir%\keep
)

endlocal


echo.
echo [2/3] 正在启动 Vite 开发服务器...
echo 提示: Vite 开发服务器将在新窗口中运行
echo 修改 Vue 代码后会自动热重载，无需手动刷新
echo.
start "Vite Dev Server" cmd /k "npm run dev"

REM 等待 2 秒，让 Vite 服务器启动
timeout /t 2 /nobreak >nul

echo [3/3] 正在启动 Go 应用...
echo.
echo ===================================
echo 应用已启动！
echo ===================================
echo.
echo 使用说明:
echo   - Vite 开发服务器: http://localhost:5173
echo   - 修改 Vue 代码会自动热重载
echo   - 关闭 Vite 窗口可停止开发服务器
echo   - 按 Ctrl+C 可停止 Go 应用
echo.
go run VueAndVite.go

echo.
echo 应用已关闭
pause
