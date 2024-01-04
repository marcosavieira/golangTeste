Teste Técnico Backend em Golang para Programador Júnior

Este teste técnico tem como objetivo avaliar suas habilidades básicas em programação em Golang, focando em conceitos fundamentais da linguagem e práticas de desenvolvimento. Certifique-se de seguir as melhores práticas de codificação, como modularização, tratamento de erros e legibilidade do código.

1. FizzBuzz em Golang (Pontuação: 20 pontos)

Escreva um programa em Golang que imprima os números de 1 a 100. No entanto, para múltiplos de 3, imprima "Fizz" em vez do número, e para múltiplos de 5, imprima "Buzz". Para números que são múltiplos de ambos (3 e 5), imprima "FizzBuzz".

2. Manipulação de Strings (Pontuação: 30 pontos)

Crie uma função em Golang que recebe uma string como entrada e retorna o número de palavras únicas na string. Considere que as palavras são separadas por espaços. Por exemplo, para a entrada "hello world world", a função deve retornar 2.

3. API Simples em Golang (Pontuação: 50 pontos)

Crie uma API simples em Golang que realiza operações básicas de CRUD (Create, Read, Update, Delete) em uma entidade chamada "Produto". Cada produto deve ter um ID único, nome e preço. Utilize um banco de dados em memória para armazenar os produtos.

A API deve fornecer os seguintes endpoints:

    POST /produto: Cria um novo produto.
    GET /produto/{id}: Retorna os detalhes de um produto específico.
    GET /produtos: Retorna a lista de todos os produtos.
    PUT /produto/{id}: Atualiza os detalhes de um produto existente.
    DELETE /produto/{id}: Exclui um produto específico.

Instruções Gerais:

    Utilize a biblioteca padrão do Golang para construir a API.
    Forneça uma breve documentação sobre como executar e testar a API.
    Certifique-se de lidar adequadamente com erros e casos de borda.
    Siga as melhores práticas de codificação em Golang.

Avaliação:

    20 pontos: Implementação correta e eficiente do FizzBuzz.
    30 pontos: Implementação correta e eficiente da função CountUniqueWords.
    50 pontos: Implementação correta e eficiente da API, seguindo as especificações fornecidas.

Observações:

    Você pode usar qualquer biblioteca externa que julgar necessário.
    Priorize a simplicidade e clareza do código.