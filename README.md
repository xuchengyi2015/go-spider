# go-spider
learn spider by golang

#### 1、运行程序 bee run 
> 注意：如果第一次请求是404的话，有可能是beego在根据router.go与各个controller的注解路由生成`commentsRouter_controllers.go`。此时，只需要ctrl+c停止项目并重新运行bee run即可。

#### 2、在手写controller的时候为避免新手踩坑，都是controllers的扩展方法
```
func (m *MovieInfoController)GetMovieInfo() 
而不是
func GetMovieInfo(m *MovieInfoController)
```

#### 3、正则匹配使用非贪婪模式
```$xslt
reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*?)</span>`)
```
> 参考代码：// https://www.cnblogs.com/wt645631686/p/9702572.html