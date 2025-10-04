# Go Hexagonal Architecture

Este projeto é uma implementação de exemplo da Arquitetura Hexagonal em Go. Ele demonstra como estruturar uma aplicação para separar a lógica de negócios principal das preocupações com a infraestrutura, como bancos de dados, interfaces de usuário e APIs externas.

## Arquitetura

A estrutura do projeto segue os princípios da Arquitetura Hexagonal:

-   **`application`**: Contém a lógica de negócios principal e as interfaces (portas) que definem como o núcleo se comunica com o mundo exterior.
-   **`adapters`**: Contém as implementações concretas das portas (adaptadores). Isso inclui adaptadores para a interface de linha de comando (CLI), a API da web e a persistência do banco de dados.
-   **`cmd`**: Utiliza a biblioteca Cobra para configurar os comandos da CLI que servem como pontos de entrada para interagir com a aplicação.

## Pré-requisitos

-   Go 1.24 ou superior
-   SQLite3

## Começando

1.  **Clone o repositório:**
    ```sh
    git clone https://github.com/raffreitas/fc-hex-go.git
    cd fc-hex-go
    ```

2.  **Instale as dependências:**
    ```sh
    go mod tidy
    ```

3.  **Configure o banco de dados:**
    A aplicação usa um banco de dados SQLite. Crie o arquivo `db.sqlite` e a tabela `products` executando o seguinte comando SQL:
    ```sh
    sqlite3 db.sqlite 'CREATE TABLE products ("id" string, "name" string, "price" float, "status" string);'
    ```

## Gerando Mocks

O projeto utiliza `mockgen` para gerar mocks para os testes. Para gerar novamente os mocks, execute o seguinte comando a partir da raiz do projeto:

```sh
go generate ./...
```
ou
```sh
mockgen -destination=application/mocks/application.go -source=application/product.go application
```

## Uso

A aplicação pode ser executada tanto como uma CLI quanto como um servidor web.

### Interface de Linha de Comando (CLI)

A CLI permite gerenciar produtos.

-   **Criar um produto:**
    ```sh
    go run main.go cli -a create -n "Meu Produto" -p 29.99
    ```

-   **Obter um produto:**
    ```sh
    go run main.go cli -i "uuid-do-produto"
    ```

-   **Habilitar um produto:**
    ```sh
    go run main.go cli -a enable -i "uuid-do-produto"
    ```

-   **Desabilitar um produto:**
    ```sh
    go run main.go cli -a disable -i "uuid-do-produto"
    ```

### Servidor Web

Para iniciar o servidor web, execute:

```sh
go run main.go http
```

O servidor será iniciado em `http://localhost:8080`.

#### Endpoints da API

-   `POST /product`: Cria um novo produto.
-   `GET /product/{id}`: Obtém um produto pelo seu ID.
-   `PUT /product/{id}/enable`: Habilita um produto.
-   `PUT /product/{id}/disable`: Desabilita um produto.