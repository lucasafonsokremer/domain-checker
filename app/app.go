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
	app.Version = "1.0.0"

	// como as flags são iguais para o segundo comando eu deixo aqui fora
	// caso contrario eu botaria dentro de commands
	flags := []cli.Flag{
		// devbook vai ser o valor padrão caso nada seja passado
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	// é um slice de cli
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			// chama a variavel flags declarada acima
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "nameservers",
			Usage:  "Busca nome dos name servers na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	// pega o valor da flag host e joga nessa var
	host := c.String("host")

	// pacote net busca os ips
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
