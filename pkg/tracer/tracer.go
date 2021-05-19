package tracer
//Jaeger 链路追踪

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

//新建Jaeger链路追踪器
func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer,
	io.Closer, error) {
	//jaeger 客户端配置
	cfg := &config.Configuration{
		ServiceName: serviceName,
		//采样配置
		Sampler: &config.SamplerConfig{
			//常量采样:采样器始终对所有traces做出相同的决定
			//Ref:https://rocdu.gitbook.io/jaeger-doc-zh/architecture/sampling
			Type:  "const",
			//采样所有链路
			Param: 1,
		},
		//报告器配置
		Reporter: &config.ReporterConfig{
			LogSpans:            true,	//日志记录Span
			BufferFlushInterval: 1 * time.Second,	//缓冲区刷新频率
			LocalAgentHostPort:  agentHostPort,		//上报的代理主机及端口
		},
	}
	//初始化链路追踪对象
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	//设置全局链路追踪对象
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
