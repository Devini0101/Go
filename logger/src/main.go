package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const amountMonitorings = 5
const delay = 2

func main() {

	greetingIdentifier()

	for {

		action := readAction()

		switch action {
		case 1:
			initializeMonitoring()
		case 2:
			fmt.Println("EXIBINDO LOGS...")
			showLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0) //informs the OS that the exit was sucessful
		default:
			fmt.Println("Desconheço essa operação")
			os.Exit(-1)
		}
	}

}

func greetingIdentifier() {
	fmt.Println("Olá Sr. Vinicius")
}

func readAction() int {
	fmt.Println("============")
	fmt.Println("1 - INICIAR MONITORAMENTO")
	fmt.Println("2 - EXIBIR LOGS")
	fmt.Println("0 - SAIR DO PROGRAMA")
	fmt.Println("============")

	var action int
	fmt.Scan(&action)
	return action
}

func initializeMonitoring() {
	fmt.Println("Monitorando...")
	sites := readSitesFromFile()

	for i := 0; i < amountMonitorings; i++ {
		for _, site := range sites {
			fmt.Println(site)
			testSite(site)
		}
		//timeout after test
		time.Sleep(delay * time.Minute) //1 vezes unidade de medida de minutos (de um em um minuto)
		fmt.Println("")
	}
}

func testSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro ao acessar o site:", err)
		return
	}
	statusCode := resp.StatusCode

	if statusCode == 200 {
		fmt.Println("Site está no ar, Ok. status: ", statusCode)
		registerLog(site, true)
	} else {
		fmt.Println("Site com status: ", statusCode, "!!")
		registerLog(site, false)
	}
}

func readSitesFromFile() []string {
	var sites []string
	archive, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return sites
	}

	leitor := bufio.NewReader(archive)

	for {
		line, err := leitor.ReadString('\n') //lê a linha até a quebra de linhas
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line) //remove espaços em branco
		sites = append(sites, line)
	}
	archive.Close()
	return sites
}

func registerLog(site string, status bool) {
	archive, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log:", err)

		return
	}

	archive.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - " + strconv.FormatBool(status) + "\n")
	archive.Close()
}

func showLogs() {
	archive, err := os.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro ao abrir o arquivo de logs: ", err)
	}

	fmt.Println(string(archive)) //string already can convert a array into a string
}
