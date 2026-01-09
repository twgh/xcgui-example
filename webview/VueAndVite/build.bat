@echo off
@chcp 65001 >nul
REM 生产模式构建脚本

echo ===================================
echo Vue + Vite 桌面应用 - 生产构建
echo ===================================
echo.

REM 检查 VueAndVite.go 中的 isDebug 是否为 false
findstr /C:"isDebug = false" VueAndVite.go >nul
if errorlevel 1 (
    echo 警告: VueAndVite.go 中的 isDebug 仍为 true
    echo 请手动将 VueAndVite.go 中的 var isDebug = true 改为 var isDebug = false
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


echo.
echo [2/3] 正在构建前端资源 (生产模式)...
call npm run build
if errorlevel 1 (
    echo 构建失败！
    pause
    exit /b 1
)

echo.
echo [3/3] 正在编译 Go 程序 (单文件)...
go build -ldflags="-s -w -H windowsgui" -trimpath -o VueDesktopApp.exe VueAndVite.go
if errorlevel 1 (
    echo 编译失败！
    pause
    exit /b 1
)

echo.
echo ===================================
echo 构建完成！
echo Output: VueDesktopApp.exe
echo 说明: Vue 资源已嵌入到 exe 中，无需额外文件
echo ===================================
echo.
pause

