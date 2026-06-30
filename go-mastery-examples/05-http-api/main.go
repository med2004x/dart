package main

import (
    "encoding/json"
    "errors"
    "log/slog"
    "net/http"
    "os"
    "strings"
    "sync"
    "time"
)

type Product struct {
    ID         string `json:"id"`
    Name       string `json:"name"`
    PriceCents int    `json:"price_cents"`
}

type CreateProductRequest struct {
    Name       string `json:"name"`
    PriceCents int    `json:"price_cents"`
}

type ValidationError struct{ Message string }
func (e ValidationError) Error() string { return e.Message }

type Store struct {
    mu sync.RWMutex
    products map[string]Product
}

func NewStore() *Store { return &Store{products: map[string]Product{}} }

func (s *Store) Save(p Product) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.products[p.ID] = p
}

func (s *Store) List() []Product {
    s.mu.RLock()
    defer s.mu.RUnlock()
    out := make([]Product, 0, len(s.products))
    for _, p := range s.products {
        out = append(out, p)
    }
    return out
}

type Handler struct {
    store *Store
    logger *slog.Logger
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
    var req CreateProductRequest
    decoder := json.NewDecoder(r.Body)
    decoder.DisallowUnknownFields()
    if err := decoder.Decode(&req); err != nil {
        writeError(w, http.StatusBadRequest, "invalid JSON")
        return
    }
    if strings.TrimSpace(req.Name) == "" || req.PriceCents <= 0 {
        writeError(w, http.StatusBadRequest, ValidationError{Message: "invalid product"}.Error())
        return
    }
    p := Product{ID: time.Now().Format("20060102150405.000000000"), Name: strings.TrimSpace(req.Name), PriceCents: req.PriceCents}
    h.store.Save(p)
    writeJSON(w, http.StatusCreated, p)
}

func (h *Handler) ListProducts(w http.ResponseWriter, r *http.Request) {
    writeJSON(w, http.StatusOK, h.store.List())
}

func writeJSON(w http.ResponseWriter, status int, v any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    _ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, message string) {
    writeJSON(w, status, map[string]string{"error": message})
}

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
    h := &Handler{store: NewStore(), logger: logger}

    mux := http.NewServeMux()
    mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
    mux.HandleFunc("POST /products", h.CreateProduct)
    mux.HandleFunc("GET /products", h.ListProducts)

    server := &http.Server{
        Addr:              ":8080",
        Handler:           mux,
        ReadHeaderTimeout: 5 * time.Second,
        ReadTimeout:       10 * time.Second,
        WriteTimeout:      30 * time.Second,
        IdleTimeout:       60 * time.Second,
    }

    logger.Info("server starting", "addr", server.Addr)
    if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
        logger.Error("server failed", "error", err)
        os.Exit(1)
    }
}
