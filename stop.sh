pid=`ps -ef|grep goadmin|grep -v grep|awk '{print $2}'`
if [ -z "$pid" ];then
    exit
fi

echo kill goadmin pid: $pid
kill -9 $pid
