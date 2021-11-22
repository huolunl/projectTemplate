/*
   @Author:huolun
   @Date:2021/11/22
   @Description
*/
package main

import (
	"github.com/opentracing/opentracing-go"
	traceLog "github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/transport"
	"io"
	"log"
)
const JaegerHostPort = "http://127.0.0.1:14268/api/traces"

func NewJaegerTracer(service string) (opentracing.Tracer, io.Closer) {

	sender := transport.NewHTTPTransport(
		JaegerHostPort,
	)
	tracer, closer:= jaeger.NewTracer(service,
		jaeger.NewConstSampler(true),
		jaeger.NewRemoteReporter(sender))
	return tracer, closer
}

func main() {
	tracer,closer:=NewJaegerTracer("test-client")
	defer func() {
		err := closer.Close()
		if err!= nil{
			log.Println(err)
		}
	}()
	// 直接创建一个span
	span := tracer.StartSpan("span first")
	span.SetTag("a","b")
	span.LogFields(traceLog.String("aaa","bbb"))
	defer  span.Finish()

	childSpan := tracer.StartSpan(
		"span second",
		opentracing.ChildOf(span.Context()),
	)
	childSpan.SetTag("c","d")
	defer childSpan.Finish()
}

