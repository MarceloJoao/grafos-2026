package conversoes

import "github.com/PauloFH/grafos-2026/internal/grafo"

// MatrizIncidencia retorna a matriz de incidência (n×m) e a lista de arestas.
// Não-direcionado: célula = 1 se vértice é extremo da aresta.
// Direcionado: +1 para origem, -1 para destino.
func MatrizIncidencia(g *grafo.Grafo) ([][]int, [][2]string) {
	type par = [2]string

	// Enumera arestas
	visitados := make(map[par]bool)
	var arestas [][2]string
	for _, v := range g.Vertices {
		for _, viz := range g.ListaAdj[v] {
			var chave par
			if !g.Direcionado && viz < v {
				chave = par{viz, v}
			} else {
				chave = par{v, viz}
			}
			if visitados[chave] {
				continue
			}
			visitados[chave] = true
			arestas = append(arestas, par{v, viz})
		}
	}

	n := len(g.Vertices)
	m := len(arestas)

	// Índice de vértices
	idx := make(map[string]int, n)
	for i, v := range g.Vertices {
		idx[v] = i
	}

	// Monta matriz n×m
	matriz := make([][]int, n)
	for i := range matriz {
		matriz[i] = make([]int, m)
	}

	for j, aresta := range arestas {
		origem, destino := aresta[0], aresta[1]
		i := idx[origem]
		k := idx[destino]
		if g.Direcionado {
			matriz[i][j] = +1
			matriz[k][j] = -1
		} else {
			matriz[i][j] = 1
			matriz[k][j] = 1
		}
	}

	return matriz, arestas
}
