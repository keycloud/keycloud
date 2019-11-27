create table if not exists users
(
    uuid varchar(36) not null
        constraint users_pk
            primary key,
    name text not null
        constraint users_name_check
            check (name <> ''::text),
    mail text not null,
    masterpasswd varchar(32) not null,
    createdatee timestamp
);

alter table users owner to postgres;

create table if not exists passwds
(
    entryid serial not null
        constraint passwds_pk
            primary key,
    uuid varchar(36) not null
        constraint passwds_users_uuid_fk
            references users,
    url text,
    passwd varchar(32) not null,
    username varchar(36)
);

alter table passwds owner to postgres;

create table if not exists sessions
(
    uuid varchar(36) not null,
    session_token varchar(32) not null,
    constraint sessions_pk
        primary key (session_token, uuid)
);

alter table sessions owner to postgres;

create table if not exists authenticators
(
    id bytea not null,
    credentialid bytea not null,
    publickey bytea,
    aaguid bytea not null,
    signcount integer not null,
    userid varchar(36)
);

alter table authenticators owner to postgres;