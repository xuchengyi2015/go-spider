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

#### 4、曲线救国实现go语言的函数参数默认值
代码写到既简单又清晰，还是比较好懂了。

```$xslt
package tools

type XResult struct {
	Code    int
	Data    interface{}
	Message string
}

var defaultXResult = XResult{
	Code:    0,
	Data:    nil,
	Message: "success",
}

type XResultOption func(*XResult)

func WithCode(code int) XResultOption {
	return func(result *XResult) {
		result.Code = code
	}
}

func WithData(data interface{}) XResultOption {
	return func(result *XResult) {
		result.Data = data
	}
}

func WithMessage(message string) XResultOption {
	return func(result *XResult) {
		result.Message = message
	}
}

// 函数默认参数的Go实现（真麻烦！(╯﹏╰)）
func GetResult(opts ...XResultOption) XResult {
	result := defaultXResult

	for _, o := range opts {
		o(&result)
	}
	return result
}

```

#### 5、对于beego api的统一错误处理

在controller的函数中实现一个匿名函数，类似于：
```$xslt
// @router /start_crawl [get]
func (m *MovieInfoController) CrawlMovie() {
	defer func() {
		if err := recover(); err != nil {
			m.Data["json"] = tools.GetResult(
				tools.WithCode(1),
				tools.WithMessage(fmt.Sprint(err)),
			)
			m.ServeJSON()
			return
		}
	}()
	
	...其他代码
	
	m.Data["json"] = tools.GetResult()
        m.ServeJSON()
}
```