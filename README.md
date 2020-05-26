# kuu-locale

[![Build Status](https://travis-ci.org/kuuland/locale_cli.svg?branch=master)](https://travis-ci.org/kuuland/locale_cli)

## Usage 

```shell script
kuu-locale dir1 dir2 ...
# OR
DIRS=dir1,dir2 kuu-locale
```

## Examples

```shell script
~ kuu-locale /example/myapp
总耗时：181.675484ms
扫描文件数：18
命中声明数：44
详细日志文件：locale.log
翻译参考文件：locale.csv

~ kuu-locale /example/myadmin/src
总耗时：71.617358ms
扫描文件数：23
命中声明数：194
详细日志文件：locale.log
翻译参考文件：locale.csv

~ kuu-locale /example/myapp /example/myadmin/src
总耗时：291.343124ms
扫描文件数：41
命中声明数：238
详细日志文件：locale.log
翻译参考文件：locale.csv

~ DIRS=/example/myapp,/example/myadmin/src kuu-locale
总耗时：265.100336ms
扫描文件数：41
命中声明数：238
详细日志文件：locale.log
翻译参考文件：locale.csv
```

## Environments

```shell script
EXTS=.go,.js,.jsx,.ts,.tsx,.vue # default
```

```shell script
DIRS=/example/dir1,/example/dir2,/example/dir3
```