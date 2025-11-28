@echo off
@REM do not get scared if many tabs open `its not a bug its a feature` 🤗

for /d %%i in (*) do (
if "%%i" == "backend" (
    cd %%i
for /d %%j in (*) do (
cd %%j
if not "%%j"=="SQL-Manager" (
    
    if not "%%j"=="protos" (
      
      echo Running app in %%j
 start "" "app.exe"

    )
)

    echo sub-folder: %%j
    
cd ..
)
    cd .. 
              ) else  (
                  echo folder: %%i
                  cd %%i && cd nyg-app && start "" ng s
                  cd .. 
)

)
pause
exit