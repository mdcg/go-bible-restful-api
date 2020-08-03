# Routes

## Versões

* *Listar todas as versões de traduções disponíveis* - **GET** - `api/v0/versions`
* *Visualizar dados de uma versão específica* - **GET** - `api/v0/versions/{abreviação}`

## Testamentos

* *Listar todos os testamentos* - **GET** - `api/v0/testaments`
* *Visualizar dados de um testamento específico* - **GET** - `api/v0/testament/{parte dos testamentos}`

## Livors

* *Listar todos os livros* - **GET** - `api/v0/books`
* *Visualizar dados de um livro específico* - **GET** - `api/v0/books/{abreviação}`
* *Listar livros pertencentes a um testamento* - **GET** - `api/v0/testaments/{parte dos testamentos}/books`

## Versículos

* *Listar todos os versículos de um capítulo a partir de uma versão de tradução + livro* - **GET** - `api/v0/versions/{abreviação}/books/{abreviação}/chapter/{número do capítulo}/verses`
* *Listar todos os versículos de uma versão de tradução + livro* - **GET** - `api/v0/versions/{abreviação}/books/{abreviação}/verses`
* *Visualizar dados de um versículo a partir de uma versão + capítulo + livro* - **GET** - `api/v0/versions/{abreviação}/books/{abreviação}/chapter/{número do capítulo}/verses/{número do versículo}`
* *Pesquisar versículos de uma versão a partir de uma palavra/frase* - **POST** - `api/v0/versions/{abreviação}/search` - Durante a pesquisa, envie uma requisição JSON (content-type: application/json) com o campo "`phrase`" (Ex: `{"phrase": "Jesus"}`).

# ToDo

* Adicionar logs
* Criar imagem Docker
* Criar sistema de cache > Pensar em async
