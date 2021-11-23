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
    print "Usage:" + sys.argv[0] + " <isbn> <filename>"
    sys.exit(0)


headers = {
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36'
}

isbn=sys.argv[1]
filename=sys.argv[2]
response = requests.get('https://book.douban.com/isbn/'+isbn, headers=headers)
soup = BeautifulSoup(response.text, 'lxml')
#soup = BeautifulSoup(open(isbn))

### 书名基本信息
title=""
author=""
pubdate=""
publisher=""
pages=""
price=""
isbn13=""
author_intro=""
summary=""
image=""

title=soup.h1.find('span').string
print "title:"+title
info=soup.find('div', id='info').get_text()
info=info.replace(" ", "").replace("\n", "")
info=info.replace("作者:", "$作者:")
info=info.replace("出版社:", "$出版社:")
info=info.replace("副标题:", "$副标题:")
info=info.replace("出版年:", "$出版年:")
info=info.replace("页数:", "$页数:")
info=info.replace("定价:", "$定价:")
info=info.replace("装帧:", "$装帧:")
info=info.replace("原作名:", "$原作名:")
info=info.replace("译者:", "$译者:")
info=info.replace("出品方:", "$出品方:")
info=info.replace("ISBN:", "$ISBN:")
info=info.replace("丛书:", "$丛书:")

for line in info.split('$'):
    if line.startswith("作者"):
        author=line.split(":", 1)[1]
        print "author:"+author
    if line.startswith("出版年"):
        pubdate=line.split(":", 1)[1]
        print "pubdate:"+pubdate
    if line.startswith("出版社"):
        publisher=line.split(":", 1)[1]
        print "publisher:"+publisher
    if line.startswith("页数"):
        pages=line.split(":", 1)[1]
        print "pages:"+pages
    if line.startswith("定价"):
        price=line.split(":", 1)[1]
        print "price:"+price
    if line.startswith("ISBN"):
        isbn13=line.split(":", 1)[1]
        print "isbn13:"+isbn13

### 图片
img=soup.find('a', class_="nbg").find("img")["src"]
image=img.replace("subject/s", "subject/m")
print "image:"+image


### 内容简介
summary1=soup.find('div', class_="indent", id="link-report")
if summary1:
    summary2=summary1.find_all("p")
    s=""
    for p in summary2:
        if p.string == None:
            continue
        if s != "":
            s=s+"\\n"+p.string
        else:
            s=s+""+p.string
    print "summary:\n\t"+s
    summary=s
else:
    print "summary: null"

### 作者简介
releated_info=soup.find('div', class_="related_info").find_all('div', class_='indent')
for r in releated_info:
    if not r.has_attr('id'):
        author_intro1 = r.find('div', class_="intro")
        if author_intro1:
            a=author_intro1.find_all('p')
            aa=""
            for p in a:
                if p.string == None:
                    continue
                if aa != "":
                    aa=aa+"\\n"+p.string
                else:
                    aa=aa+p.string
            print "author_intro:\n\t"+aa
            author_intro=aa

### TAGS
tags_list=[]
tags1=soup.find('div', class_="blank20", id="db-tags-section")
if tags1:
    tags=tags1.find('div', class_='indent').find_all("span")
    print "tags:"
    for t in tags:
        s=t.find('a').string
        tags_list.append(s)
        print "\t"+s

book={"author":author,"pubdate":pubdate,"image":image, "pages":pages, "isbn13":isbn13, "title":title,"author_intro":author_intro, "summary":summary, "publisher":publisher, "price":price, "tags": tags_list}

with io.open(filename, "w", encoding="utf-8") as f:
    print json.dumps(book, indent=2, ensure_ascii=False)
    f.write(json.dumps(book, indent=2, ensure_ascii=False))

