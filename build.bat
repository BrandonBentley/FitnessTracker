@echo off
SETLOCAL

IF EXIST buildInfo (
	del buildInfo
)

IF EXIST Server.exe (
	del Server.exe
	echo Removing Previous Server
)
IF NOT EXIST Server.exe (
	go build -i -o Server.exe go/main.go go/handlers.go go/prompt.go go/structs.go go/helpers.go go/global.go >> buildInfo
) ELSE (
	echo Previous Server Removal Failed
)
IF EXIST Server.exe (
	echo Build Successful: Server.exe
) ELSE (
	SET "failed=true"
	type buildInfo
)

IF NOT DEFINED failed (
	del buildInfo
)
echo Build Complete