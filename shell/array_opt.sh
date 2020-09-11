#!/bin/bash
# banqinghong@126.com
# 数组初始化 数组遍历 字符串转换为数组



# 数组初始化: my_array=(A B "C" D)
# ${var//pattern/substr} 
# 全部。查找var所表示的字符串中，所有能被pattern所匹配到的字符串，以substr替换之

str="aaa,bbb,ccc,ddd"

arr=(${str//,/ })
for st in ${arr[*]};do
    echo "string = $st"
done
