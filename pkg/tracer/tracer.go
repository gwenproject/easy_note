package tracer

import (
    "fmt"
    "github.com/gwen0x4c3/easy_note/pkg/constants"
    "github.com/opentracing/opentracing-go"
    "github.com/uber/jaeger-client-go"
    jaegercfg "github.com/uber/jaeger-client-go/config"
)

func InitJaeger(service string) {
    cfg := &jaegercfg.Configuration{
        ServiceName: service,
        Disabled:    false,
        Sampler: &jaegercfg.SamplerConfig{
            Type:              "const",
            Param:             1,
            SamplingServerURL: fmt.Sprintf("http://%s:%d/sampling", constants.JaegerHost, jaeger.DefaultSamplingServerPort),
        },
        Reporter: &jaegercfg.ReporterConfig{
            LogSpans:           true,
            LocalAgentHostPort: fmt.Sprintf("%s:%d", constants.JaegerHost, jaeger.DefaultUDPSpanServerPort),
        },
    }

    tracer, _, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
    if err != nil {
        panic(fmt.Sprintf("Error: cannot init Jaeger: %v\n", err))
    }
    opentracing.SetGlobalTracer(tracer)
}
