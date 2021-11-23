#!/usr/bin/python
# -*- coding: utf-8 -*-

import requests
from bs4 import BeautifulSoup
import time
import string
import sys
import os
import json
import io

reload(sys)
sys.setdefaultencoding('utf8')

if len(sys.argv) < 3:
    print "Usage:" + sys.argv[0] + "<dir> <filename>"
    sys.exit(0)

directory=sys.argv[1]
filename=sys.argv[2]

books=[]
list = os.listdir(directory)
for i in range(0, len(list)):
    path = os.path.join(directory, list[i])
    with io.open(path, "r", encoding="utf-8") as f:
        s1 = json.load(f)
        title = s1['title']
        image = s1['image']
        author = s1['author']
        pages = s1['pages']
        price = s1['price']
        id= s1['isbn13']
        publisher = s1['publisher']
        book = {'title':title, 'image':image, "id":id, "author": author, "publisher": publisher, "pages":pages, "price":price}
        books.append(book)

with io.open(filename, "w", encoding="utf-8") as f:
    f.write(json.dumps(books, indent=4, ensure_ascii=False))
