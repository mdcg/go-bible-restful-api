# GO BIBLE RESTFUL API

[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/mdcg/go-bible-restful-api/blob/master/LICENSE)

## Introdução

Olá! Esse é o meu primeiro desenvolvimento de API RESTful usando a linguagem Go. Eu escolhi fazer algo relacionado à minha religião, que é o cristianismo, e antes de mais nada, agradeço a Cristo pela oportunidade e pela benção de poder contribuir um pouco mais na sua obra. 

Estou utilizando o banco de dados disponibilizado pelo [Thiago Bodruk](https://github.com/thiagobodruk), que você pode ter acesso ao repositório clicando [aqui](https://github.com/thiagobodruk/biblia). Desde já, eu o agradeço pelo seu trabalho em disponibilizar a biblía em diversos formatos.

Essa API RESTful é bem simples, dado que meus conhecimentos na linguagem Go ainda são bem superficiais, mas mesmo assim, saiba que dei o meu melhor para escrever da maneira mais otimizada e limpa possível, para que qualquer pessoa que deseje utilizá-la, possa ter o mínimo de esforço possível para estendê-la.

Muito obrigado pela sua atenção, e que Deus lhe abençoe.

## Primeiros passos

### Configurando as variáveis de ambiente

Antes de mais nada, precisamos configurar nossas variáveis de ambiente. Existe um arquivo dentro do diretório `config` chamado `.env-example` . No seu terminal, dentro do diretório `config` , execute o seguinte comando:

``` 
$ cp .env-example .env
```

Basicamente, o comando acima irá copiar o conteúdo de `.env-example` para `.env` . Dentro do arquivo `.env` você verá que existem várias variáveis de ambiente, no qual duas já estão preenchidas, que são `MYSQL_DATABASE` e `MYSQL_HOST` . Sugiro que você não as mude, do contrário, possa ser que o sistema não venha funcionar de maneira apropriada, no mais, fique a vontade e preencha as que não tenha nenhum valor já atribuido da maneira que você bem entender.

### Iniciando e restaurando o banco de dados

Agora que já temos as nossas variáveis de ambiente configuradas, precisamos iniciar o nosso banco de dados. Estamos utilizando o **Docker** e o **Docker Compose** para facilitar a execução de todo o sistema, por isso, peço que você os tenha instalados em sua máquina antes de executar os comandos a seguir.

Para inicializar o nosso banco de dados, iremos executar o seguinte comando:

``` 
$ docker-compose up -d bible_db
```

Esse comando pode demorar um pouco até ser concluído e deixar o servidor MySQL pronto para uso, por isso, seja paciente.

A flag `-d` deixará o container rodando em "backgroud", nesse caso, não deixará o seu terminal "bloqueado". Caso você deseje visualizar os logs do servidor MySQL, basta omitir a flag `-d` . 

Uma vez que o container esteja rodando, você precisará restaurar o banco de dados com as versões de traduções disponíveis. Existe um script chamado `restore_db.sh` para facilitar que você restaure o banco de dados sem muita dor de cabeça. Para isso, você precisará executar dois comandos:

``` 
$ chmod +x restore_db.sh
$ ./restore_db.sh
```

O primeiro comando dará permissão de execução ao script, o segundo irá executar o processo de restaução do banco de dados.

### Iniciando a API RESTful

Para iniciarmos a API, precisaremos utilizar apenas um comando:

``` 
$ docker-compose up app
```

Novamente, o comando acima irá demorar um pouco até ser totalmente concluído, mas no final, você terá o servidor com a API RESTful funcionando com sucesso.

## Rotas disponíveis

As rotas disponíveis são essas:

### Versões

* **Listar todas as versões de traduções disponíveis** - *GET* - `localhost:8080/api/v0/versions`
* **Visualizar dados de uma versão específica** - *GET* - `localhost:8080/api/v0/versions/{abreviação}`

### Testamentos

* **Listar todos os testamentos** - *GET* - `localhost:8080/api/v0/testaments`
* **Visualizar dados de um testamento específico** - *GET* - `localhost:8080/api/v0/testament/{parte dos testamentos}`

### Livros

* **Listar todos os livros** - *GET* - `localhost:8080/api/v0/books`
* **Visualizar dados de um livro específico** - *GET* - `localhost:8080/api/v0/books/{abreviação}`
* **Listar livros pertencentes a um testamento** - *GET* - `localhost:8080/api/v0/testaments/{parte dos testamentos}/books`

### Versículos

* **Listar todos os versículos de um capítulo a partir de uma versão de tradução + livro** - *GET* - `localhost:8080/api/v0/versions/{abreviação}/books/{abreviação}/chapter/{número do capítulo}/verses`
* **Listar todos os versículos de uma versão de tradução + livro** - *GET* - `localhost:8080/api/v0/versions/{abreviação}/books/{abreviação}/verses`
* **Visualizar dados de um versículo a partir de uma versão de tradução + capítulo + livro** - *GET* - `localhost:8080/api/v0/versions/{abreviação}/books/{abreviação}/chapter/{número do capítulo}/verses/{número do versículo}`
* **Pesquisar versículos de uma versão de tradução a partir de uma palavra/frase** - *POST* - `localhost:8080/api/v0/versions/{abreviação}/search` - Durante a pesquisa, envie uma requisição JSON (content-type: application/json) com o campo " `phrase` " (Ex: `{"phrase": "Jesus"}` ).

## Contribuindo

Caso você tenha achado o projeto legal e queira contribuir, fique a vontade para abrir um PR no momento que quiser! Ficarei muito feliz em receber a sua contribuição! :-)
