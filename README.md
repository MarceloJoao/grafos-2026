# Trabalho de Grafos - 2026

## Estrutura

```
cmd/main.go                 → ponto de entrada
internal/grafo/grafo.go     → estrutura do grafo
internal/algoritmos/        → cada membro implementa aqui
internal/conversoes/        → cada membro implementa aqui
internal/leitor/leitor.go   → lê os arquivos da pasta inputs
internal/relatorio/relatorio.go → gera a saída padronizada
```

## Como rodar

```
go build -o projeto ./cmd/main.go
./projeto
```

## Como cada membro adiciona sua parte

### Passo 1 — Implemente a função

Abra o arquivo em `internal/conversoes/` ou `internal/algoritmos/` correspondente ao seu item.

Exemplo: `conversoes/lista_para_matriz.go`

```go
func ListaParaMatriz(g *grafo.Grafo) string {
 TODO: Fazer a conversao

}
```

### Passo 2 — Adicione ao relatório

No `cmd/main.go`, depois das linhas que já existem:

```go
matriz := conversoes.ListaParaMatriz(g)
r.Adiciona("MATRIZ_DE_ADJACENCIA", matriz)
```

### Passo 3 — Rode e verifique a saída

```
./projeto
cat outputs/GRAFO_1.txt
```

## Estrutura do Grafo

O grafo é simples, em lista de adjacência:

```go
type Grafo struct {
    Vertices    []string            // lista na ordem
    ListaAdj    map[string][]string // vértice -> vizinhos
    Direcionado bool                // true = digrafo
}
```

### Funções já prontas no grafo:

| Função | O que faz |
|---|---|
| `g.AdicionarVertice(id)` | cria vértice se não existir |
| `g.RemoverVertice(id)` | remove vértice e conexões |
| `g.AdicionarAresta(a, b)` | conecta dois vértices |
| `g.RemoverAresta(a, b)` | remove conexão |
| `g.NumVertices()` | total de vértices |
| `g.NumArestas()` | total de arestas |
| `g.SaoAdjacentes(a, b)` | verifica se são vizinhos |
| `g.GetVizinhos(id)` | retorna vizinhos de um vértice |

