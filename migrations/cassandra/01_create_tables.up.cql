USE goauth_service;

CREATE TABLE IF NOT EXISTS client_platform (
    id uuid,
    name text,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS client_type (
    id uuid,
    name text,
    login_strategy text,
    confirm_by text,
    self_register boolean,
    self_recover boolean,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS client (
    client_platform_id uuid,
    client_type_id uuid,
    primary key (client_platform_id, client_type_id)
);

CREATE TABLE IF NOT EXISTS role (
    client_type_id uuid,
    name text,
    id uuid,
    primary key (client_type_id, name)
);
CREATE INDEX IF NOT EXISTS ON role(id);

CREATE TABLE IF NOT EXISTS scope (
    method text,
    url text,
    requests counter,
    primary key (method,  url)
);

CREATE TABLE IF NOT EXISTS permission (
    parent_id uuid,
    name text,
    id uuid,
    primary key (parent_id, name)
);
CREATE INDEX IF NOT EXISTS ON permission(id);

CREATE TABLE IF NOT EXISTS permission_scope (
    permission_id uuid,
    method text,
    url text,
    primary key (permission_id, method,  url)
);

CREATE TABLE IF NOT EXISTS role_permission (
    role_id uuid,
    permission_id uuid,
    primary key (role_id, permission_id)
);

CREATE TABLE IF NOT EXISTS user (
    id uuid,
    client_type_id uuid,
    role_id uuid,
    password text,
    active tinyint,
    expires_at timestamp,
    created_at timestamp,
    updated_at timestamp,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS user_phone (
    phone text,
    user_id uuid,
    created_at timestamp,
    updated_at timestamp,
    primary key (phone)
);
CREATE INDEX IF NOT EXISTS ON user_phone(user_id);

CREATE TABLE IF NOT EXISTS user_email (
    email text,
    user_id uuid,
    created_at timestamp,
    updated_at timestamp,
    primary key (email)
);
CREATE INDEX IF NOT EXISTS ON user_email(user_id);

CREATE TABLE IF NOT EXISTS user_login (
    login text,
    user_id uuid,
    created_at timestamp,
    updated_at timestamp,
    primary key (login)
);
CREATE INDEX IF NOT EXISTS ON user_login(user_id);

CREATE TABLE IF NOT EXISTS user_key (
    key text,
    user_id uuid,
    created_at timestamp,
    updated_at timestamp,
    primary key (key)
);
CREATE INDEX IF NOT EXISTS ON user_key(user_id);

CREATE TABLE IF NOT EXISTS session (
    client_platform_id uuid,
    client_type_id uuid,
    user_id uuid,
    id uuid,
    role_id uuid,
    ip inet,
    data text,
    created_at timestamp,
    updated_at timestamp,
    expires_at timestamp,
    primary key ((client_platform_id, client_type_id), user_id, id)
);
CREATE INDEX IF NOT EXISTS ON session(user_id);