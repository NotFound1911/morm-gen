# morm_gen
 mini orm code generation.  [mini orm](https://github.com/NotFound1911/morm)框架的代码辅助生成工具

## 使用方法
1.下载代码
``` 
git clone https://github.com/NotFound1911/morm_gen.git
```
2.编译代码
``` 
go build -v
```

3.执行代码, 生成模板文件
```
./morm_gen -f input_file.go
```
会在对应路径下生成文件`input_file_gen.go`