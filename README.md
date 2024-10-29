# email-service-notification

A **Email Service Notification API** é uma API desenvolvida em Golang com o objetivo de enviar notificações por e-mail e armazenar um registro dos e-mails enviados em um banco de dados MongoDB. Este projeto ainda está em desenvolvimento e possui funcionalidades básicas, com melhorias e ajustes em andamento.

## Funcionalidades

- Envio de e-mails via protocolo SMTP.
- Registro de cada e-mail enviado em uma coleção MongoDB.

## Endpoints

- `POST /send-email`: Envia um e-mail e registra o envio no MongoDB.
