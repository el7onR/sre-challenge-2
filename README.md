# Exame - Bexs DevOps


## Introdução

Resultado do teste técnico da Bexs Banco para posição de SRE. Detalhes dos requisitos:

* Crie imagens Docker para ambas as aplicações. 
* Preencher este arquivo README.md com os detalhes, linha de raciocínio e dicas para os desenvolvedores que utilizarão sua solução.
* Considere que os desenvolvedores estão iniciando carreira e precisarão de mais detalhes de como executar sua solução.
* A Stack pode usar os recursos do próprio desenvolvedor(ex. VirtualBox, Docker, Docker-Compose) ou recursos de um provedor de cloud (Amazon Web Service ou Google Cloud)
* Não é necessário a criação de um pipeline. Considere que sua solução fará o bootstrap da Stack em questão.
* Não se preocupe em montar uma solução complexa. Dê preferência em montar uma solução simples que permita que o desenvolvedor realize melhorias.
* Apresente um desenho macro de arquitetura de sua solução.

## A Solução

O projeto consiste de 3 componentes, sendo:

### Frontend

* Linguagem: Python 3x
* Framework Web: Flask
* WSGI: Gunicorn
* Método de execução: Docker, Docker Compose ou local
* Dependências: Backend
* Porta de exposição: 8000/TCP
* Healthcheck URL: <http://localhost:8000/healthz>
* Base URL: <http://localhost:8000/>
* Metrics: <http://localhost:8001/metrics>

### Backend

* Linguagem: Go 1.15.x
* Framework Web: gin
* Framework ORM: gorm
* Método de execução: Docker, Docker Compose ou local
* Dependências: MySQL
* Porta de exposição: 8080/TCP
* Healthcheck URL: <http://localhost:8080/healthz>
* Base URL: <http://localhost:8080/>
* Metrics: <http://localhost:8080/metrics>
* Postman: Folder postman possui um export do Postman utilizado pra testar o backend

### MySQL

* Método de execução: Docker ou Docker Compose
* Porta de exposição: 3306/TCP

## Como Utilizar

Como executar o stack localmente
### Docker Compose - **Preferencial**

O stack possui o arquivo docker-compose.yml na raiz do projeto. O script possui mapeamento para os arquivos dev.Dockerfile de cada componente. Para executar o docker compose, é necessário possuir o docker e docker-compose instalado no computador, além de conexão à internet

Testes foram realizados em ambiente **Linux Ubuntu 20.04**, utilizando as versões **docker-compose 1.27.4** e **docker 19.03.13**.

**Importante**: Caso queria utilizar o Windows, será necessário ajustar os mapeamentos de volume das aplicações para que live reload funcione adequadamente

Comandos (devem ser executados na raiz do projeto)

* Para subir a stack em background: `docker-compose up --build -d` 
* Para subir a stack na console: `docker-compose up --build`
* Para verificar o status da stack: `docker-compose ps` **<- liberado para acesso quando todos os containers estverem healthy**
* Para parar a stack sem remover o conteúdo: `docker-compose stop`
* Para iniciar a stack: `docker-compose start`
* Para remover a stack e apagar os volumes e networks: `docker-compose down --remove -v`
* Para verificar os logs: `docker-compose logs -f`

**Importante**: Os logs em DEV estão configurados para o modo debug

### Local

Necessário possui o docker instalado (testado no **Linux Ubuntu 20.03** com **docker 19.03.13**). Além disso, é necessário possuir conexão à internet.

As varíaveis default da stack já estão apontando para os endereços e portas, caso queira ajustar, favor verificar as variáveis suportadas no arquivo docker-compose.yml

**Importante**: Comandos ajustados para executar no bash, para uso do Windows, será necessário ajustar os comandos.

Iniciar o MySQL

`docker run --name mysql -e MYSQL_ROOT_PASSWORD=root_pass -p 3306:3306 -e MYSQL_DATABASE=database --rm -d mysql`

Iniciar o Backend

Dentro da pasta backend, executar:

`go run cmd/main.go`

Iniciar o Frontend

Dentro da pasta frontend, executar:

`python3 -m venv .venv`
`source .venv/bin/activate`
`pip3 install -r requirements.txt`
`export prometheus_multiproc_dir=/tmp`
`gunicorn -c gunicorn_config.py wsgi:app --reload -w 4 --log-level debug`

## Bonus

Melhorias e ajustes adicionados ao código das aplicações de frontend e backend.

### Código Frontend

* Adoção do gunicorn
* Atualização de libs
* Separação de responsabilidades
* Novos comportamentos (listagem, alteração e remoção de usuários)
* Revisão do layout e estilo das páginas
* Controle de erros quando backend indisponível
* Adição de Dockerfile para produção e para Dev Local
* Suporte ao live reload no docker-compose para auxiliar no dev local

### Código Backend

* Mudança de mux para gin
* Adoção do gorm
* Mudança do sqllite para mysql
* Atualização de libs
* Adoção de go modules
* Separação de responsabilidades
* Novos comportamentos (alteração de usuário)
* Maior resiliência quando banco de dados está indisponível.
* Novos retornos HTTP em acordo ao Rest API
* Adição de Dockerfile para produção e para Dev Local
* Suporte ao live reload no docker-compose para auxiliar no dev local (go air)
* Ajuste no comportamento de criação do usuário para não considerar ID ao criar novo usuário e utilizar o ID do banco.
* Novas variáveis para conexão ao banco de dados.

## TODO

### Observability

* Instrumentalizar tracing no frontend e no backend
* Adicionar Prometheus no docker-compose e configurar para monitorar frontend e backend
* Adicionar Grafana no docker-compose e configurar com dashs do frontend e backend
* Configurar Jaeger no docker-compose.
* Adicionar stack do ELK para capturar os logs das aplicações

### Testes

* Adicionar testes unitários e integrados no frontend e no backend
* Adicionar testes de carga do frontend e backend

### Aplicação Backend

* Adicionar documentação do swagger para as apis
* Ajustar os paths para versionamento
* Adicionar validações no middleware
* Adicionar caching para consultas ao banco

### Aplicação Frontend

* Adicionar caching para consultas ao backend
* Novas validações nos formulários


### Geral

* Facilitar setup local com makefile
* Adicionar documentação de cleanup