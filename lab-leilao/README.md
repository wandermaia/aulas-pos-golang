# Laboratório de Concorrência com Golang - Leilão

Este módulo é referente as aulas do Lab Concorrência com Golang (sistema de leilão).


## Anotações


```bash

go mod init github.com/wandermaia/aulas-pos-golang/lab-leilao

go get github.com/joho/godotenv

```

Para as variáveis de ambiente, será utilizada a biblioteca godotenv

```bash

go get github.com/joho/godotenv

```

Para os logs será utilizada uma biblioteca da uber que gera os logs em formato json

```bash

go get go.uber.org/zap

```

Para o desenvolvimento do projeto, será utilizado o mogo-db (executando em docker)

```bash

go get go.mongodb.org/mongo-driver/mongo

```
Para criar o container do mongodb:


```bash

docker container run -d -p 27017:27017 --name auctionsDB mongo

```

## Referências