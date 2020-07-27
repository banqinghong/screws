#!/bin/bash
# liyang@g7.com.cn
# 2019-08-21
# 自动格式化磁盘
# 因为是通过命令 fdisk -l | grep "/dev/vd*" 获取的磁盘列表，所以脚本不能支持所有的磁盘格式化，其他磁盘的格式化需要修改脚本内容

log_file=/var/log/disk_init.log
i=0
b=1

# 判断是否有磁盘需要格式化
disk_count=`fdisk -l | grep "/dev/vd*" | awk '{if($2~/vd/ && $2!~"/dev/vda") print substr($2,0,8)}'|wc -l`
if [ ${disk_count} -lt 1 ];then
    echo "no disk need parted" >> ${log_file}
fi


if [ ${disk_count} -eq 1 ];then
    echo "only one disk found" >> ${log_file}
    disk=`fdisk -l | grep "/dev/vd*" | awk '{if($2~/vd/ && $2!~"/dev/vda") print substr($2,0,8)}'`
    echo "start parting disk --> $disk" >> ${log_file}
    parted $disk  << EXIT
        mklabel gpt
        mkpart primary 0 -1
        ignore
        quit
EXIT
    echo "disk ${disk} parted successful" >> ${log_file}
    mkfs.ext4 $disk$b
    echo "${disk}${b} create file system successful" >> ${log_file}
    sleep 1s
    mkdir /data
    mount ${disk}${b} /data
    uuid=`blkid ${disk}${b} | awk '{print $2}'|awk -F"\"" '{print $2}'`
    echo "UUID=${uuid}      /data      ext4    defaults                0 0 " >> /etc/fstab
    echo "disk ${disk}${b} mounted successful" >> ${log_file}
    # 创建 app目录
    mkdir /data/app; mkdir /data/app_log;
    # 修改目录权限
    chmod 1777 /data/app; chmod 1777 /data/app_log
    # 启动脚本中注释 磁盘格式化脚本
    sed -i 's@/usr/bin/init_disk.sh@#/usr/bin/init_disk.sh@g' /etc/rc.d/rc.local
    exit 0
fi



# 格式化并挂载磁盘
echo "${disk_count} disks found" >> ${log_file}
for  disk in `fdisk -l | grep "/dev/vd*" | awk '{if($2~/vd/ && $2!~"/dev/vda") print substr($2,0,8)}'`;do
    echo "start parting disk --> $disk" >> ${log_file}
    parted $disk  << EXIT
        mklabel gpt
        mkpart primary 0 -1
        ignore
        quit
EXIT
    echo "disk ${disk} parted successful" >> ${log_file}
    mkfs.ext4 $disk$b
    echo "${disk}${b} create file system successful" >> ${log_file}
    sleep 1s
    mkdir /data${i}
    mount ${disk}${b} /data${i}
    uuid=`blkid ${disk}${b} | awk '{print $2}'|awk -F"\"" '{print $2}'`
    echo "UUID=${uuid}      /data${i}      ext4    defaults                0 0 " >> /etc/fstab
    echo "disk ${disk}${b} mounted successful" >> ${log_file}
    let i+=1
done

