package tracer

import (
	"io"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

// NewJaegerTracer 新增链路追踪
func NewJaegerTracer(serviceName string, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	// 为 jaeger client 的配置项，主要设置应用的基本信息，详情见源码
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			// 固定采样
			Type: "const",
			// 所有数据采样
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			// 是否启用
			LogSpans: true,
			// 刷新缓冲区的频率
			BufferFlushInterval: 1 * time.Second,
			// 上报的tracer agent地址
			LocalAgentHostPort: agentHostPort,
		},
	}
	// 根据配置初始化Tracer对象，返回的是通用标准的opentracing.Tracer对象
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	// 设置全局Tracer对象，在一个应用里一般只用一个追踪系统
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, nil
}
