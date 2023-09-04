create extension if not exists "uuid-ossp";

create table player (
    id uuid not null 
        constraint pk_player_id primary key
        constraint df_player_id default uuid_generate_v4(),
    name varchar(255) not null,
    birth_date timestamp not null,
    height double precision not null
);

create table team (
    id uuid not null 
        constraint pk_team_id primary key
        constraint df_team_id default uuid_generate_v4(),
    name varchar(255) not null,
    description text,
    founded_at date not null,
    terminated_at date,
    is_active boolean not null default true
);

create table team_player (
    player_id uuid not null
        constraint team_player_player_id references player (id),
    team_id uuid not null
        constraint team_player_team_id references team (id),
    player_number integer not null,
    joined_at timestamp not null,
    left_at timestamp null,

    constraint pk_team_player primary key (player_id, team_id)
);

-- Inserir jogadores fictícios
INSERT INTO player (id, name, birth_date, height)
VALUES
  ('0a322e71-6696-43e6-88ae-114b1dc7da01', 'LeBron James', '1984-12-30', 2.06),
  ('4f2ee54b-ef18-4a69-a71c-df309fc1ffeb', 'Stephen Curry', '1988-03-14', 1.91),
  ('5f66ad81-bc34-4c36-93de-24a4c5519d6c', 'Kevin Durant', '1988-09-29', 2.08);

-- Inserir times fictícios
INSERT INTO team (id, name, description, founded_at, terminated_at, is_active)
VALUES
  ('aa5e5a81-fb2a-4004-915d-315781e523a6', 'Los Angeles Lakers', 'One of the most successful NBA franchises', '1947-01-01', NULL, true),
  ('a19cfc09-8e7d-4daa-8ce3-7d098df88f19', 'Golden State Warriors', 'Known for three-point shooting', '1946-01-01', NULL, true),
  ('e5b0c53a-3f77-4f0a-b6b7-50bf54a48e9d', 'Brooklyn Nets', 'Recently acquired several star players', '1967-01-01', NULL, true);

-- Inserir associações jogador-time fictícias
INSERT INTO team_player (player_id, team_id, player_number, joined_at, left_at)
VALUES
  ('0a322e71-6696-43e6-88ae-114b1dc7da01', 'aa5e5a81-fb2a-4004-915d-315781e523a6', 6, '2021-01-01', NULL),
  ('4f2ee54b-ef18-4a69-a71c-df309fc1ffeb', 'a19cfc09-8e7d-4daa-8ce3-7d098df88f19', 30, '2021-01-15', NULL),
  ('5f66ad81-bc34-4c36-93de-24a4c5519d6c', 'e5b0c53a-3f77-4f0a-b6b7-50bf54a48e9d', 7, '2021-02-01', NULL);
