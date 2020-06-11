# Govm

Govm 是一个简单的虚拟机。很适合你学习研究

其中顺带写了一门自定义的简单语言，以供学习

作者的编写笔记：https://www.pefish.club/categories/%E5%8A%A8%E6%89%8B%E5%86%99%E8%99%9A%E6%8B%9F%E6%9C%BA/

## Quick start

```shell script
go get -u github.com/pefish/go-vm/bin/govm

govm "
CONST 'Hello World'
PRINT 
halt
"

govm "
CONST 1
CONST 2.1
ADD
PRINT 
halt
"
```


## Security Vulnerabilities

If you discover a security vulnerability, please send an e-mail to [pefish@qq.com](mailto:pefish@qq.com). All security vulnerabilities will be promptly addressed.

## License

This project is licensed under the [Apache License](LICENSE).



