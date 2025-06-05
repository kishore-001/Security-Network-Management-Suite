create table if not exists users (
    id uuid primary key default gen_random_uuid(),
    name varchar(255) not null,
    role varchar(50) not null,
    email varchar(255) unique not null,
    password_hash varchar(100) not null 
);
