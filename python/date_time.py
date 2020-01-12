#!/usr/bin/python
import time 
import datetime
from datetime import date


for plus_time   in  range(0,10):
    time_day=(datetime.datetime.now()+datetime.timedelta(days=plus_time)).strftime('%Y-%m-%d')
    time_day_week=(datetime.datetime.now()+datetime.timedelta(days=plus_time)).strftime('%w')
    print ("time day is :%s   week_day:%s "%(time_day,time_day_week))
    #print ('datw  +%s'%(time_day.weekday()))
last_month=""
for plus_time   in  range(0,100):    
    now_month=(datetime.datetime.now()+datetime.timedelta(days=plus_time)).strftime('%Y-%m')
    if last_month != now_month:
        print ("now_month is %s"%(now_month))
        last_month=now_month
    else:
        continue;
