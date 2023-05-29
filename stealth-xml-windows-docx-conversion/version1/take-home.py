import zipfile
import re

import xml.dom.minidom


class Solution(object):
    head1_pattern = re.compile(r'<w:pStyle w:val="Heading1"/>.*?<w:t>(.*?)</w:t>', flags=re.DOTALL)
    head2_pattern = re.compile(r'<w:pStyle w:val="Heading2"/>.*?<w:t>(.*?)</w:t>', flags=re.DOTALL)
    # make it a configurable variable later
    input_file_name = 'windows8.docx'

    def __init__(self):
        self.header_map = {}

    def read_input_and_print_output(self):
        # step 1 : read the input
        document = zipfile.ZipFile(self.input_file_name)
        uglyXml = xml.dom.minidom.parseString(document.read('word/document.xml')).toprettyxml(indent='  ')

        with open('input.txt', 'w+') as output_:
            for line in uglyXml:
                try:
                    output_.write(line)
                except Exception as ex:
                    pass

        flag = False
        # step 2 : map the header1 with header2 in different sections
        with open('input.txt') as input_:
            with open('output.txt', 'w+') as output_:
                for line in input_:
                    if (line.find("Heading1") != -1):
                        if (not flag):
                            flag = True
                            output_.write(line)
                            continue
                        else:
                            self._process_file()
                            output_.truncate(0)
                    output_.write(line)
            self._process_file()

        # step3 : convert output to HTML
        self._print_output_to_html()

    def _process_file(self):
        with open('output.txt') as input_:
            doc = input_.read()
            header1 = self.head1_pattern.findall(doc)
            print(header1)
            header2 = self.head2_pattern.findall(doc)
            print(header2)
            self.header_map[header1[0]] = header2

    def _print_output_to_html(self):
        # print(self.header_map)
        html_pattern = """<ul>
                            {}
                            <ul>
                               {}
                                </ul>
                            </ul>"""
        print("-----------------final output will be :-----------------")
        with open('output.html', 'w+') as output_:
            for header1, header2 in self.header_map.items():
                header1_list = f'<li>{header1}</li>'
                if not header2:
                    header2_list = ""
                else:
                    header2_list = "".join(f"<li>{item}</li>" for item in header2)
                doc_html = html_pattern.format(header1_list, header2_list)
                print(doc_html)
                output_.write(doc_html)


if __name__ == '__main__':
    s = Solution()
    s.read_input_and_print_output()
