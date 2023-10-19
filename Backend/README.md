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

Para teste do filtro por data de registro, caso todas os dados da coluna data_hora_registro estejam no mesmo dia(já que é configurado para pegar a data atual), você pode alterar a data diretamente no banco de dados, siga o exemplo abaixo:
UPDATE sinais_vitais
	SET data_hora_registro = '2023-10-18 08:16:43'
	WHERE frequencia_cardiaca = 65