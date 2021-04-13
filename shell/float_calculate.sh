#!/bin/bash
# liyang
# 浮点数计算方法

a=115
total=206

res1=$(echo "scale=4;$a / $total"|bc)
res2=$(echo "scale=4;$a / $total * 100"|bc)
echo "$res1"
echo "$res2"

awk 'BEGIN{printf "%.2f\n",'$a'/'$total'*100}'


