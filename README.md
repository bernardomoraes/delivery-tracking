# Delivery Tracking com Kafka e Golang

## Descrição

Este projeto tem como objetivo criar um sistema de rastreamento de entregas, utilizando Apache Kafka e Golang para que seja possível compreender o ciclo de vida de mensagens em tópicos kafkas. É um projeto desenvolvido com intuito de estudo e aprendizado.

## Configurar /etc/hosts

A comunicação entre as aplicações se dá de forma direta através da rede da máquina.
Para isto é necessário configurar um endereços que todos os containers Docker consigam acessar.

Acrescente no seu /etc/hosts (para Windows o caminho é C:\Windows\system32\drivers\etc\hosts):
```
127.0.0.1 host.docker.internal
```
Em todos os sistemas operacionais é necessário abrir o programa para editar o *hosts* como Administrator da máquina ou root.

## Rodar a aplicação


Dentro da pasta raíz, execute os passos abaixo para rodar a aplicação:

1. Para subir os producers e consumers do Kafka, execute os comandos abaixo:
    - `docker-compose --project-directory ./.docker/kafka up -d`
    - `docker exec -it <Kafka-Container-Name> bash`
      - Isso irá abrir o bash do container do Kafka para que possamos executar os comandos do Kafka.
    - Dentro do terminal aberto execute: 
        - `kafka-console-producer --bootstrap-server localhost:9092 --topic route.new-direction`
        - Esse comando irá abrir o producer do Kafka, onde podemos enviar mensagens para o Kafka.
    - Abra outro terminal e execute:
        - `docker exec -it <Kafka-Container-Name> bash`
        - `kafka-console-consumer --bootstrap-server localhost:9092 --topic route.new-position`
        - Esse comando irá abrir o consumer do Kafka, onde podemos receber as mensagens enviadas para o tópico especificado.
2. Agora em outro terminal, suba o container da aplicação:
    - `docker-compose up -d`
    - `docker-compose exec app bash`
    - Dentro do container, execute o comando abaixo para rodar a aplicação:
        - `go run main.go`
3. No terminal do producer, envie uma mensagem para o Kafka:
    - `{"clientId":"1","routeId":"1"}`
    - Para enviar outra mensagem, basta enviar outra mensagem no terminal do producer.

## Conclusão

- A mensagem será enviada para o topico ***route.new-direction***, a aplição está consumindo esse tópico. 
- Dessa forma a mensagem será recebida pela aplicação, ela fará a lógica de negócio e depois ira entregar mensagens simulando um rastreamento de entrega enviando mensagens com a posição em tempo real para o tópico ***route.new-position***
- Esse tópico por sua vez está sendo consumido pelo terminal consumer que foi criado.

> Dessa forma, o objetivo do repositório é atingido, sendo possível visualizar todo o fluxo de entrega e recebimento de mensagens no Kafka.