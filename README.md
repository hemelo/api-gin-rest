# Simples API

API desenvolvida para me introduzir no Framework Web Gin feito em golang

**Dependências:** 
```
- Go Compiler
- Docker (opcional)
- Postgree
```    

## Rodando ambiente com Docker

Acesse o diretório em que o repositório foi clonado através do terminal e
execute os comandos:
 - `docker-compose build` para criar containers 
 - `docker-compose up` para criar banco de dados e inicializar o servidor

### Iniciando e finalizando containers
Para inicializar execute o comando `docker-compose start` e
para finalizar `docker-compose stop`

### Observação

Será necessário editar o arquivo `.env.example` para `.env` e inserir as credênciais
e criar o servidor no pgAdmin4 com o IP de conexão do Docker caso esteja utilizando, que é `host.docker.internal`.