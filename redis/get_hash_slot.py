#!/usr/bin/python
import crc16
import sys

def get_key(key_values):
     hash_key=(crc16.crc16xmodem(key_values))%16384
     return hash_key

if  __name__ == "__main__":
    if len(sys.argv)==2:    
        print (get_key(sys.argv[1]))
    else:
        print("ERROR  plase enter string")
