# 概述
自己开发公用工具集合

## 安装
> brew tap dinglevin/pub-tools
> 
> brew install levin-pub-tools

## 使用
### appendstr
A small tool to append strings to a file, so it can change its MD5 hash

Usage:
  appendstr [flags]

Flags:
  -a, --append string   Append string to the file (default "@@LEVIN@@")
  -h, --help            help for appendstr

### getlastn
A small tool to get last n bytes from a file

Usage:
  getlastn [flags]

Flags:
  -h, --help      help for getlastn
  -n, --num int   Number of last bytes to read (default 9)

