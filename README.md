---

### Primeira API na linguagem GO

Este projeto consiste em uma API simples em Go capaz de realizar cálculos matemáticos básicos (adição, subtração, multiplicação e divisão).

#### Endpoints disponíveis:

- `/adicao`: Realiza a adição de um número à lista de números armazenados.
- `/subtracao`: Realiza a subtração de um número da lista de números armazenados.
- `/multiplicacao`: Realiza a multiplicação de todos os números armazenados por um número.
- `/divisao`: Realiza a divisão de todos os números armazenados por um número.
- `/total`: Retorna o total dos números armazenados.
- `/limpar`: Limpa a lista de números armazenados.

#### Funcionamento:

A API espera um corpo de requisição no formato JSON contendo um número para realizar as operações de adição, subtração, multiplicação ou divisão. Os números são armazenados até que sejam limpos através do endpoint `/limpar`.

### Uso:

Para utilizar a API, envie requisições HTTP para os endpoints listados acima, especificando o número desejado no corpo da requisição em formato JSON.

---
