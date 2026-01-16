# Feature: Import and Save CSV Transactions

## Objetivo

Esta feature permite que usuários façam upload de arquivos CSV contendo transações financeiras, que são automaticamente processadas, categorizadas e salvas no sistema.

## Como funciona

### Fluxo de Processamento

1. **Upload**: O usuário envia um arquivo CSV via endpoint `/api/v1/parser/import-and-save`
2. **Parsing**: O sistema lê e valida o arquivo CSV
3. **Análise**: As transações são analisadas para identificar padrões
4. **Categorização Automática**: Transações sem categoria recebem uma categoria sugerida baseada em palavras-chave
5. **Salvamento**: As transações são salvas no banco de dados vinculadas ao usuário autenticado

### Categorização Automática

O sistema utiliza palavras-chave para sugerir categorias automaticamente:

- **Transporte**: uber, 99, taxi, ride, dl*, pg *, estacionamento
- **Alimentação**: ifood, restaurante, padaria, pizza, lanche, açai, food, bar, cafe, tempero
- **Compras**: amazon, mercado, loja, mercadolivre
- **Assinaturas**: spotify, netflix, prime, dm *
- **Taxas**: iof
- **Crédito**: estorno, crédito de, pagamento recebido
- **Outros**: qualquer transação não categorizada

### Endpoints

#### 1. Upload CSV (apenas parsing)
```
POST /api/v1/parser/upload/csv
```
Processa o CSV e retorna as transações sem salvar.

#### 2. Import and Save (novo)
```
POST /api/v1/parser/import-and-save
```
Processa, categoriza e salva automaticamente as transações.

**Response:**
```json
{
  "message": "CSV processado e salvo com sucesso",
  "processed": 62,
  "saved": 62,
  "transactions": [...]
}
```

## Formato CSV Esperado

O CSV deve conter as seguintes colunas (case-insensitive):
- **date** (obrigatório): Data da transação (formatos aceitos: YYYY-MM-DD, DD/MM/YYYY, MM/DD/YYYY)
- **amount** (obrigatório): Valor da transação
- **description** (opcional): Descrição da transação
- **category** (opcional): Categoria (se vazia, será sugerida automaticamente)

### Exemplo

```csv
date,description,amount,category
2025-09-01,Dl*99 Ride,14.6,
2025-08-30,Amazon - Parcela 1/2,24.48,
2025-08-17,Amazonprimebr,19.9,Assinaturas
```

## Benefícios

- Importação em lote de transações
- Categorização automática inteligente
- Validações robustas
- Logs detalhados para debugging
- Integração completa com expense e analysis modules

## Arquitetura

### Componentes

1. **Parser Service**: Processa o arquivo CSV
2. **Integration Service**: Coordena parser, analysis e expense
3. **Analysis Service**: Analisa padrões nas transações
4. **Expense Service**: Persiste as transações no banco

### Fluxo de Dados

```
CSV File → Parser → Integration Service → Analysis → Expense Repository → Database
```

## Segurança

- Requer autenticação JWT (Bearer token)
- Validação de user_id
- Validação de formato de arquivo
- Tratamento de erros robusto

