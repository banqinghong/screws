#!/bin/bash
# banqinghong@126.com
# 定时压缩日志，清理老的日志

pid_file=/var/run/nginx.pid
log_dir=/data/web_log/altair.zuul/
log_file=access.log
date=`date -d -1days +%F`

function cut_log () {
    if [ ! -d ${log_dir} ];then
        echo "`date '+%Y-%m-%d %H:%M:%S'` --> 日志目录不存在"
        exit 100
    fi
    cd ${log_dir}
    echo "`date '+%Y-%m-%d %H:%M:%S'` --> 开始切割日志"
    mv ${log_file} ${log_file}-${date}
    echo "`date '+%Y-%m-%d %H:%M:%S'` --> 日志切割完成"
    kill -USR1 `cat ${pid_file}`
    echo "`date '+%Y-%m-%d %H:%M:%S'` --> 服务重载完成"
}

function compress_log () {
    echo "`date '+%Y-%m-%d %H:%M:%S'` --> 开始压缩日志"
    tar czf ${log_file}-${date}.tar.gz ${log_file}-${date} --remove-files
    result=`echo $?`
    if [ $result == 0 ];then
        echo "`date '+%Y-%m-%d %H:%M:%S'` --> 日志压缩完成"
    else
        echo "`date '+%Y-%m-%d %H:%M:%S'` --> 日志压缩失败"
    fi
}

function delete_log () {
    echo "`date '+%Y-%m-%d %H:%M:%S'` --> 开始删除老日志"
    find ${log_dir} -mtime +15  -a  -type f -name "*.tar.gz" -exec  rm -f {} \;
    result=`echo $?`
    if [ $result == 0 ];then
        echo "`date '+%Y-%m-%d %H:%M:%S'` --> 日志删除完成"
    else
        echo "`date '+%Y-%m-%d %H:%M:%S'` --> 日志删除失败"
        exit  100
    fi
}

function main () {
    echo ""
    echo "`date '+%Y-%m-%d %H:%M:%S'` --> 任务开始"
    cut_log
    compress_log
    delete_log
    echo "`date '+%Y-%m-%d %H:%M:%S'` --> 任务结束"
}

main
