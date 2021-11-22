# bookinfo

Get book info from douban.com by using isbn13 or isbn10, It also can generate book summary info and
save them to file.


# How to use?

Get book info

	$ ./bookinfo.py  <isbn> <filename>

Generate book summary

	$ ./summary.py <dir> <filename>

# Example

Get book info which isbn13 is `9787508623016` and save the result to `book.json`

	$ ./bookinfo.py  9787508623016  book.json

Generate book summary info and save the result to `books.json`

	$ ./summary.py books books/books.json

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

