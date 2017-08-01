package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tagsOrigem := readFile("readOrigin")
	tagsTrad := readFile("readTrad")
	fmt.Println(diff(tagsOrigem, tagsTrad))
}

func readFile(fileName string) map[string]string {
	tags := map[string]string{}
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// encontrar tag
		linha := scanner.Text()
		inicioDaTag := strings.Index(linha, "<")
		fimDaTag := strings.Index(linha, ">")
		tag := linha[inicioDaTag+1 : fimDaTag]
		// adicionar somente a tag, com conteúdo vazio
		tags[tag] = ""
	}
	return tags
}

func diff(origem, traducao map[string]string) []string {
	faltando := []string{}
	for chave := range origem {
		if _, ok := traducao[chave]; !ok {
			faltando = append(faltando, chave)
		}
	}
	return faltando
}
