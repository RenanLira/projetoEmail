# Projeto Email

Este projeto gerencia campanhas de e-mail, oferecendo rotas para criar, listar e manipular campanhas.

## Instalação

1. Certifique-se de ter o [Go](https://go.dev) instalado (versão 1.20+).
2. Clone o repositório:

   ```bash
   git clone https://github.com/seu-usuario/projetoEmail.git
   ```

3. Entrar no projeto:

    ```bash
    cd projetoEmail
    ```

4. Instalando e executando dependencias externas com `docker compose`
    > Keycloak e banco de dados Postgres

    ```bash
    docker compose up -d
    ```

5. Configurando `.env`

    1. crie um arquivo chamado `.env`
    2. utilize o `.env.exemple` como exemplo

6. Iniciando o projeto:

    ```bash
    go run cmd/api/main.go
    ```
