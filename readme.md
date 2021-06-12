# Descrição

Projeto em go com hot reload usando. A lib usada para o hot realod é [https://github.com/canthefason/go-watcher](https://github.com/canthefason/go-watcher). 

# Como usar

O docker file inclue o binario do projeto acima que quando está em uma pasta com um projeto em GO ele é capaz de fazer o hot reload a cada alteração nos arquivos.

Lembrando que para o projeto funciona deve ser iniciado o go modules.

```
go mod init ...
```

# Rodando

O volume já está configurado o que facilita o uso basta um docker-compose up