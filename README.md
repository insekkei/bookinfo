你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# bookinfo

Get book info from douban.com by using isbn13 or isbn10, It also can generate book summary info and
save them to file.

## prepare works

1、Firstly, install golang...

	Download the newest package here: http://www.golangtc.com/download

2、Now the go command became available...

3、config the path:
	
	$ vim ~/.bash_profile
	
	# GOPATH is an existed content
	export GOPATH=/Users/.../works
	export GOBIN=$GOPATH/bin
	export PATH=$PATH:$GOBIN

	# Remember to source the new file...
	$ source ~/.bash_profile

4、Get godep

	$ go get github.com/tools/godep

----------------------------------------------------------------------------------

## Usage

Clone the source code

	$ git clone https://github.com/insekkei/bookinfo.git

Get into bookinfo, `$ godep  go build`, then bookinfo appears~~

###How to use?

Get book info

	$ ./bookinfo get -i [isbn] -o [filename]

Generate book summary

	# the default dir is books
	$ ./bookinfo generate -d [dir] -o [filename]

`help` is always there...

### Example

Get book info which isbn13 is `9787508623016` and save the result to `book.json`

	$ ./bookinfo get -i 9787508623016 -o book.json

Generate book summary info and save the result to `books.json`

	$ ./bookinfo generate -o books.json

## Batch action

	$ ./update_book_data.sh


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

