@ECHO OFF

echo "Cleaning .exe files"
del main_equation_reader.exe
del main_equation_arranger.exe
del main_equation_disassembler.exe
del main_equation_solver.exe
del main_equation_reporter.exe

echo "Testing equation message"
go test equationmessage\tests\
if %errorlevel% neq 0 exit /b %errorlevel%

echo "Testing equation scanner"
go test equationscanner\tests\
if %errorlevel% neq 0 exit /b %errorlevel%

echo "Testing equation arranger"
go test equationarranger\tests\
if %errorlevel% neq 0 exit /b %errorlevel%

echo "Testing equation disassembler"
go test equationdisassembler\tests\
if %errorlevel% neq 0 exit /b %errorlevel%

echo "Testing equation solver"
go test equationsolver\tests\
if %errorlevel% neq 0 exit /b %errorlevel%

echo "Testing equation reporter"
go test equationreporter\tests\
if %errorlevel% neq 0 exit /b %errorlevel%

echo "Building equation reader"
go build main_equation_reader.go
if %errorlevel% neq 0 exit /b %errorlevel%

echo "Building equation arranger"
go build main_equation_arranger.go
if %errorlevel% neq 0 exit /b %errorlevel%

echo "Building equation disassembler"
go build main_equation_disassembler.go
if %errorlevel% neq 0 exit /b %errorlevel%

echo "Building equation solver"
go build main_equation_solver.go
if %errorlevel% neq 0 exit /b %errorlevel%

echo "Building equation reporter"
go build main_equation_reporter.go
if %errorlevel% neq 0 exit /b %errorlevel%

echo "Done testing and building successfully"