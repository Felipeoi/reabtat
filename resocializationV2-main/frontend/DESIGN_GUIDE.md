# Guia de Design - Sistema de Ressocialização

## 🎨 Paleta de Cores

### Verde Principal (Primary)
A paleta verde foi projetada para transmitir renovação, esperança e crescimento:

```css
primary-50:  #f0fdf4  /* Verde muito claro - fundos */
primary-100: #dcfce7  /* Verde suave - bordas e cards */
primary-200: #bbf7d0  /* Verde claro - hover states */
primary-300: #86efac  /* Verde médio-claro */
primary-400: #4ade80  /* Verde médio - foco */
primary-500: #22c55e  /* Verde principal - botões */
primary-600: #16a34a  /* Verde escuro - hover botões */
primary-700: #15803d  /* Verde mais escuro */
primary-800: #166534  /* Verde muito escuro - textos */
primary-900: #14532d  /* Verde profundo */
```

### Cores Complementares

- **Grays**: Para textos e elementos neutros
- **Red**: Para ações destrutivas (deletar, sair)
- **White**: Para fundos de cards e modais

## 🎭 Gradientes

### Gradiente Verde (gradient-green)
```css
background: linear-gradient(to right, #22c55e, #10b981);
```
Usado em: Botões primários, logo, navegação ativa

### Gradiente Verde Claro (gradient-green-light)
```css
background: linear-gradient(to bottom right, #f0fdf4, #dcfce7);
```
Usado em: Headers de cards e modais

### Gradiente de Fundo (body)
```css
background: linear-gradient(to bottom right, #f0fdf4, #dcfce7, #ccfbf1);
```
Fundo suave e agradável para toda a aplicação

## 💫 Sombras

### Sombra Verde (shadow-green)
```css
box-shadow: 0 4px 14px 0 rgba(34, 197, 94, 0.15);
```
Sombra sutil com toque verde para cards e botões

### Sombra Verde Grande (shadow-green-lg)
```css
box-shadow: 0 10px 40px 0 rgba(34, 197, 94, 0.2);
```
Sombra mais pronunciada para modais e elementos destacados

## 🧩 Componentes

### Button
**Variantes:**
- `primary`: Gradiente verde com sombra verde
- `secondary`: Branco com borda verde
- `danger`: Gradiente vermelho
- `ghost`: Transparente com hover verde claro

**Tamanhos:**
- `sm`: Pequeno (px-4 py-2)
- `md`: Médio (px-5 py-2.5)
- `lg`: Grande (px-8 py-3.5)

**Efeitos:**
- Hover: Scale 1.05
- Active: Scale 0.95
- Transição suave de cores

### Input
**Características:**
- Border 2px com hover verde
- Focus ring verde (primary-400)
- Padding confortável (px-4 py-3)
- Bordas arredondadas (rounded-xl)

### Card
**Estrutura:**
- Borda verde claro (border-primary-100)
- Header com gradiente verde claro
- Sombra verde
- Hover aumenta sombra

### Select
**Estilo:**
- Idêntico ao Input para consistência
- Border 2px verde no hover
- Padding confortável

### Modal
**Design:**
- Backdrop com blur
- Borda verde
- Header com gradiente verde claro
- Botão fechar com hover vermelho

## 📱 Layout

### Navegação
**Características:**
- Logo com gradiente verde (12x12)
- Altura 20 (h-20)
- Links com indicador verde quando ativos
- Botão sair com hover vermelho

### Espaçamento
- Padding principal: `py-8`
- Cards: `p-6`
- Forms: `space-y-5`
- Elementos: `space-x-2` ou `space-y-2`

## 🎯 Páginas

### Login / Signup
**Layout:**
- Logo verde grande (20x20) com hover scale
- Título grande (text-4xl)
- Card branco com borda verde
- Fundo com gradiente suave

### Dashboard (Usuários, Reeducandos, Correlações)
**Estrutura:**
- Container max-w-7xl
- Cards com sombra verde
- Tabelas com hover verde claro
- Botões de ação com cores apropriadas

### 404 Not Found
**Design:**
- Número 404 com gradiente verde
- Card centralizado
- Botão de retorno destacado

## ✨ Animações e Transições

### Transições Globais
```css
* {
  transition: colors 200ms;
}
```

### Hover Effects
- Botões: `scale-105`
- Logo: `scale-110`
- Links: Mudança de cor suave

### Active States
- Botões: `scale-95`
- Links: Cor mais escura

## 📏 Bordas e Raios

### Border Radius
- Pequeno: `rounded-lg` (8px)
- Médio: `rounded-xl` (12px)
- Grande: `rounded-2xl` (16px)
- Extra grande: `rounded-3xl` (24px)

### Border Width
- Padrão: `border` (1px)
- Destaque: `border-2` (2px)

## 🎨 Uso das Cores

### Textos
- Títulos principais: `text-gray-900`
- Títulos secundários: `text-primary-800`
- Corpo: `text-gray-600`
- Links: `text-primary-600` hover `text-primary-700`

### Fundos
- Principal: Gradiente verde suave
- Cards: `bg-white`
- Hover: `bg-primary-50`
- Ativo: `gradient-green`

### Bordas
- Padrão: `border-gray-200`
- Hover: `border-primary-200`
- Foco: `border-primary-400`
- Destaque: `border-primary-100`

## 🚀 Melhores Práticas

1. **Sempre use as classes customizadas** quando disponível:
   - `gradient-green` em vez de escrever o gradiente
   - `shadow-green` em vez de shadow-md

2. **Mantenha consistência**:
   - Todos inputs com mesmo estilo
   - Todos botões primários com gradiente verde
   - Todas bordas com border-2

3. **Use hover states**:
   - Sempre adicione `hover:` classes
   - Use `transform hover:scale-105` para interatividade

4. **Transições suaves**:
   - Adicione `transition-all` quando necessário
   - Duração padrão: 200ms

5. **Acessibilidade**:
   - Contraste adequado em todos os textos
   - Focus rings visíveis
   - Tamanhos de click adequados (min 44x44px)

## 🎨 Exemplos de Uso

### Botão Primário
```tsx
<Button variant="primary" size="lg">
  Criar Novo
</Button>
```

### Card com Título
```tsx
<Card title="Usuários">
  {/* Conteúdo */}
</Card>
```

### Input com Label
```tsx
<Input
  label="Nome"
  placeholder="Digite seu nome"
  value={name}
  onChange={handleChange}
/>
```

### Modal
```tsx
<Modal
  isOpen={isOpen}
  onClose={handleClose}
  title="Adicionar Usuário"
>
  {/* Formulário */}
</Modal>
```

## 🎯 Próximos Passos

Se quiser personalizar ainda mais:

1. **Adicionar modo escuro**: Criar variantes dark: no Tailwind
2. **Mais cores**: Adicionar paletas secundárias (azul, roxo)
3. **Animações complexas**: Usar Framer Motion
4. **Ícones**: Adicionar biblioteca como Lucide React
5. **Ilustrações**: Adicionar SVGs customizados

## 📱 Responsividade

Todos os componentes são responsivos por padrão:
- Mobile first
- Breakpoints: sm, md, lg, xl
- Padding e spacing adaptivos
- Grid e flex layouts responsivos
