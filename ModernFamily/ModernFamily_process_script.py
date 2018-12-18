# -*- coding:utf-9 -*-

"""
You must make sure all files are encoded with utf-8 format
If not,you can use `change_file_format.py` to convert.
"""

import os
import re

# 提取出中英文的正则表达式
script_pattern = re.compile("^.*,,(.*?)\\N{.*}(.*)")

dir_ = "directory which contains file"
target_file_name = "You target filename.suffix"
# 标明第几季。因为字幕名称不统一所以无法从文件名处判断
episode = 1


def main():
    files_list = read_os()
    for index, file_ in enumerate(files_list):
        deal_with_content(index, file_)


def read_os():
    files_list = []
    for root, dirs, files in os.walk(dir_, topdown=False):
        for name in files:
            files_list.append(os.path.join(root, name))
    print(files_list)
    return files_list


def deal_with_content(index, file_name):
    """

    :param index: chapter number
    :param file_name: episode script filename
    :return:
    """
    with open(file_name, encoding='utf-8') as input_:
        with open(target_file_name, 'a+', encoding='utf-8') as output_:
            title = "0{}".format(index + 1) if index < 9 else index + 1
            output_.write("S0{}E{} \n----------------------\n".format(episode, title))
            for line in input_:
                # 有些字幕会存在推广内容，寻找关键字并屏蔽
                if not any(x in line for x in ["更多影视更新", "本字幕首发于", "打开微信","最新美剧下载","海外影视剧下载",
                                               "字幕组","最快首发"]):
                    try:
                        if line.startswith('Dialogue'):
                            chs = script_pattern.match(line).group(1).strip('\\-')
                            eng = script_pattern.match(line).group(2).strip('\\-')
                            output_.write("- {}\n".format(eng))
                            output_.write(" {} \n\n".format(chs))
                    except AttributeError:
                        pass


if __name__ == '__main__':
    main()
