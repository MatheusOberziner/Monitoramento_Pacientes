Observações para utilização do Backend

Criação database "monitoramento_pacientes":
CREATE DATABASE monitoramento-pacientes;

Criação tabela "pacientes":
CREATE TABLE pacientes (
  id SERIAL PRIMARY KEY,
	nome char(255),
	cpf char(14),
	sexo char(9),
	data_nascimento date,
	cidade char(100)
);

Criação tabela "sinais_vitais":
CREATE TABLE sinais_vitais (
  id_paciente INT,
	frequencia_cardiaca INT,
	pressao_arterial VARCHAR(20),
	temperatura DECIMAL(3, 1),
  saturacao_oxigenio DECIMAL(3, 1),
  data_hora_registro TIMESTAMP,
  FOREIGN KEY (id_paciente) REFERENCES pacientes (id)
);

Se quiser adicionar alguns pacientes pode seguir o sql abaixo:
INSERT INTO pacientes (nome, cpf, sexo, data_nascimento, cidade)
	VALUES ('Giovana Rafaela Regina Nogueira', '161.899.786-65', 'Feminino', '27/05/1960', 'Rio Branco');
INSERT INTO pacientes (nome, cpf, sexo, data_nascimento, cidade)
	VALUES ('Yuri Leandro Porto', '439.449.218-10', 'Masculino', '03/11/1988', 'João Pessoa');
INSERT INTO pacientes (nome, cpf, sexo, data_nascimento, cidade)
	VALUES ('Renato Ian Cardoso', '035.214.808-01', 'Masculino', '02/08/1981', 'Blumenau');
INSERT INTO pacientes (nome, cpf, sexo, data_nascimento, cidade)
	VALUES ('Cecília Sebastiana Nunes', '894.726.079-77', 'Feminino', '17/01/1989', 'Florianópolis');

Para teste do filtro de pesquisa por data de registro, caso todas os dados da coluna data_hora_registro estejam no mesmo dia(já que é configurado para pegar a data atual), você pode alterar a data diretamente no banco de dados, siga o exemplo abaixo:
UPDATE sinais_vitais
	SET data_hora_registro = '2023-10-18 08:16:43'
	WHERE frequencia_cardiaca = 65
