package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/vastrock-huang/gotour-blogservice/global"
)

//链路追踪中间件
func Tracing() func(c *gin.Context) {
	return func(c *gin.Context) {
		var newCtx context.Context	//上下文
		var span opentracing.Span	//链路Span
		//使用GlobalTracer()返回全局链路追踪器
		//按给定格式从carrier对象中提取Span上下文
		if spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,	//HTTP首部格式
			//由gin请求首部构成的HTTP Carrier对象
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		); err != nil {		//出错
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),	//请求的上下文
				global.Tracer,	//全局链路追踪器
				c.Request.URL.Path,	//请求的URL路径作为Span的操作名OperationName
			)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{
					Key:   string(ext.Component),
					Value: "HTTP",
				},
			)
		}
		defer span.Finish()	//设置Span的结束时间戳已经完成其终止状态

		var traceID string
		var spanID string
		var spanContext = span.Context()
		switch spanContext.(type) {		//断言类型
		case jaeger.SpanContext:	//Span上下文类型
			jaegerContext:=spanContext.(jaeger.SpanContext)
			traceID = jaegerContext.TraceID().String()	//链路ID
			spanID = jaegerContext.SpanID().String()	//链路中该Span的ID
		}
		//设置链路ID和SpanID到上下文中
		c.Set("X-Trace-ID",traceID)
		c.Set("X-Span-ID",spanID)

		c.Request = c.Request.WithContext(newCtx)	//更新请求上下文
		c.Next()
	}
}
