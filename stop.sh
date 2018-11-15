pid=`ps -ef|grep class|grep -v grep|awk '{print $2}'`
if [ -z "$pid" ];then
    exit
fi

echo kill config_server pid: $pid
kill -9 $pid
