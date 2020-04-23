# Cadastro de Usuário usando Golang e MySQL

Aplicação Web que permite cadastrar usuário, esta aplicação utiliza os serviços  [GO-AUTH API](https://github.com/marceloagmelo/go-auth). Esta aplicação possue as seguintes funcionalidades.

- Listar Usuários
- Cadastrar Usuário
- Excluir Usuário

----

# Instalação

```
go get -v github.com/marceloagmelo/go-auth-web
```
```
cd go-auth-web
```

## Build da Aplicação

```
./image-build.sh
```

## Iniciar as Aplicações de Dependências
```
./dependecy-start.sh
```

## Preparar o MySQL

```
docker  exec -it mysqldb bash -c "mysql -u root -p"
```
- Criar a tabela
	> use goauthdb;
	
	> CREATE TABLE usuario (
id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
login VARCHAR(20), senha VARCHAR(100),
email VARCHAR(255), status INTEGER,
PRIMARY KEY (id)
);

## Iniciar a Aplicação
```
./start.sh
```
```
http://localhost:7070
```

## Finalizar a Aplicação
```
./stop.sh
```

## Finalizar a Todas as Aplicações
```
./stop-all.sh
```