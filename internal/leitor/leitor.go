package leitor

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/PauloFH/grafos-2026/internal/grafo"
)

// LerArquivo lê um arquivo e devolve o grafo
func LerArquivo(caminho string) (*grafo.Grafo, error) {
	arquivo, err := os.Open(caminho)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()

	// Decide se é grafo ou digrafo pelo nome
	direcionado := strings.Contains(strings.ToUpper(filepath.Base(caminho)), "DIGRAFO")

	nome := filepath.Base(caminho)
	g := grafo.NovoGrafo(direcionado, nome)

	// Lê o arquivo
	scanner := bufio.NewScanner(arquivo)
	linhaNum := 0

	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text())
		linhaNum++

		if linha == "" {
			continue
		}

		// Primeira linha: número de vértices (ignorar)
		if linhaNum == 1 {
			continue
		}

		// Demais linhas: origem,destino
		partes := strings.Split(linha, ",")
		if len(partes) != 2 {
			continue
		}

		origem := strings.TrimSpace(partes[0])
		destino := strings.TrimSpace(partes[1])

		g.AdicionarAresta(origem, destino)
	}

	return g, nil
}

// LerDiretorio lê todos os .txt de uma pasta
func LerDiretorio(caminho string) (map[string]*grafo.Grafo, error) {
	resultado := make(map[string]*grafo.Grafo)

	entradas, err := os.ReadDir(caminho)
	if err != nil {
		return nil, err
	}

	for _, e := range entradas {
		if e.IsDir() {
			continue
		}
		if !strings.HasSuffix(strings.ToLower(e.Name()), ".txt") {
			continue
		}

		caminhoArquivo := filepath.Join(caminho, e.Name())
		g, err := LerArquivo(caminhoArquivo)
		if err != nil {
			fmt.Println("Erro ao ler", e.Name(), ":", err)
			continue
		}

		nome := strings.TrimSuffix(e.Name(), ".txt")
		resultado[nome] = g
	}

	return resultado, nil
}
