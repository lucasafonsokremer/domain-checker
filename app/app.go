package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Exporta função da cli globalmente
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e nomes"
	app.Author = "lucasafonsokremer"
	app.Version = "0.0.1"

	// Criação das flags
	flags := []cli.Flag{
		// google.com será o valor padrão caso nada seja passado
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			// chama a variavel flags declarada acima
			Flags:  flags,
			// chama a função buscaIps declarada abaixo
			Action: buscarIps,
		},
		{
			Name:   "nameservers",
			Usage:  "Busca nome dos name servers na internet",
			// chama a variavel flags declarada acima
			Flags:  flags,
			// chama a função buscarServidores declarada abaixo
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	nameservers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range nameservers {
		fmt.Println(servidor.Host)
	}

}
