# goexpert-auction
Projeto do Laboratório "Concorrência com Golang - Leilão" da Pós Graduação GoExpert(FullCycle).



## O desafio
Adicionar uma nova funcionalidade ao projeto já existente para o leilão fechar automaticamente a partir de um tempo definido.

Toda rotina de criação do leilão e lances já está desenvolvida, entretanto, [o projeto clonado](https://github.com/devfullcycle/labs-auction-goexpert) necessita de melhoria: adicionar a rotina de fechamento automático a partir de um tempo.

Para essa tarefa, você utilizará o go routines e deverá se concentrar no processo de criação de leilão (auction). A validação do leilão (auction) estar fechado ou aberto na rotina de novos lançes (bid) já está implementado.



## Como rodar o projet
``` shell
## put the docker-compose containers up
make up 

## put the docker-compose containers down
make down

## make some request
make run
```

1. POST /auction
Rota: http://localhost:8080/auction
Método: POST
Descrição: Cria um novo leilão.
JSON de entrada (create_auction.json):

json

{
  "product_name": "Carro Antigo",
  "category": "Veículos",
  "description": "Um carro antigo em ótimo estado de conservação.",
  "condition": 1
}

2. GET /auction
Rota: http://localhost:8080/auction
Método: GET
Descrição: Lista todos os leilões, com filtros opcionais.
Parâmetros de query:

status (int): Status do leilão.

category (string): Categoria do leilão.

product_name (string): Nome do produto.

Exemplo de URL com query parameters:


http://localhost:8080/auction?status=1&category=Veículos&product_name=Carro



3. GET /auction/:auctionId
Rota: http://localhost:8080/auction/123e4567-e89b-12d3-a456-426614174000
Método: GET
Descrição: Busca um leilão pelo ID.
Parâmetro de URL:

4. GET /auction/winner/:auctionId
Rota: http://localhost:8080/auction/winner/123e4567-e89b-12d3-a456-426614174000
Método: GET
Descrição: Busca o lance vencedor de um leilão pelo ID do leilão.
Parâmetro de URL:

5. POST /bid
Rota: http://localhost:8080/bid
Método: POST
Descrição: Cria um novo lance.
JSON de entrada (create_bid.json):

json

{
  "user_id": "123e4567-e89b-12d3-a456-426614174001",
  "auction_id": "123e4567-e89b-12d3-a456-426614174000",
  "amount": 16000.50
}

6. GET /bid/:auctionId
Rota: http://localhost:8080/bid/123e4567-e89b-12d3-a456-426614174000
Método: GET
Descrição: Lista todos os lances de um leilão pelo ID do leilão.
Parâmetro de URL:

7. GET /user/:userId
Rota: http://localhost:8080/user/123e4567-e89b-12d3-a456-426614174001
Método: GET
Descrição: Busca um usuário pelo ID.

## Funcionalidades da Linguagem Utilizadas
- context
- net/http
- encoding/json
- testing
- testify



## Requisitos: implementação
- [ ] Uma função que irá calcular o tempo do leilão, baseado em parâmetros previamente definidos em variáveis de ambiente
- [ ] Uma nova go routine que validará a existência de um leilão (auction) vencido (que o tempo já se esgotou) e que deverá realizar o update, fechando o leilão (auction);
- [ ] Um teste para validar se o fechamento está acontecendo de forma automatizada;



## Requisitos: entrega
- [x] O código-fonte completo da implementação.
- [x] Documentação explicando como rodar o projeto em ambiente dev.
- [x] Utilize docker/docker-compose para podermos realizar os testes de sua aplicação.
