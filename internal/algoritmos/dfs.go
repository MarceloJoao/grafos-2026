package algoritmos

import (
	"fmt"
	"strings"

	"github.com/PauloFH/grafos-2026/internal/grafo"
)

type ResultadoDFS struct {
	Visitados   []string
	Predecessor map[string]string
	Entrada     map[string]int
	Saida       map[string]int
}

func DFS(g *grafo.Grafo, inicio string) ResultadoDFS {
	visitado := make(map[string]bool)
	predecessor := make(map[string]string)
	entrada := make(map[string]int)
	saida := make(map[string]int)
	visitados := []string{}
	tempo := 0

	var visitar func(u string)
	visitar = func(u string) {
		visitado[u] = true
		visitados = append(visitados, u)
		entrada[u] = tempo
		tempo++
		for _, w := range g.GetVizinhos(u) {
			if !visitado[w] {
				predecessor[w] = u
				visitar(w)
			}
		}
		saida[u] = tempo
		tempo++
	}

	visitar(inicio)

	return ResultadoDFS{
		Visitados:   visitados,
		Predecessor: predecessor,
		Entrada:     entrada,
		Saida:       saida,
	}
}

func FormataDFS(g *grafo.Grafo, inicio string) string {
	res := DFS(g, inicio)
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Vértice inicial: %s\n", inicio))
	sb.WriteString(fmt.Sprintf("Ordem de visita: %s\n\n", strings.Join(res.Visitados, " -> ")))
	sb.WriteString(fmt.Sprintf("%-10s %-8s %-8s %s\n", "Vértice", "Entrada", "Saída", "Predecessor"))
	sb.WriteString(strings.Repeat("-", 38) + "\n")
	for _, v := range res.Visitados {
		pred := res.Predecessor[v]
		if pred == "" {
			pred = "-"
		}
		sb.WriteString(fmt.Sprintf("%-10s %-8d %-8d %s\n", v, res.Entrada[v], res.Saida[v], pred))
	}

	return sb.String()
}
