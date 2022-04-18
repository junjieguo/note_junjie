#!/bin/sh
###stirng use ";" split 
first_all="$1"
second_all="$2"


first_all_arr=(${first_all//;/ })  
second_all_arr=(${second_all//;/ })  

if [ ${#first_all_arr[@]} -ne ${#second_all_arr[@]} ];then

    echo "first num not eq second num "
    exit 1 
else 
    echo "ok"
fi 

for ((i=0;i<${#first_all_arr[@]};i++))
do
    First="${first_all_arr[$i]}"
    Fsecond="${second_all_arr[$i]}"
    echo "$First $Fsecond"
done
