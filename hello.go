package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	registraLog("site-falso", false)

	exibeIntroducao()
	leSitesDoArquivo()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do Programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
		}
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func exibeIntroducao() {
	nome := "Guilherme"
	versao := 1.1

	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLindo int

	fmt.Scan(&comandoLindo)

	fmt.Println("O endereço da minha variável comando é", &comandoLindo)
	fmt.Println("O comando escolhido foi", comandoLindo)
	fmt.Println("")

	return comandoLindo
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")

	//  sites := []string{
	//	"https://random-status-code.herokuapp.com/",
	//	"https://cursos.alura.com.br/",
	//	"https://cursos.caelum.com.br/",
	//}

	sites := leSitesDoArquivo()

	for rodada := 0; rodada < monitoramentos; rodada++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")

	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um Erro", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}
func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.Txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + " " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(arquivo))
}
