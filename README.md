Primeira API na linguagem GO



Foi solicitado por meu professor uma API capaz de realizar cálculos matemáticos simples (adição, subtração, divisão e multiplicação).

Em vista do que foi pedido, criei endpoints proporcionais ao que iriam realizar, que foram:\n
|  "/adicao" - Adição                |\n
|  "/subtracao" - Subtração          |\n
|  "/multiplicacao" - Multiplicação  |\n
|  "/divisao" - Divisão              |\n
Além disso incrementei endpoints para visualização desses endpoints que foi o "/home" e a mais, endpoints para ver o total e para excluir os números digitados em determinada sessão.
|  "/total" - Total                  |
|  "/limpar" - Delete                |
O programa solicita em um corpo com formatação JSON, um número, a cada envio, o número é armazenado na determinada função, até que o usuário encerre o programa ou delete os números informados!
