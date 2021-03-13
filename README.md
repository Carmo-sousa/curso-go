# Curso de Go

Repositório com todos os códigos do curso de Golang da Cod3r.

## Índice
- [Fundamentos](01-fundamentos)
- [Controles](02-controles)
- [Array, Slice e Map](03-arrayslicemap)
- [Funções](04-funcoes)
- [Tipos](05-tipos)
- [Pacotes](06-pacotes)
- [Concorrência](07-concorrencia)
- [Anotações](#annotations)

## <a name=annotations><a/> Anotações

**Concorrência vs Paralelismo**

Go é uma linguagem concorrente.

- Paralelismo - executar código simultaneamente em processadores físicos diferentes.
- Concorrência - intercalar (administrar) vários processos ao mesmo tempo e isso pode ocorrer em um único processador físico.
- Concorrência - é a forma de estruturar o seu programa, é a composição de processos (tipicamente funções) que executam de forma independente.

Concorrência viabiliza paralelismo.

- É possível que a concorrência seja melhor que o paralelismo!
- Paralelismo exige muito mais do *SO* e do *hardware*.
