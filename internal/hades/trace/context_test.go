package trace

/*func TestContext(t *testing.T) {
	_, closer := Init("room.backend.service")
	defer closer.Close()
	taskLen := 100
	for i := 0; i < taskLen; i++ {
		// Notice: Every iteration need create new trace context
		ctx := Context()
		span := opentracing.SpanFromContext(ctx)
		sp, _ := span.(*jaeger.Span)
		//assert.Equal(t, "Backend root", sp.OperationName())
		t.Logf("oper name:%s", sp)
	}
	ctx := Context("HTTP Client Get /api/get")
	span := opentracing.SpanFromContext(ctx)
	sp, _ := span.(*jaeger.Span)
	//assert.Equal(t, "HTTP Client Get /api/get", sp.OperationName())
	t.Logf("oper name:%s", sp)
}
*/
