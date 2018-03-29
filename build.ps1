Param([string]$command = "build")
$goPath= "go/"
$goFiles = "main.go", "handlers.go", "prompt.go", "structs.go", "helpers.go", "global.go"
$goString = ""
foreach($s in $gofiles) {
    $goString += ($goPath + $s + " ")
}
If($command -eq "build" -OR $command -eq "run") {
    If(Test-Path "./Server.exe") {
        Remove-Item -Path "./Server.exe"
        Write-Output "Removing Previous Server"
    }
    If(-Not (Test-Path "./Server.exe")) {
        go build -i -o Server.exe $goString >> buildInfo
    } Else {
        Write-Output "Previous Server Removal Failed"
        Exit
    }
}
If($command -eq "run") {
    ./Server.exe
} ElseIf ($command -eq "clean") {
    Remove-Item -Path "./Server.exe"
} ElseIf ($command -ne "build") {
    Write-Output "invalid command: " + $command
}
