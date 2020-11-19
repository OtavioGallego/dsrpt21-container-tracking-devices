CREATE DATABASE IF NOT EXISTS dsrpt21;
USE dsrpt21;

CREATE TABLE produtos(
    id int auto_increment primary key,
    nome varchar(50) not null,
    tipo varchar(50) not null,
    nocivo boolean,
    preco_medio float,
    peso_medio float
);

CREATE TABLE containers(
    id int auto_increment primary key,
    produto_id int not null, FOREIGN KEY(produto_id) REFERENCES produtos(id) ON DELETE CASCADE,
    quantidade int,
    comprimento_externo float,
    largura_externa float,
    altura_externa float,
    comprimento_interno float,
    largura_interna float,
    altura_interna float,
    valor_carregamento float,
    peso_carregamento float,
    volume float
);

CREATE TABLE localizacoes(
    id int auto_increment primary key,
    container_id int not null, FOREIGN KEY(container_id) REFERENCES containers(id) ON DELETE CASCADE,
    latitude decimal(11, 8),
    longitude decimal(11, 8),
    data timestamp default current_timestamp()
);