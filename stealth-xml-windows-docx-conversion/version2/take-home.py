import zipfile
import re

import xml.dom.minidom
import xml.etree.ElementTree as ET


class Solution(object):
    # make it a configurable variable later
    input_file_name = "windows8.docx"

    def __init__(self):
        self.header_map = {}

    def parse_file_with_xml(self):
        """
        Parse the file with XML library
        This is the XML structure:
        <w:body>
            <w:p>
                <w:pPr w:pStyle="Heading1/Heading2">
                    <wr>
                        <wt>Header name</wt>
                    </wr>
                </w:pPr>
            </w:p>
        </w:body>

        And we need to filter the None(empty) value

        """
        document = zipfile.ZipFile(self.input_file_name)
        pretty_xml = xml.dom.minidom.parseString(document.read("word/document.xml")).toprettyxml(indent="  ")

        parser = ET.XMLParser(encoding="utf-8")
        tree = ET.parse("input.xml", parser=parser)
        root = tree.getroot()
        header1 = []
        ns = {"w": "http://schemas.openxmlformats.org/wordprocessingml/2006/main"}

        for item in root.findall("w:body/w:p", ns):
            header = item.find("w:pPr/w:pStyle", ns)
            if header is not None:
                value = item.find("w:r/w:t", ns)
                if value is not None:
                    header = list(header.attrib.values())[0]
                    value = value.text
                    if header == "Heading1":
                        header1.append(value)
                        # We can use a defaultDict instead, yet currently we still use dict
                        # so we donot need to redefine the header_map
                        self.header_map[value] = []
                    else:
                        self.header_map[header1[-1]].append(value)
                    print(f"{header}   \n {value}")
        print(self.header_map)
        self._print_output_to_html()

    def _print_output_to_html(self):
        """
        header_map is a key-value pair like : {header1:[header2,header2]}
        This function will generate the HTML output, fill the template with header_map
        :return:
        """
        html_pattern = """{}{}"""
        print("-----------------final output will be :-----------------")
        with open("output.html", "w+") as output_:
            output_.write("<ul>")
            for header1, header2 in self.header_map.items():
                header1_list = f"<li>{header1}</li>"
                if not header2:
                    header2_list = ""
                else:
                    generated_list = "".join(f"<li>{item}</li>" for item in header2)
                    header2_list = f"<ul>{generated_list}</ul>"
                doc_html = html_pattern.format(header1_list, header2_list)
                print(doc_html)
                output_.write(doc_html)
            output_.write("</ul>")
        print("-----------------\nPlease check it in output.html file")


if __name__ == "__main__":
    s = Solution()
    s.parse_file_with_xml()
