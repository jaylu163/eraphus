package trace

const operationLimit = 512

/*// Context returns a ctx with root span
// only for backend task service
func Context(operationName ...string) context.Context {
	var operation = "Backend root"
	if len(operationName) > 0 && len(operationName[0]) < operationLimit {
		operation = operationName[0]
	}
	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan(operation)
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	tls.SetContext(ctx)
	return ctx
}

// Deprecated init global tracer
func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1.0,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	tracer, closer, err := cfg.NewTracer(func(c *config.Options) {

	}, config.Logger(jaeger.StdLogger))
	opentracing.SetGlobalTracer(tracer)
	if err != nil {
		panic(err)
	}
	return tracer, closer
}
*/
