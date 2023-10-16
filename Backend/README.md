Observações para utilização do Backend

Criação database "monitoramento_pacientes":
CREATE DATABASE monitoramento_pacientes;

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