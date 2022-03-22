create table alunos(
    id serial primary key,
    nome varchar,
    cpf varchar,
    rg varchar,
);

INSERT INTO alunos(nome, cpf, rg) VALUES
('Henrique', '149.149.149-45', '079870567'),
('Marcela', '222.222.222-90', '32312320')   