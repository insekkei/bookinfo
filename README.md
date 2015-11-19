# bookinfo

Get book info from douban.com by using isbn13 or isbn10

## Usage

1、首先安装go语言：

	下载最新版本安装包：http://www.golangtc.com/download

2、安装完成后go命令就有效了...

3、配置环境变量：`$ vim ~/.bash_profile`
	
	# GOPATH 是已有的一个目录
	export GOPATH=/Users/.../works
	export GOBIN=$GOPATH/bin
	export PATH=$PATH:$GOBIN

4、`$ source ~/.bash_profile`

5、`$ go get github.com/tools/godep`

6、git clone https://github.com/insekkei/bookinfo.git

进入bookinfo目录，执行`$ godep  go build`，之后，当前目录下就能看到bookinfo文件了。

----------------------------------------------------------------------------------

Get book info which isbn13 is `9787508623016` and save the result to `book.json`

	$ ./bookinfo -i [isbn] -o [filename]

	# 例如：

	$ ./bookinfo -i 9787508623016 -o book.json


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

