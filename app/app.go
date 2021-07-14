package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar a aplicacao de linha de comando para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "check IP"
	app.Usage = "Busca IPs e nome de Servidores"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIP,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIP(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal((erro))
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal((erro))
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

}
