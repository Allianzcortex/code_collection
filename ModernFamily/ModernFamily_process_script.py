import os
import re

script_pattern = re.compile("^.*,,(.*?)\\N{.*}(.*)")


def main():
    files_list = read_os()
    for index, file_ in enumerate(files_list):
        deal_with_content(index, file_)


def read_os():
    files_list = []
    top = 'F:\MF_scripts'
    for root, dirs, files in os.walk(top, topdown=False):
        for name in files:
            files_list.append(os.path.join(root, name))
    return files_list


def deal_with_content(index, file_name):
    """

    :param index: chapter number
    :param file_name: episode script filename
    :return:
    """
    with open(file_name, encoding='utf-8') as input_:
        with open('new_title.txt', 'a+', encoding='utf-8') as output_:
            title = "0{}".format(index+1) if index<9 else index+1
            output_.write("S06E{} \n----------------------\n".format(title))
            for line in input_:
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
