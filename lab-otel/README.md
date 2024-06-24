# Laboratório de Observabilidade e Open Telemetry

O OTEL é o segundo projeto mais popular da CNCF (perdendo apenas para o Kubernetes)


## O que é Observabilidade

Monitoramento e Observabilidade são complementares, mas são diferentes.


- Monitoramento nos mostra que há algo errado

- Monitoramento se baseia em saber com antecedência quais sinais você deseja monitorar.

- Observabilidade nos permite perguntar porquê

- Monitoramento é "O quê" e a Observabilidade é o "porquê"


### 3 Pilares da Observabilidade

- Métricas - Qualquer coisa que possa ser medida, normalmente de forma absoluta. POdem ser de sistemas ou de negócio.

- Logs - São os eventos no passado, o que aconteceu.

- Tracing - Rastrear a requisição, olhar a aplicação por dentro, passando por todas as camdas. Muito utilizado de forma distribuída.

## Contexto Open Telemetry

Observabilidade
    - Logs
    - Métricas
    - Tracing
    - Cloud Native Telemetry

Centralização da Informação
Necessidade de customização das Informações
    - Geração de métricas de Negócio
    - Tracing das rotinas e blocos internos da aplicação
Vendors e Tools com padrões diferentes = Lock in

Surgiu a partir do opentracing e do opencensus.

## Componentes OpenTelemetry

### O que é o OpenTelemetry

É um framework de observabilidade para softwares cloud native
Conjunto de ferramentas, APIs e SDKs.
Instrumentação, geração, coleta e exportação de dados de telemetria

### Componentes principais

- Especificações
    - Dados
    - SDKs
    - APIs
- Collector
    - Pode coletar métricas, tracing e logs
    - Consegue funcionar como um componente independente
    - Agente ou Serviço
    - Pipeline (pega o dado, trata o dado, converte, exportar e enviar para outro sistema)
        - Recebimento
        - Processamento
        - Envio dados
    - Vendor-agnostic
- Libs
    - Vendor-agnostic
    - Tracing e Logs
    - Auto tracing
- Logs: draft

## Arquitetura Básica

- Sem Collector
- Com coletor no modo agente (side-car application)
- Coletor no modo server (collector server)


## Instrumentação

Como vocẽ lida com os dados de telemetria dentro da aplicação.

- Automática

- Manual

**spans** são blocos de código que estamos monitorando e gerando tracing.

## Exemplo Prático

Na pasta `dice` está o exemplo prático disponibilizado na documentação para Go do próprio site do OpenTelemetry.

```bash

go mod init dice


```

Na pasta `otel` estão os arquivos disponibilizados nas aulas.




### Adicioando as dependências

O comando à seguir instala os componentes OpenTelemetry SDK e instrumentação net/http.

```bash

go get "go.opentelemetry.io/otel" \
  "go.opentelemetry.io/otel/exporters/stdout/stdoutmetric" \
  "go.opentelemetry.io/otel/exporters/stdout/stdouttrace" \
  "go.opentelemetry.io/otel/exporters/stdout/stdoutlog" \
  "go.opentelemetry.io/otel/sdk/log" \
  "go.opentelemetry.io/otel/log/global" \
  "go.opentelemetry.io/otel/propagation" \
  "go.opentelemetry.io/otel/sdk/metric" \
  "go.opentelemetry.io/otel/sdk/resource" \
  "go.opentelemetry.io/otel/sdk/trace" \
  "go.opentelemetry.io/otel/semconv/v1.24.0" \
  "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"\
  "go.opentelemetry.io/contrib/bridges/otelslog"


```

### Inicialize o OpenTelemetry SDK

Primeiro, inicializaremos o OpenTelemetry SDK. Isto é necessário para qualquer aplicativo que exporte telemetria.

Para isso, foi criado o arquivo `otel.go`com o código de inicialização do OpenTelemetry SDK.

Se você estiver usando apenas tracing ou métricas, poderá omitir o código de inicialização TracerProvider ou MeterProvider correspondente.


### Adicionar instrumentação personalizada

As bibliotecas de instrumentação capturam a telemetria nas bordas dos seus sistemas, como solicitações HTTP de entrada e saída, mas não capturam o que está acontecendo no seu aplicativo. Para isso, você precisará escrever alguma instrumentação manual personalizada.


```bash

go mod tidy
export OTEL_RESOURCE_ATTRIBUTES="service.name=dice,service.version=0.1.0"
go run .


```

Abra http://localhost:8080/rolldice/Alice no seu navegador. Ao enviar uma solicitação ao servidor, você verá dois spans no rastreamento emitido para o console. O intervalo gerado pela biblioteca de instrumentação rastreia o tempo de vida de uma solicitação para a rota /rolldice/{player}. O intervalo denominado roll é criado manualmente e é filho do intervalo mencionado anteriormente.


## Referências

Cloud Native Foundation

https://www.cncf.io/

OpenTelemetry

https://opentelemetry.io/

Getting started for Developers

https://opentelemetry.io/docs/getting-started/dev/


OpenTelemetry-Go Contrib

https://github.com/open-telemetry/opentelemetry-go-contrib



Getting Started (OpenTelemetry)
https://opentelemetry.io/docs/languages/go/getting-started/