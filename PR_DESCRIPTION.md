# Pull Request: Add SQLite database layer with flexible interface

## Link para criar o PR

https://github.com/JulioPeixoto/gastei-quanto/pull/new/feature/sqlite-database-layer

## Título

```
feat: Add SQLite database layer with flexible interface
```

## Descrição

Esta PR adiciona uma camada de banco de dados flexível usando SQLite com uma interface que permite trocar facilmente para outros bancos SQL.

## Mudanças

- Criada interface `Database` em `pkg/database` para abstrair a implementação do banco de dados
- Implementado cliente SQLite com migrations automáticas
- Criados repositories SQL para `auth` e `expense` que usam a interface
- Adicionadas tabelas com foreign keys e índices para performance
- Migrations executadas automaticamente na inicialização
- Suporte para variáveis de ambiente (DB_DRIVER, DB_DSN)
- Atualizado README com documentação completa

## Estrutura

- Interface flexível permite adicionar PostgreSQL, MySQL, etc facilmente
- Mantidos os repositories em memória para testes
- Foreign keys e índices configurados corretamente
- Isolamento total por usuário

## Endpoints adicionados

Todos os endpoints de expenses agora persistem no banco:
- `POST /api/v1/expenses` - Criar despesa
- `GET /api/v1/expenses` - Listar despesas (com filtros)
- `GET /api/v1/expenses/stats` - Estatísticas de despesas
- `GET /api/v1/expenses/:id` - Buscar despesa específica
- `PUT /api/v1/expenses/:id` - Atualizar despesa
- `DELETE /api/v1/expenses/:id` - Deletar despesa
- `POST /api/v1/expenses/import` - Importar transações

## Como testar

1. Checkout da branch
   ```bash
   git checkout feature/sqlite-database-layer
   ```

2. Baixar dependências
   ```bash
   go mod download
   ```

3. Executar aplicação
   ```bash
   go run src/cmd/api/main.go
   ```

4. O banco SQLite será criado automaticamente em `gastei-quanto.db`

5. Testar endpoints via Swagger: http://localhost:8080/swagger/index.html

## Variáveis de Ambiente

```env
JWT_SECRET=your-secret-key-change-in-production
DB_DRIVER=sqlite
DB_DSN=./gastei-quanto.db
```

## Commits (6 total)

1. `feat: add database interface and SQLite implementation`
2. `feat: add SQL repository implementations for auth and expense`
3. `feat: integrate SQLite database layer in main application`
4. `chore: add SQLite dependency and update gitignore`
5. `docs: update README with database layer and new endpoints documentation`
6. `fix: remove duplicate function declaration in expense repository`

## Arquivos modificados

- `src/pkg/database/database.go` - Interface Database
- `src/pkg/database/sqlite.go` - Implementação SQLite
- `src/internal/auth/repository_sql.go` - Repository SQL para auth
- `src/internal/expense/repository_sql.go` - Repository SQL para expense
- `src/cmd/api/main.go` - Integração do banco de dados
- `README.md` - Documentação atualizada
- `.gitignore` - Ignorar arquivos .db
- `go.mod` / `go.sum` - Dependência sqlite3

## Próximos passos

Após aprovação desta PR, pode-se facilmente adicionar suporte para:
- PostgreSQL
- MySQL
- SQL Server

Basta criar um novo arquivo em `pkg/database/` implementando a interface `Database`.

