# Monitoramento_Pacientes
Alunos: Matheus Schmidt Oberziner e Diego Muniz

Para configuração de ambiente você pode seguir este passo a passo: (Antes de tudo, clonar o repositório)

Para configuração do Frontend:
1. Abrir o terminal e rodar `cd Frontend`
2. Rodar `npm install` ou `yarn` para instalação de dependências e pacotes
3. Para subir o servidor basta rodar `npm run dev` ou `yarn quasar dev`
4. Irá abrir no navegador na porta 9000

Para configuração do Backend e do Banco de Dados(Postgres):
1. Para criação de database e tabelas você pode seguir no arquivo readme presente no diretório `./Backend/README.md`
2. Para conexão do banco com o backend você pode verificar no arquivo `./Backend/configs/config.go`, onde há uma função `init`, que é responsável pela configuração da conexão. (Eu deixei com a configuração padrão do postgres, mas se quiser mudar basta alterar os valores)
3. Abrir o terminal e rodar `cd Backend`
4. Para subir o servidor basta rodar `go run main.go`
5. Assim, todos os endpoints já estão disponíveis para serem consumidos no Frontend, e já pode usar o sistema
