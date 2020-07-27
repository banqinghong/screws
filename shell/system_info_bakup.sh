#!/bin/bash
# banqinghong@126.com
# 服务器需要重启或者其他情况下，备份当前系统的一些信息

# 定义变量
date=`date "+%Y%m%d-%H"`
bak_dir=/root/${date}_bakup

# 新建备份目录
mkdir $bak_dir

#开始备份iptables、route、进程信息。
cd $bak_dir

# 备份路由信息
route >> route.bak

# 备份防火墙信息
iptables-save >> iptables.bak

# 备份进程信息
ps -ef >> process.bak

# 备份域名解析信息
cat /etc/resolv.conf >> resolv.bak

# 备份历史操作命令
history -a
cat ~/.bash_history >> history.bak

# 备份磁盘挂载信息
echo "--- space ---" >> disk.bak
df -h >> disk.bak
echo >> disk.bak
echo "--- inode ---" >> disk.bak
df -i >> disk.bak

# 备份/data目录文件信息
ls -al /data >> data_info.bak

# 备份hostname信息
hostname >> hostname.bak

# 备份hosts文件
cp /etc/hosts ./hosts.bak

# 备份 profile 文件
cat /etc/profile  >> profile.bak

