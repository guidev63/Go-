package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeNomes()
	//exibeIntroducao()
	for {

		//exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			fmt.Println("Monitorando...")
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
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://cursos.alura.com.br/"
	sites[1] = "https://cursos.caelum.com.br/"

	fmt.Println(sites)

	site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "Foi carregado com Sucesso!")
	} else {
		fmt.Println("Site:", site, "Esta com Problemas. Status Code:", resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	fmt.Println(nomes)
}
