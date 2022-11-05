 #!/bin/sh
 #2012年11月23日 11:36:16
 #统计特定文件中单词的词频
 if [ $# -ne 1 ];then
 echo "Usage: $0 filename"
 exit -1
 fi

 filename=$1

 egrep -o "\b[[:alpa]]+\b" $filename|awk '{print count[$0]++} END{printf("%-14s%s\n","Word","Count");for(ind in count) {printf("%-14s%d\n",ind,count[ind]);}}'

 #输出为：./word_freq.sh words.txt
