#!/usr/bin/python 
#coding:UTF-8
import time
import requests 
import json
import datetime
import sys
import os 
from requests.structures import CaseInsensitiveDict


def get_error_log_msg(start_time,end_time):
    '''
    get data from monitor.oa.com  .return 2 months average data
    根据intfid确认是否过滤，确定不同的head_data
    '''
    headers = CaseInsensitiveDict()
    headers["Content-Type"] = "application/x-www-form-urlencoded"
    headers["Cookie"] = "username=XXX"
    head_data = '''page=1&pageSize=200&a=%s&b=%s'''%(start_time,end_time)
    #print (head_data)
    #print(headers)
    #head_data='''{"dimMap":{"intfID":"%s"},"millSt":%s,"millEt":%s,"metric":"%s","groups":["%s"],"desc":true, "topN":30}
 #       '''%(intfid,start_time,end_time,request_type,groups_id)
    curl_url="http://xxxxx.com"
    #print (head_data)
    try:
        r=requests.post(curl_url,headers=headers,data=head_data,timeout=2000)
        data_dir=json.loads(r.text)

    except requests.RequestException as e:
        print(e)
    info_dic={}
    errlog_datas=data_dir['data']['data']
    for i, var in enumerate(errlog_datas):
        Acct_Id=var["AcctId"]
        Mon_desc=var["mon_desc"]
        Firstkey=var["mon_desc"].split(';')[0].split('=')[1]
        #print(Acct_Id)
        #print(Firstkey)
        info_dic[Firstkey]=Acct_Id
    #print (info_dic)
    if info_dic:
        for (key ,value) in info_dic.items():
            print ("%s is %s "%(key,value))
            os.popen('ls -lrth  %s %s'%(value,key))


if  __name__ == "__main__":
    now_time=time.mktime(datetime.datetime.now().timetuple())
    end_time=int(now_time)*1000
    begin_time=(int(now_time)-3600-200)*1000
#    print(end_time)
#    print(begin_time)
    get_error_log_msg(begin_time,end_time)
