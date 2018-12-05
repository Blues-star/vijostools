# vijostools
> powered by 星夜的蓝天 www.poi.ac  

![欢迎来我的主站玩耍](https://i.loli.net/2018/12/05/5c07655a514ae.jpg)
## 描述
vijos数据文件打包器，自用
## 用法
将exe文件放置在数据目录里，打开即可。exe需要手动构建，或者直接使用python脚本也可以    

golang使用方法：
```bat
go run main.go
```
## 版本状态
1.0到3.0版本为python脚本，在不装python的情况下难以支持跨平台（主要是XP），所以5.0版本使用golang重写。性能也得到了提升

## 已知bug
- [ ] 不支持无题目名的数据点
- [ ] 报错时需要使用cmd环境下打开才能知道报错信息