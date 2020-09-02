#/bin/bash
# banqinghong@126.com
# 根据输入的日期区间，输出一段连续的日期
# 日期格式为 YYYY-mm-dd

if [ $# -ne 2 ]; then
    echo -e "usage:\n  $0 <start_date> <end_date>\nExample:\n  $0 2020-08-01 2020-08-31"
    exit 2
fi

start_date=`date -d "$1" +%Y%m%d`
end_date=`date -d "$2" +%Y%m%d`

if [ ${start_date} -ge ${end_date} ];then
    echo "The start date should be less than the end date"
fi

while [ ${start_date} -le ${end_date} ];do
    if [ ${start_date} -lt ${end_date} ];then
        date=`date -d "${start_date}" +%Y-%m-%d`
        echo $date
        start_date=`date -d "+1 day ${start_date}" +%Y%m%d`
    else
        date=`date -d "${end_date}" +%Y-%m-%d`
        echo $date
        exit 2
    fi
done
