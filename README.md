# mlog

#### 一个超轻量的日志输出的包

- install package
```
 go get github.com/wangxudong123/mlog
```


- use
```go               
import "github.com/wangxudong123/mlog"

func main(){

    m := mlog.New("./log/")     
    
    m.Info("这是一个输出信息")      
    
    m.Warning("这是一个警告的错误")   
    
    m.Error("这是一个致命的错误")     

}
```
