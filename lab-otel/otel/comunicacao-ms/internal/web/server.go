package web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"text/template"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

//go:embed template/*
var templateContent embed.FS

/*
Dica:
Essa tag acima adiciona o template ao binário, de forma que não vai precisar de uma pasta extra.
Nesse caso acima, ele pega a tag fo:embed e vai levar todo o conateúdo da pasta "template/*"
*/

// Struct para receber os dados para o webserver. Foi criado no arquivo main
type Webserver struct {
	TemplateData *TemplateData
}

// NewServer creates a new server instance
func NewServer(templateData *TemplateData) *Webserver {
	return &Webserver{
		TemplateData: templateData,
	}
}

// createServer creates a new server instance with go chi router
func (we *Webserver) CreateServer() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.Timeout(60 * time.Second))
	// promhttp. Usado para gerar as métricas automáticas do prometheus
	router.Handle("/metrics", promhttp.Handler())
	router.Get("/", we.HandleRequest)
	return router
}

type TemplateData struct {
	Title              string
	BackgroundColor    string
	ResponseTime       time.Duration
	ExternalCallMethod string
	ExternalCallURL    string
	Content            string
	RequestNameOTEL    string
	OTELTracer         trace.Tracer
}

func (h *Webserver) HandleRequest(w http.ResponseWriter, r *http.Request) {
	// Carregandos o header para gerar o request id para conseguir a rastreabilidade
	carrier := propagation.HeaderCarrier(r.Header)
	ctx := r.Context()
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	//Span adicional
	ctx, spanInicial := h.TemplateData.OTELTracer.Start(ctx, "SPAN_INICIAL "+h.TemplateData.RequestNameOTEL)
	time.Sleep(time.Second) // Adicionado mais um segundo para aparecer no trace
	spanInicial.End()

	// Criação de um span
	ctx, span := h.TemplateData.OTELTracer.Start(ctx, "Chama externa"+h.TemplateData.RequestNameOTEL)
	defer span.End()

	// Sleep apenas para gerar o tempo de processamento
	time.Sleep(time.Millisecond * h.TemplateData.ResponseTime)

	// Desvio de fluxo, caso não seja o último container de app.
	if h.TemplateData.ExternalCallURL != "" {
		var req *http.Request
		var err error

		// Se não for chamado GET ou POST, retorna internal server error
		if h.TemplateData.ExternalCallMethod == "GET" {
			req, err = http.NewRequestWithContext(ctx, "GET", h.TemplateData.ExternalCallURL, nil)
		} else if h.TemplateData.ExternalCallMethod == "POST" {
			req, err = http.NewRequestWithContext(ctx, "POST", h.TemplateData.ExternalCallURL, nil)
		} else {
			http.Error(w, "Invalid ExternalCallMethod", http.StatusInternalServerError)
			return
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Para facilitar o entendimento, esta request foi retirada do if acima
		// req, err = http.NewRequestWithContext(ctx, "GET", h.TemplateData.ExternalCallURL, nil)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// Injetando o header do request id. Necessário para realizar o tracker
		otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

		// Executando a request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Pega os dados que vieram da request
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Insere os dados no template.
		h.TemplateData.Content = string(bodyBytes)
	}

	tpl := template.Must(template.New("index.html").ParseFS(templateContent, "template/index.html"))
	err := tpl.Execute(w, h.TemplateData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}
