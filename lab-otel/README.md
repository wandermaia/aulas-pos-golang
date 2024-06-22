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

### O que oe o OpenTelemetry
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







## Referências

Cloud Native Foundation

https://www.cncf.io/