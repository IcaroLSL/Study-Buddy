
# Study Buddy

Study Buddy é uma aplicação web simples para gerenciar usuários e dados de estudo. Este repositório contém o backend em Go, alguns arquivos estáticos para a interface e um armazenamento simples em arquivos JSON.

## Estrutura principal

- `backend/` - código fonte do servidor em Go (API e rotas).
- `static/` - arquivos estáticos (HTML) servidos pelo backend.
- `storage/` - armazenamento local em JSON (usuários e dados).

## Requisitos

- Go (versão 1.24.2 ou superior recomendada, conforme `backend/go.mod`).

## Como executar (desenvolvimento)

Abra um terminal e execute o servidor com o seguinte comando:

```bash
cd backend && go run main.go
```

Depois de iniciado, o servidor fica disponível em: http://localhost:8080

Endpoints úteis:

- Páginas estáticas: `/` (index), `/login`, `/register`, `/forgot-password` (servos via `/static`).
- Autenticação pública: `/auth/login`, `/auth/register`.
- API protegida (requere token): `/api/data`, `/api/events/:id`.

## Observações

- Os dados persistem em arquivos JSON na pasta `storage/`.
- Se quiser criar um binário em vez de usar `go run`:

```bash
cd backend 
go build -o study-buddy main.go
./study-buddy
```

Se precisar de ajuda extra para rodar em outra porta ou com um sistema de produção, posso adicionar instruções e um Dockerfile.

