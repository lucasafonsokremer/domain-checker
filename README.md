# domain-checker

domain-checker is a GO CLI tool to check domains

## Instalation

### Manual

Go to the [releases](https://github.com/lucasafonsokremer/domain-checker/releases) page and download the latest one for your platform. Just place the binary in your $PATH and you are good to go.

## Usage

```
domain-checker [command] [--flags]
```

- A few examples

```
$ domain-checker nameservers --host amazon.com.br

$ domain-checker ip --host amazon.com.br
```

- Get Help

```
domain-checker --help
```

- Output:

```
NAME:
   Aplicação de linha de comando - Busca IPs e nomes

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   lucasafonsokremer

COMMANDS:
   ip           Busca IPs de endereços na internet
   nameservers  Busca nome dos name servers na internet
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## ToDo

- Automated tests
- CLI auto-completion
- New functions like HTTP status check
