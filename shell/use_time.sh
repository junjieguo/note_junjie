#!/bin/sh

begin_time="$1"
tar_days=$2
for((i=0;i<$2;i++))
do
        deal_day=`date  -d " ${begin_time} $i days  ago" "+%Y-%m-%d"`
        let "bf=i+1"
        let "af=i+2"
        deal_yestarday=`date  -d " ${begin_time} $bf days  ago" "+%Y-%m-%d"`
        deal_afterday=`date  -d " ${begin_time} $af days  " "+%Y-%m-%d"`
        echo "${deal_yestarday} ${deal_day} ${deal_afterday}"
done 
