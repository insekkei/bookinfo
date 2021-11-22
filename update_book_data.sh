#!/usr/bin/env bash
#
# update_book_data.sh
# Copyright (C) 2015 wanglong <wanglong@laoqinren.net>
# 
# Usage:
#	./update_book_data.sh  <filename>
#
# This script update the book data according by the isbn 
# number in the file(defalt name: book.txt)

# The default name of the file which contains isbn number per line.
FILENAME=book.txt

# Get the filename
while [ "$#" -gt 0 ]; do
	case $1 in
		--file)
			shift 1
			FILENAME="$1"
			;;
		*)
			break
			;;
	esac
done

# Create the books dir if it not exist
BOOKS_DIR=$(dirname $0)/books
if [ ! -d ${BOOKS_DIR} ]; then
	mkdir -p ${BOOKS_DIR}
fi

# Read the file and update the book data
while read isbn
do
	FILE=${BOOKS_DIR}/${isbn}.json
	if [ ! -f ${FILE} ]; then
		echo "Get book data :" $isbn
		./bookinfo.py ${isbn} ${FILE}
	else
		echo "Book data :" $isbn "already exists"
	fi
done < $FILENAME

# Generate the summary book data
# and save it to books.json
rm -fr  books/books.json
./summary.py books books/books.json
