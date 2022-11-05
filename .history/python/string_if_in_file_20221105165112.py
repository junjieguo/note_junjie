#!/usr/bin/python
#coding:UTF-8

import argparse
import time
import requests 
import json
import datetime
import sys
import os
import ast
reload(sys)
sys.setdefaultencoding('utf-8')
import os
os.environ['NLS_LANG'] = 'Simplified Chinese_CHINA.ZHS16GBK'

def  table_if_in_while_list(file_name,table_name ):
        fh=open(file_name)
        while fh:
            line=fh.readline().strip('\n')
            if not line:
                return "False"
                break
            if table_name  in line:
                #print (table_name)
                return "True"
                break
            else:
                continue 

if  __name__ == "__main__":
  
    if len(sys.argv) < 2:
        print ("plase  into  function name")
        exit()
    function_name=sys.argv[1]

    file_name=sys.argv[2]

    table_if_in_while_list(function_name,file_name,"skyguo")
    print(table_if_in_while_list(file_name,"skyguo")    )
    print(table_if_in_while_list(file_name,"skyguo1")    )    
    print(table_if_in_while_list(file_name,"guojunjie")    )   
