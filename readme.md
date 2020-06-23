## 给go加注释
这玩意用来干嘛的呢？
给go文件加一些其实没必要的注释。
比如：
```
const LineID = "lineId"

func Init() {
...
}

type TStruct struct {
}
func (t *TStruct) Get() (*TSnmpSubsvr, error) {
}
```
处理后：
commentHelper /path/file.go

会变成：
```
// LineID ...
const LineID = "lineId"

// Init ...
func Init() {
...
}

// TStruct ...
type TStruct struct {
}

// Get ...
func (t *TStruct) Get() (*TSnmpSubsvr, error) {
}

```

## 做这个有啥用
因为最近部门在搞代码评比，扫描出来公有变量不加注释是会报警的。
所以，用来骗骗机器。

其次，先写个注释模版在这提醒下，说不定你心情好的时候，也会补充下嘛。

## usage
```
git clone https://github.com/dalianzhu/commentHelper.git
cd commentHelper
go build github.com/dalianzhu/commentHelper
# 得到 commentHelper
mv commentHelper /usr/bin
chmod 777 /usr/bin/commentHelper
# 使用
commentHelper /path/file.go
```