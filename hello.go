package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Não Conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLindo int

	fmt.Scan(&comandoLindo)

	fmt.Println("O endereço da minha variável comando é", &comandoLindo)
	fmt.Println("O comando escolhido foi", comandoLindo)

	return comandoLindo
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://cursos.alura.com.br/", "https://cursos.caelum.com.br/"}

	//fmt.Println(sites)
	for i, site := range sites {
		fmt.Println("Testando Site", i, ":", site)
		testaSite(site)
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Foi carregado com Sucesso!")
	} else {
		fmt.Println("Site:", site, "Esta com Problemas. Status Code:", resp.StatusCode)
	}
}
