# Study-Buddy

Este repositório contém a aplicação "Study-Buddy" (frontend em TypeScript/React Native) e uma implementação anterior/alternativa localizada na pasta `deprecated` (backend em Go e páginas estáticas). Este README descreve a stack tecnológica, organização das pastas e observações sobre a versão `deprecated`.

**Visão Geral**

- **Objetivo:** Aplicação de apoio aos estudos com frontend moderno e uma implementação legada em Go.

**Stack**

- **Frontend:** React Native com TypeScript, bundler Metro, e configuração Babel.
- **Estilização:** Tailwind via NativeWind (integração Tailwind para React Native), com `tailwind.config.js` e `nativewind-env.d.ts`.
- **Ferramentas de dev:** `eslint` (configuração em `eslint.config.js`), `prettier` (`prettier.config.js`), `tsconfig.json` para TypeScript.
- **Empacotamento / bundling:** `metro.config.js` (padrão para apps React Native).

**Decisões de arquitetura**

- **Expo (managed workflow):** o projeto usa o ecossistema Expo para simplificar builds e deploys (veja `package.json` e dependências do Expo).
- **Roteamento baseado em arquivos:** adotado o `expo-router` para navegação file-based e convenções de rotas pelo sistema de pastas do `app/`.
- **TypeScript:** código escrito em TypeScript para tipagem estática e melhor DX; `tsconfig.json` já configurado.
- **Estilização com Tailwind/NativeWind:** escolha por `nativewind` para manter consistência de classes utilitárias similares à web.
- **Lint & Format:** `eslint` + `prettier` com `prettier-plugin-tailwindcss` para padronizar código e ordem de classes Tailwind.
- **Backend legado separado:** manter a implementação anterior em Go dentro de `deprecated/` como referência e possível backend local; não foi integrada como serviço remoto automático.
- **Persistência temporária (legado):** a versão `deprecated` utiliza JSON em disco para armazenamento simples — decisão consciente para prototipagem rápida.


**Backend (versão legada em `deprecated`)**

- Implementado em **Go** (módulo em `deprecated/go.mod`).
- Estrutura de backend inclui `main.go`, handlers (`handlers/`), middleware (`middleware/`), modelos (`models/`) e armazenamento local (`storage/`).
- Páginas estáticas para autenticação/fluxos simples em `deprecated/static/` (ex.: `login.html`, `register.html`).
- Armazenamento: arquivos JSON locais em `deprecated/storage` (ex.: `users.json`, `materials.json`) — não há banco de dados relacional nesta versão.

**Dados & Persistência**

- Versão legada usa JSON em disco (`deprecated/storage/*.json`).
- Frontend atual não inclui um backend remoto no repositório principal — a pasta `deprecated` representa a implementação anterior que pode servir como referência ou backend local.

**Principais dependências / requisitos**

- Node.js (para ferramentas JS/TS, Metro, bundling)
- Go (para compilar/rodar o backend presente em `deprecated`)
- Gerenciador de pacotes JS: `npm` / `yarn` (controle via `package.json`)

**Estrutura de pastas (resumo)**

- `app/` — código do frontend (arquivos TypeScript/React Native, ex.: [app/_layout.tsx](app/_layout.tsx#L1), [app/index.tsx](app/index.tsx#L1)).
- `components/` — componentes React reutilizáveis.
- `assets/` — imagens e recursos estáticos do frontend.
- `deprecated/` — implementação anterior com backend em Go e páginas estáticas (ver [deprecated/backend/main.go](deprecated/backend/main.go#L1) e [deprecated/static/index.html](deprecated/static/index.html#L1)).
- Arquivos de configuração: `package.json`, `tsconfig.json`, `tailwind.config.js`, `babel.config.js`, `metro.config.js`, `eslint.config.js`, `prettier.config.js`.

**Observações sobre `deprecated`**

- Trate `deprecated` como a mesma aplicação, porém construída com uma arquitetura diferente (monolito Go + páginas estáticas). Pode ser útil como referência para rotas, modelos de dados e lógica de negócio.

---
