package proxy

import (
	"io"
	"net/http"
)

type Proxy struct {
	client *http.Client
	target string
}

func NewProxy(target string, client *http.Client) *Proxy {
	return &Proxy{
		client: client,
		target: target,
	}
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(r.Method, p.target, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := p.client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "erro ao ler resposta", http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(resp.StatusCode)
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
