To make golang application compile
1.  Set the GOLANG environment be your desired path.
    Your go files should be under $GOLANG/src/<golang files and packages>

2.  go get "github.com/jjeffery/stomp"
3.  go get "github.com/google/uuid"
4.  go get "github.com/gomodule/redigo/redis"

To make golang application be built via command line:
go build <main_go_file_name>.
It will create exe file named <main_go_file_name.exe>
It is runnable by ./main_go_file_name.exe

Use test_and_build.bat to test and build all services.