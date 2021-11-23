# bookinfo 介绍

* bookinfo.py 根据isbn13，抓取豆瓣网上图书的信息。
* summary.py 将所有的图书信息，抽取部分信息，汇总到一个json文件中。


# 如何使用?

## 安装运行时的依赖库

```
$ sudo easy_install requests
$ sudo easy_install beautifulsoup4==4.9.3
$ sudo easy_install  lxml
```

## 获取图书信息

```
$ ./bookinfo.py  <isbn> <filename>
```
### 示例：
```
[root@localhost bookinfo]# ./bookinfo.py 9787100086820 9787100086820.json
[root@localhost bookinfo]# cat 9787100086820.json
{
  "pubdate": "2012-6", 
  "tags": [
    "尼采", 
    "哲学", 
    "西方哲学", 
    "德国", 
    "美学", 
    "悲剧", 
    "商务印书馆", 
    "古希腊"
  ], 
  "price": "25.00元", 
  "pages": "195", 
  "publisher": "商务印书馆", 
  "isbn13": "9787100086820", 
  "author": "[德]尼采", 
  "title": "悲剧的诞生", 
  "author_intro": "", 
  "summary": "《悲剧的诞生》是对20世纪的精神生活起了最大影响的思想家。《悲剧的诞生》是尼采一鸣惊人的巨作，也是读者理解尼采美学和哲学的入门书，尼采自称这是一本为那些兼有分析和反省能力的艺术家写的书，充满心理学的创见和艺术的奥秘，是“一部充满青年人的勇气和青年人的忧伤的青年之作”。在这部著作中，尼采用日神阿波罗和酒神狄奥尼索斯的象征说明了艺术的起源、本质、功用以及人生的意义。\\n《悲剧的诞生》一书的最独特之处是对古希腊酒神现象的极端重视。 这种现象基本上靠民间口头秘传，缺乏文字资料，一向为正宗的古典学术所不屑。尼采却立足于这种不登大雅之堂的现象，把它当作理解高雅的希腊悲剧、希腊艺术、希腊精神的钥匙，甚至从中提升出了一种哲学来。他能够凭借什么来理解这种史料无征的神秘现象呢？只能是凭借猜测。然而，他不是凭空猜测，而是根据自己的某种体验，也就是上述所谓“一种被确证的、亲身...\\n(展开全部)\\n《悲剧的诞生》是对20世纪的精神生活起了最大影响的思想家。《悲剧的诞生》是尼采一鸣惊人的巨作，也是读者理解尼采美学和哲学的入门书，尼采自称这是一本为那些兼有分析和反省能力的艺术家写的书，充满心理学的创见和艺术的奥秘，是“一部充满青年人的勇气和青年人的忧伤的青年之作”。在这部著作中，尼采用日神阿波罗和酒神狄奥尼索斯的象征说明了艺术的起源、本质、功用以及人生的意义。\\n《悲剧的诞生》一书的最独特之处是对古希腊酒神现象的极端重视。 这种现象基本上靠民间口头秘传，缺乏文字资料，一向为正宗的古典学术所不屑。尼采却立足于这种不登大雅之堂的现象，把它当作理解高雅的希腊悲剧、希腊艺术、希腊精神的钥匙，甚至从中提升出了一种哲学来。他能够凭借什么来理解这种史料无征的神秘现象呢？只能是凭借猜测。然而，他不是凭空猜测，而是根据自己的某种体验，也就是上述所谓“一种被确证的、亲身经历的神秘主义”。对于这一点，尼采自己有清楚的意识。还在写作此书时，一个朋友对他的酒神理论感到疑惑，要求证据，他在一封信中说：“证据怎样才算是可靠的呢？有人在努力接近谜样事物的源头，而现在，可敬的读者却要求全部问题用一个证据来办妥，好像阿波罗亲口说的那样。”在晚期著述中，他更明确地表示，在《悲剧的诞生》中，他是凭借他“最内在的经验”理解了“奇异的酒神现象”，并“把酒神精神转变为一种哲学激情”。", 
  "image": "https://img2.doubanio.com/view/subject/m/public/s11098302.jpg"
}
```

## 汇总图书信息

```
$ ./summary.py <dir> <filename>

```
### 示例：
```
[root@localhost bookinfo]# ls -l books
total 40
-rw-r--r--  1 root root    590 Nov 23 11:34 9787040209280.json
-rw-r--r--  1 root root   5511 Nov 23 11:34 9789570846300.json
-rw-r--r--  1 root root   2092 Nov 23 11:34 9789576219726.json
-rw-r--r--  1 root root   2262 Nov 23 11:34 9789866562747.json
[root@localhost bookinfo]# ./summary.py books books/books.json 
[root@localhost bookinfo]#  ls -l books/books.json
-rw-r--r--  1 root root 1070 Nov 23 11:54 books/books.json
cat books/books.json
[
  {
    "publisher": "聯經出版公司", 
    "author": " 葛亮", 
    "image": "https://img1.doubanio.com/view/subject/m/public/s28373628.jpg", 
    "title": "北鳶", 
    "id": "9789570846300", 
    "price": "NT390", 
    "pages": "568"
  }, 
  {
    "publisher": "獨步文化", 
    "author": " [日]道尾秀介", 
    "image": "https://img9.doubanio.com/view/subject/m/public/s4546485.jpg", 
    "title": "鼠男", 
    "id": "9789866562747", 
    "price": "NT", 
    "pages": "310"
  }, 
  {
    "publisher": "高等教育出版社", 
    "author": "[美]西尔伯查茨(Silberschatz，A.)", 
    "image": "https://img3.doubanio.com/view/subject/m/public/s26014820.jpg", 
    "title": "操作系统概念", 
    "id": "9787040209280", 
    "price": "72.00元", 
    "pages": "921"
  }, 
  {
    "publisher": "天下文化", 
    "author": "保罗科尔赫", 
    "image": "https://img9.doubanio.com/view/subject/m/public/s29283125.jpg", 
    "title": "我坐在琵卓河畔，哭泣", 
    "id": "9789576219726", 
    "price": "NT", 
    "pages": ""
  }
]
```

## 批量操作

将要获取的图书的`isbn13`的信息，写入到`book.txt`文件中，执行如下命令，会自动抓取图书信息，并保存在`books`目录下，同时会在`books`目录下，生成汇总的文件`books.json`。
```
	$ ./update_book_data.sh
```

# License

Copyright (c) 2014-2015 

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

