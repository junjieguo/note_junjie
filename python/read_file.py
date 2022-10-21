#!/usr/bin/python
f = open("need_redo_1020")


for  lines  in f.readlines():
    print (lines)
    key=lines.split("\t")[0]
    tag=lines.split("\t")[1]
    mes=lines.split("\t")[2].strip('\n')
    print(mes)
    filename=key+".data"
    f = open(filename, "a")
    messages='''{"service":"AAA", "function":"process", "data":{"KEY":"%s","TAG":"%s","MESSAGE":"%s"}}'''%(key,tag,mes)
    print(messages)
    f.write(messages)
    f.close()
f.close()
