# Contributing

Obrigado por contribuir com o nosso projeto! Siga as diretrizes abaixo para garantir um fluxo de colaboração claro e eficiente.

---

## 1. Código de Conduta

- Respeite todos os participantes.
- Mantenha um ambiente amigável e profissional.
- Evite linguagem ofensiva ou discriminatória.

Para mais detalhes, consulte o [CODE_OF_CONDUCT.md](./CODE_OF_CONDUCT.md) (se disponível).

---

## 2. Como Reportar Bugs

1. Verifique se já existe uma issue similar.
2. Abra uma nova issue usando o template de bug:
   - **Título**: `bug: [seção] breve descrição`
   - **Descrição**: passos para reproduzir, resultado atual, resultado esperado.
   - **Ambiente**: versão do Node/Golang/Terraform, sistema operacional.

---

## 3. Como Propor Novas Funcionalidades

1. Abra uma issue com o template de feature:
   - **Título**: `feat: [seção] breve descrição`
   - Descreva o caso de uso, benefício e critérios de aceitação.

---

## 4. Fluxo de Trabalho

### 4.1 Branches

- **Base**: `main`
- **Naming**:
  - `feature/<descrição>` para novas funcionalidades
  - `fix/<descrição>` para correções de bug
  - `chore/<descrição>` para tarefas gerais
  - `hotfix/<descrição>` para correções críticas em produção

### 4.2 Commits

Use **Conventional Commits**:

```
<type>(<escopo>): <descrição curta>

[corpo opcional]

[rodapé opcional]
```

- **Tipos**: `feat`, `fix`, `chore`, `docs`, `refactor`, `test`, `ci`
- **Exemplo**: `feat(server): adicionar endpoint /health`

### 4.3 Pull Requests

- Abra PR apontando para `main`.
- Preencha o template de PR:
  - Objetivo da mudança
  - Issues relacionadas (`Closes #123`)
  - Detalhes de implementação
  - Passos para testar localmente
- Aguarde revisão de pelo menos um maintainer.
- Após aprovação e CI verde, faça merge usando "Squash and merge".

---

## 5. Estilo de Código e Linting

### 5.1 PWA (React + pnpm)

- Utilize **ESLint** e **Prettier**.
- Regras definidas em `/pwa/.eslintrc.js` e `/pwa/.prettierrc`.
- Scripts úteis:
  - `pnpm install`
  - `pnpm lint`
  - `pnpm test`

### 5.2 Server (Golang)

- Formate com `gofmt` e organize imports com `goimports`.
- Estruture pacotes seguindo convenções de Go.
- Testes: `go test ./... -cover`

### 5.3 Infra (Terraform)

- Formate com `terraform fmt`.
- Valide sintaxe com `terraform validate`.
- Não comite arquivos de estado (tfstate).

---

## 6. Testes e Cobertura

- **PWA**: `pnpm test`, relatórios em `coverage/`.
- **Server**: `go test ./... -coverprofile=coverage.out`.
- **Infra**: `terraform validate` + `terraform plan` em CI.

---

## 7. Atualização de Documentação e Diagramas

- Diagramas C4 em PlantUML em `docs/architecture/*.puml`.
- Gere PNGs com `make docs` (via GitHub Actions).
- Commit dos PNGs atualizados é obrigatório (veja configuração CI para `EndBug/add-and-commit`).

---

## 8. Integração Contínua (CI)

- Pipeline executa:
  1. Lint e testes da PWA
  2. Testes e cobertura do servidor
  3. Validação e formatação do Terraform
  4. Geração de documentação (`make docs`)
- Merges só são permitidos com todos os checks verdes.

---

## 9. LGPD e Privacidade

- Nunca exponha dados sensíveis ou tokens reais.
- Use variáveis de ambiente para credenciais.
- Remova informações pessoais antes de commitar logs ou amostras.

---

## 10. Agradecimentos

Obrigado por ajudar a tornar este projeto melhor! 🎉
