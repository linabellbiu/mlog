# mlog

#### 一个超轻量日志输出的包

- install package
```
 go get github.com/wangxudong123/mlog
```


- use
```go               
import "github.com/wangxudong123/mlog"

func main(){
    m := mlog.New("./log/")     
    
    //[INFO]: 2020/04/02 11:31:10 main.go:8: 这是一个输出信息
    m.Info("这是一个输出信息")      
    
    //[WARNING]: 2020/04/02 11:31:10 /Users/xudong/code/my/githup/main.go:10: 这是一个警告的错误
    m.Warning("这是一个警告的错误")   
    
    //[ERROR]: 2020/04/02 11:31:10.014125 /Users/xudong/code/my/githup/main.go:12: 这是一个致命的错误
    m.Error("这是一个致命的错误")     

}
```
