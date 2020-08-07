#!/usr/bin/python
import pdb
import paramiko
import  sys

def exec_sh_remote_host(host_name,user_name,passwd):
    s = paramiko.SSHClient()
    s.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    s.connect(hostname=host_name,username=user_name,port=36000,password=passwd)
    sftp = s.open_sftp()
    sftp.put('./remort_host_check.sh', '/data/remort_host_check.sh')

    sftp.close()
    stdin,stdout,stderr=s.exec_command("cd /data/;sh remort_host_check.sh")
    #print(stdin.read().decode())
    print(stdout.read().decode())
    
    #print("error"+stderr.read().decode())
    s.close()

    exit  

def main():
    file_name=sys.argv[1]
    fh=open(file_name)
    log_name=file_name+".log"
    file_log=open(log_name,"w")
    for line  in  fh.readlines():
        host_name=line.split(' ')[0]
        user_name=line.split(' ')[1]
        pass_wd=line.split(' ')[2].strip()
        print(pass_wd)
        exec_sh_remote_host(host_name,user_name,pass_wd)
if __name__ == "__main__":
    main()
