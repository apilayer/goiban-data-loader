goiban-data-loader
=======

Used to load data into the MySQL db used at openiban.com.

Currently supported data sources:

- bundebank
- lu
- nbb
- nl
- at
- ch
- li

Setting up the database
-------

```bash
$ DATABASE_URL="mysql_user:password@tcp(localhost)/goiban?charset=utf8" make migrate
```

Providing new data
-------

Data files provided by some european bank institutes (e.g Austria and Germany) is sometimes provided as a ISO-8859-1 encoded file.
It should be converted to UTF-8 before being committed to the repository.

This is possible using `iconv`.

```bash
# Austria
$ iconv -f iso-8859-1 -t utf8 at_original.csv > at.csv
# Germany
$ iconv -f iso-8859-1 -t utf8 bundesbank_original.txt > bundesbank.txt
```

Running
-------

You should load the data into a database called 'goiban'.

```bash
$ go run loader.go bundesbank "mysql_user:password@tcp(localhost)/goiban?charset=utf8"
```

or to load all data

```bash
$ DATABASE_URL="mysql_user:password@tcp(localhost)/goiban?charset=utf8" make load
```

MySQL development instance
-------
To run MySQL inside a docker container you can use the following command:

`docker run -d --name openiban-mysql -p3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=goiban mysql`

The MIT License (MIT)
---------------
Copyright (c) 2014 Chris Grieger

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
