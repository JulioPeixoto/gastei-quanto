# Pull Request: Feature - Import and Save CSV with Auto-categorization

## Resumo

Esta PR adiciona a funcionalidade de importar arquivos CSV de transações com categorização automática e salvamento direto no banco de dados, integrando os módulos `parser`, `analysis` e `expense`.

## Mudanças Principais

### 1. Endpoint Modificado: `/api/v1/parser/upload/csv`
- Agora processa E salva arquivos CSV de transações automaticamente
- Categoriza automaticamente baseado em palavras-chave
- Salva diretamente no banco de dados vinculado ao usuário
- Retorna estatísticas do processamento (transações processadas e salvas)

### 2. Serviço de Integração (`integration_service.go`)
- Coordena parser, análise e expense
- Implementa algoritmo de categorização automática
- Validações robustas
- Logs detalhados

### 3. Categorização Automática
Identifica automaticamente as categorias baseado em palavras-chave:
- **Transporte**: uber, 99, taxi, ride, dl*, pg *
- **Alimentação**: ifood, restaurante, padaria, pizza, lanche
- **Compras**: amazon, mercado, loja, mercadolivre
- **Assinaturas**: spotify, netflix, prime
- **Taxas**: iof
- **Crédito**: estorno, pagamento recebido

### 4. Melhorias na Documentação
- Atualização do README com nova feature
- Documentação detalhada em `docs/IMPORT_FEATURE.md`
- Swagger atualizado com novo endpoint
- Arquivo CSV de exemplo em `examples/`

## Arquivos Modificados

### Novos Arquivos
- `src/internal/parser/integration_service.go` - Serviço de integração
- `docs/IMPORT_FEATURE.md` - Documentação da feature
- `examples/sample_transactions.csv` - Exemplo de CSV

### Arquivos Alterados
- `src/internal/parser/handler.go` - Novo handler ImportAndSaveCSV
- `src/internal/parser/routes.go` - Nova rota
- `src/internal/parser/model.go` - Novo modelo ImportAndSaveResponse
- `src/cmd/api/main.go` - Integração dos serviços
- `README.md` - Documentação atualizada
- `docs/` - Swagger regenerado

## Como Testar

### 1. Registrar/Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

### 2. Upload do CSV
```bash
curl -X POST http://localhost:8080/api/v1/parser/upload/csv \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -F "file=@examples/sample_transactions.csv"
```

### 3. Verificar Despesas Salvas
```bash
curl -X GET http://localhost:8080/api/v1/expenses \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Benefícios

1. **Experiência do Usuário**: Upload e salvamento em um único passo no endpoint existente
2. **Categorização Inteligente**: Reduz trabalho manual de categorização
3. **Validações Robustas**: Tratamento de erros em todas as etapas
4. **Logs Detalhados**: Facilita debugging e monitoramento
5. **Documentação Completa**: README, Swagger e documentação específica
6. **Backward Compatible**: Usa o endpoint existente, sem quebrar integração

## Testes Realizados

- ✅ Upload de CSV válido com categorização automática
- ✅ Validação de arquivo CSV inválido
- ✅ Validação de autenticação (JWT)
- ✅ Salvamento correto no banco de dados
- ✅ Logs informativos durante processamento
- ✅ Documentação Swagger atualizada

## Commits (10 total)

1. feat: adicionar suporte a user_id no parser handler e modelo de resposta para import
2. feat: criar servico de integracao e endpoint import-and-save para processar e salvar CSV automaticamente
3. feat: integrar servico de parser com expense e analysis no main
4. feat: adicionar validacoes e tratamento de erros robusto no integration service
5. feat: melhorar algoritmo de categorizacao automatica com mais palavras-chave
6. feat: adicionar logs informativos detalhados durante processamento e categorizacao
7. docs: atualizar documentacao swagger com novo endpoint import-and-save
8. docs: adicionar documentacao detalhada da feature de import and save
9. chore: adicionar arquivo CSV de exemplo para testes
10. docs: atualizar README com informacoes sobre endpoint import-and-save e auto-categorizacao

## Próximos Passos (Sugestões)

- [ ] Adicionar testes unitários para integration_service
- [ ] Implementar detecção de duplicatas
- [ ] Adicionar suporte para mais formatos de CSV
- [ ] Melhorar algoritmo de categorização com ML
- [ ] Adicionar validação de transações antes de salvar

## Notas

- Feature totalmente backwards compatible
- Endpoint antigo `/parser/upload/csv` continua funcionando
- Não há breaking changes

