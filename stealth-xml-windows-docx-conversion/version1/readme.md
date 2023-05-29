### Basic Idea

1. Use zip to parse doc file into  `xml` 
2. Split the XML file into different parts based on `Heading2` keyword
3. use regex to extract the headers we need 

---

### Steps to use

1. Run the command: `python take-home.py`

2. Final output will be displayed and stored in `output.html`

---

### TODO

1. Its inefficient to parse XML format with regex. I will try to use other high-performance libraries(e.g. `BeautifulSoup`) to parse the XML.

2. Converting doc to XML format is one way to solve the problem. We may find other better ways to do it.

3. Test more docs to make it more generic.