"""
This file aims to convert multifle files in a same folder into another format.
For example,you can convert utf-16 to utf-8
"""
import io
import os

# directory that contains files
dir_ = "YourDirectory/"

# initial input format
input_encoding = "utf-16 LE"

# target output format
output_encoding = "utf-8-sig"

files_list = []
for root, dirs, files in os.walk(dir_, topdown=False):
    for name in files:
        files_list.append(os.path.join(root, name))
    print (files_list)

for filename in files_list:
    with io.open(filename,encoding=input_encoding) as f:
        text = f.read()
    # process Unicode text
    with io.open(filename,'w',encoding=output_encoding) as f:
        f.write(text)
