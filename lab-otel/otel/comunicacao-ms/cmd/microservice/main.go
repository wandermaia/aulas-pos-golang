package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/devfullcycle/microservices-demo/internal/web"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func initProvider(serviceName, collectorURL string) (func(context.Context) error, error) {
	ctx := context.Background()

	// Criando nome do recurso para ser utilizando no jaeger, por exemplo
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(serviceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// Alterado para contexto com timeout
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	//Criando a chamada grpc
	conn, err := grpc.DialContext(ctx, collectorURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	// Configurar o exporter do trace com grpc (poderia ser http também)
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// Criar o span em formato batch
	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)

	// Vai fazer a consolidação das informações
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()), // A amostragem que será enviada no trace.
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)

	// Propagar a informação utilizando os dados de tracing
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Shutdown graceful
	return tracerProvider.Shutdown, nil
}

// load env vars cfg
func init() {
	viper.AutomaticEnv()
	// viper.SetDefault("TITLE", "Microservice Demo")
	// viper.SetDefault("BACKGROUND_COLOR", "green")
	// viper.SetDefault("RESPONSE_TIME", "1000")
	// viper.SetDefault("EXTERNAL_CALL_URL", "http://goapp2:8181")
	// viper.SetDefault("EXTERNAL_CALL_METHOD", "GET")
	// viper.SetDefault("REQUEST_NAME_OTEL", "microservice-demo")
	// viper.SetDefault("OTEL_SERVICE_NAME", "microservice-demo")
	// viper.SetDefault("OTEL_EXPORTER_OTLP_ENDPOINT", "otel-collector:4317")
	// viper.SetDefault("HTTP_PORT", ":8080")
}

func main() {

	// Sinais para graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Shutdown do provider
	shutdown, err := initProvider(viper.GetString("OTEL_SERVICE_NAME"), viper.GetString("OTEL_EXPORTER_OTLP_ENDPOINT"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

	// Criação do tracer, que vai realmente realizer o tracing do código
	tracer := otel.Tracer("microservice-tracer")

	// Dados para a criação do servidor
	templateData := &web.TemplateData{
		Title:              viper.GetString("TITLE"),
		BackgroundColor:    viper.GetString("BACKGROUND_COLOR"),
		ResponseTime:       time.Duration(viper.GetInt("RESPONSE_TIME")),
		ExternalCallURL:    viper.GetString("EXTERNAL_CALL_URL"),
		ExternalCallMethod: viper.GetString("EXTERNAL_CALL_METHOD"),
		RequestNameOTEL:    viper.GetString("REQUEST_NAME_OTEL"),
		OTELTracer:         tracer,
	}

	// Criação do server
	server := web.NewServer(templateData)
	router := server.CreateServer()

	// Servidor inciiando em outra thread
	go func() {
		log.Println("Starting server on port", viper.GetString("HTTP_PORT"))
		if err := http.ListenAndServe(viper.GetString("HTTP_PORT"), router); err != nil {
			log.Fatal(err)
		}
	}()

	// Select para realizar o gracefull shutdown
	select {
	case <-sigCh:
		log.Println("Shutting down gracefully, CTRL+C pressed...")
	case <-ctx.Done():
		log.Println("Shutting down due to other reason...")
	}

	// Create a timeout context for the graceful shutdown
	_, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
}
