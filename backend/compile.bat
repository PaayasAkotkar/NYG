@echo off

for /d %%i in (*) do (
    echo Folder: %%i
    cd %%i
    go build
    if errorlevel 1 (
        echo Build failed in %%i
    ) else    if not "%%i"=="SQL-Manager" (
        echo Running app in %%i
        start "" "app.exe"
    )
    REM imp: come back to parent directory
    cd ..
)
pause