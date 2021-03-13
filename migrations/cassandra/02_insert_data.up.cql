USE goauth_service;

INSERT INTO client_platform(id, name) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c01, 'WEB');
INSERT INTO client_platform(id, name) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c02, 'IOS');
INSERT INTO client_platform(id, name) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c03, 'ANDROID');
INSERT INTO client_platform(id, name) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c04, 'DESKTOP');
INSERT INTO client_platform(id, name) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c05, 'TELEGRAM_BOT');

INSERT INTO client_type(id, name, register_strategy, login_strategy, recovery_strategy) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c01, 'ADMIN', null, 'STANDARD', 'OTP');
INSERT INTO client_type(id, name, register_strategy, login_strategy, recovery_strategy) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c02, 'DIRECTOR', null, 'PASSCODE', 'OTP');
INSERT INTO client_type(id, name, register_strategy, login_strategy, recovery_strategy) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c03, 'MANAGER', null, 'PASSCODE', 'OTP');
INSERT INTO client_type(id, name, register_strategy, login_strategy, recovery_strategy) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c04, 'CONSULTANT', null, 'STANDARD', 'OTP');
INSERT INTO client_type(id, name, register_strategy, login_strategy, recovery_strategy) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c05, 'CASHIER', null, 'STANDARD', 'OTP');
INSERT INTO client_type(id, name, register_strategy, login_strategy, recovery_strategy) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c06, 'GUEST', null, 'OTP', 'OTP');
INSERT INTO client_type(id, name, register_strategy, login_strategy, recovery_strategy) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c07, 'CUSTOMER', 'OTP', 'OTP', 'OTP');
INSERT INTO client_type(id, name, register_strategy, login_strategy, recovery_strategy) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c08, 'WEBHOOK', null, 'KEY', null);

INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c01, 5a3818a9-90f0-44e9-a053-3be0ba1e2c01);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c01, 5a3818a9-90f0-44e9-a053-3be0ba1e2c02);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c01, 5a3818a9-90f0-44e9-a053-3be0ba1e2c03);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c01, 5a3818a9-90f0-44e9-a053-3be0ba1e2c04);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c01, 5a3818a9-90f0-44e9-a053-3be0ba1e2c05);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c01, 5a3818a9-90f0-44e9-a053-3be0ba1e2c06);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c01, 5a3818a9-90f0-44e9-a053-3be0ba1e2c07);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c02, 5a3818a9-90f0-44e9-a053-3be0ba1e2c04);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c02, 5a3818a9-90f0-44e9-a053-3be0ba1e2c05);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c02, 5a3818a9-90f0-44e9-a053-3be0ba1e2c07);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c03, 5a3818a9-90f0-44e9-a053-3be0ba1e2c04);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c03, 5a3818a9-90f0-44e9-a053-3be0ba1e2c05);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c03, 5a3818a9-90f0-44e9-a053-3be0ba1e2c07);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c04, 5a3818a9-90f0-44e9-a053-3be0ba1e2c03);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c04, 5a3818a9-90f0-44e9-a053-3be0ba1e2c04);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c04, 5a3818a9-90f0-44e9-a053-3be0ba1e2c05);
INSERT INTO client(client_platform_id, client_type_id) VALUES (7d4a4c38-dd84-4902-b744-0488b80a4c05, 5a3818a9-90f0-44e9-a053-3be0ba1e2c08);


INSERT INTO role(client_type_id, name, id) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c01, 'DEFAULT', a1ca1301-4da9-424d-a9e2-578ae6dcde01);
INSERT INTO role(client_type_id, name, id) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c02, 'DEFAULT', a1ca1301-4da9-424d-a9e2-578ae6dcde02);
INSERT INTO role(client_type_id, name, id) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c03, 'DEFAULT', a1ca1301-4da9-424d-a9e2-578ae6dcde03);
INSERT INTO role(client_type_id, name, id) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c04, 'DEFAULT', a1ca1301-4da9-424d-a9e2-578ae6dcde04);
INSERT INTO role(client_type_id, name, id) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c05, 'DEFAULT', a1ca1301-4da9-424d-a9e2-578ae6dcde05);
INSERT INTO role(client_type_id, name, id) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c06, 'DEFAULT', a1ca1301-4da9-424d-a9e2-578ae6dcde06);
INSERT INTO role(client_type_id, name, id) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c07, 'DEFAULT', a1ca1301-4da9-424d-a9e2-578ae6dcde07);
INSERT INTO role(client_type_id, name, id) VALUES (5a3818a9-90f0-44e9-a053-3be0ba1e2c08, 'DEFAULT', a1ca1301-4da9-424d-a9e2-578ae6dcde08);


UPDATE scope SET requests = requests + 123 WHERE method='GET' AND url='/';
UPDATE scope SET requests = requests + 32 WHERE method='GET' AND url='/ping';
UPDATE scope SET requests = requests + 10 WHERE method='GET' AND url='/v1/user';
UPDATE scope SET requests = requests + 3 WHERE method='POST' AND url='/v1/user';
UPDATE scope SET requests = requests + 7 WHERE method='GET' AND url='/v1/user/:id';
UPDATE scope SET requests = requests + 2 WHERE method='PUT' AND url='/v1/user/:id';
UPDATE scope SET requests = requests + 1 WHERE method='DELETE' AND url='/v1/user/:id';

INSERT INTO permission(parent_id, name, id) VALUES (ffffffff-ffff-4fff-8fff-ffffffffffff, 'dashboard', 9cbb32da-e473-4312-8413-95524ec08c01);
INSERT INTO permission(parent_id, name, id) VALUES (ffffffff-ffff-4fff-8fff-ffffffffffff, 'user', 9cbb32da-e473-4312-8413-95524ec08c02);
INSERT INTO permission(parent_id, name, id) VALUES (9cbb32da-e473-4312-8413-95524ec08c02, 'user/list', 9cbb32da-e473-4312-8413-95524ec08c03);
INSERT INTO permission(parent_id, name, id) VALUES (9cbb32da-e473-4312-8413-95524ec08c02, 'user/create', 9cbb32da-e473-4312-8413-95524ec08c04);
INSERT INTO permission(parent_id, name, id) VALUES (9cbb32da-e473-4312-8413-95524ec08c02, 'user/update', 9cbb32da-e473-4312-8413-95524ec08c05);
INSERT INTO permission(parent_id, name, id) VALUES (9cbb32da-e473-4312-8413-95524ec08c02, 'user/delete', 9cbb32da-e473-4312-8413-95524ec08c06);
INSERT INTO permission(parent_id, name, id) VALUES (9cbb32da-e473-4312-8413-95524ec08c03, 'user/detail', 9cbb32da-e473-4312-8413-95524ec08c07);

INSERT INTO permission_scope(permission_id, method, url) VALUES (9cbb32da-e473-4312-8413-95524ec08c01, 'GET', '/');
INSERT INTO permission_scope(permission_id, method, url) VALUES (9cbb32da-e473-4312-8413-95524ec08c01, 'GET', '/ping');
INSERT INTO permission_scope(permission_id, method, url) VALUES (9cbb32da-e473-4312-8413-95524ec08c03, 'GET', '/');
INSERT INTO permission_scope(permission_id, method, url) VALUES (9cbb32da-e473-4312-8413-95524ec08c03, 'GET', '/v1/user');
INSERT INTO permission_scope(permission_id, method, url) VALUES (9cbb32da-e473-4312-8413-95524ec08c04, 'POST', '/v1/user');
INSERT INTO permission_scope(permission_id, method, url) VALUES (9cbb32da-e473-4312-8413-95524ec08c07, 'GET', '/');
INSERT INTO permission_scope(permission_id, method, url) VALUES (9cbb32da-e473-4312-8413-95524ec08c07, 'GET', '/ping');
INSERT INTO permission_scope(permission_id, method, url) VALUES (9cbb32da-e473-4312-8413-95524ec08c07, 'GET', '/v1/user/:id');
INSERT INTO permission_scope(permission_id, method, url) VALUES (9cbb32da-e473-4312-8413-95524ec08c05, 'PUT', '/v1/user/:id');
INSERT INTO permission_scope(permission_id, method, url) VALUES (9cbb32da-e473-4312-8413-95524ec08c06, 'DELETE', '/v1/user/:id');

INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde01, 9cbb32da-e473-4312-8413-95524ec08c01);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde01, 9cbb32da-e473-4312-8413-95524ec08c02);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde01, 9cbb32da-e473-4312-8413-95524ec08c03);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde01, 9cbb32da-e473-4312-8413-95524ec08c04);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde01, 9cbb32da-e473-4312-8413-95524ec08c05);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde01, 9cbb32da-e473-4312-8413-95524ec08c06);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde01, 9cbb32da-e473-4312-8413-95524ec08c07);

INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde02, 9cbb32da-e473-4312-8413-95524ec08c01);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde02, 9cbb32da-e473-4312-8413-95524ec08c02);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde02, 9cbb32da-e473-4312-8413-95524ec08c03);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde02, 9cbb32da-e473-4312-8413-95524ec08c07);

INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde03, 9cbb32da-e473-4312-8413-95524ec08c01);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde03, 9cbb32da-e473-4312-8413-95524ec08c02);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde03, 9cbb32da-e473-4312-8413-95524ec08c03);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde03, 9cbb32da-e473-4312-8413-95524ec08c04);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde03, 9cbb32da-e473-4312-8413-95524ec08c05);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde03, 9cbb32da-e473-4312-8413-95524ec08c06);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde03, 9cbb32da-e473-4312-8413-95524ec08c07);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde04, 9cbb32da-e473-4312-8413-95524ec08c01);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde05, 9cbb32da-e473-4312-8413-95524ec08c01);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde06, 9cbb32da-e473-4312-8413-95524ec08c01);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde07, 9cbb32da-e473-4312-8413-95524ec08c07);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde08, 9cbb32da-e473-4312-8413-95524ec08c03);
INSERT INTO role_permission(role_id, permission_id) VALUES (a1ca1301-4da9-424d-a9e2-578ae6dcde08, 9cbb32da-e473-4312-8413-95524ec08c06);


INSERT INTO user(id, client_type_id, role_id, password, active, expires_at, created_at, updated_at) VALUES (f799f1c5-ce5f-4fdd-ac23-f542247dcc01, 5a3818a9-90f0-44e9-a053-3be0ba1e2c01, a1ca1301-4da9-424d-a9e2-578ae6dcde01, '$argon2id$v=19$models=65536,t=3,p=4$ylXXHk9qJ4Cg4te0mAjOAQ$sCU9Ie/L4SRHbd3sOKs+qwhImzlKZr4ya6waEaCO6Uw', 1, '2021-05-01T11:21:59.001+0000', toTimeStamp(now()), toTimeStamp(now()));
INSERT INTO user(id, client_type_id, role_id, password, active, expires_at, created_at, updated_at) VALUES (f799f1c5-ce5f-4fdd-ac23-f542247dcc02, 5a3818a9-90f0-44e9-a053-3be0ba1e2c02, a1ca1301-4da9-424d-a9e2-578ae6dcde02, '$argon2id$v=19$models=65536,t=3,p=4$ylXXHk9qJ4Cg4te0mAjOAQ$sCU9Ie/L4SRHbd3sOKs+qwhImzlKZr4ya6waEaCO6Uw', 1, '2021-01-01T11:21:59.001+0000', toTimeStamp(now()), toTimeStamp(now()));
INSERT INTO user(id, client_type_id, role_id, password, active, expires_at, created_at, updated_at) VALUES (f799f1c5-ce5f-4fdd-ac23-f542247dcc03, 5a3818a9-90f0-44e9-a053-3be0ba1e2c03, a1ca1301-4da9-424d-a9e2-578ae6dcde03, '$argon2id$v=19$models=65536,t=3,p=4$ylXXHk9qJ4Cg4te0mAjOAQ$sCU9Ie/L4SRHbd3sOKs+qwhImzlKZr4ya6waEaCO6Uw', 1, '2021-05-01T11:21:59.001+0000', toTimeStamp(now()), toTimeStamp(now()));
INSERT INTO user(id, client_type_id, role_id, password, active, expires_at, created_at, updated_at) VALUES (f799f1c5-ce5f-4fdd-ac23-f542247dcc04, 5a3818a9-90f0-44e9-a053-3be0ba1e2c04, a1ca1301-4da9-424d-a9e2-578ae6dcde04, '$argon2id$v=19$models=65536,t=3,p=4$ylXXHk9qJ4Cg4te0mAjOAQ$sCU9Ie/L4SRHbd3sOKs+qwhImzlKZr4ya6waEaCO6Uw', 0, '2021-05-01T11:21:59.001+0000', toTimeStamp(now()), toTimeStamp(now()));
INSERT INTO user(id, client_type_id, role_id, password, active, expires_at, created_at, updated_at) VALUES (f799f1c5-ce5f-4fdd-ac23-f542247dcc05, 5a3818a9-90f0-44e9-a053-3be0ba1e2c05, a1ca1301-4da9-424d-a9e2-578ae6dcde05, '$argon2id$v=19$models=65536,t=3,p=4$ylXXHk9qJ4Cg4te0mAjOAQ$sCU9Ie/L4SRHbd3sOKs+qwhImzlKZr4ya6waEaCO6Uw', -1, '2021-05-01T11:21:59.001+0000', toTimeStamp(now()), toTimeStamp(now()));
INSERT INTO user(id, client_type_id, role_id, password, active, expires_at, created_at, updated_at) VALUES (f799f1c5-ce5f-4fdd-ac23-f542247dcc06, 5a3818a9-90f0-44e9-a053-3be0ba1e2c06, a1ca1301-4da9-424d-a9e2-578ae6dcde06, null, 1, '2021-05-01T11:21:59.001+0000', toTimeStamp(now()), toTimeStamp(now()));
INSERT INTO user(id, client_type_id, role_id, password, active, expires_at, created_at, updated_at) VALUES (f799f1c5-ce5f-4fdd-ac23-f542247dcc07, 5a3818a9-90f0-44e9-a053-3be0ba1e2c07, a1ca1301-4da9-424d-a9e2-578ae6dcde07, null, 1, '2021-05-01T11:21:59.001+0000', toTimeStamp(now()), toTimeStamp(now()));
INSERT INTO user(id, client_type_id, role_id, password, active, expires_at, created_at, updated_at) VALUES (f799f1c5-ce5f-4fdd-ac23-f542247dcc08, 5a3818a9-90f0-44e9-a053-3be0ba1e2c08, a1ca1301-4da9-424d-a9e2-578ae6dcde08, null, 1, '2021-05-01T11:21:59.001+0000', toTimeStamp(now()), toTimeStamp(now()));

INSERT INTO user_phone(phone, user_id, created_at, updated_at) VALUES ('+998987654322', f799f1c5-ce5f-4fdd-ac23-f542247dcc02, toTimeStamp(now()), toTimeStamp(now()));
INSERT INTO user_phone(phone, user_id, created_at, updated_at) VALUES ('+998987654323', f799f1c5-ce5f-4fdd-ac23-f542247dcc03, toTimeStamp(now()), toTimeStamp(now()));
INSERT INTO user_phone(phone, user_id, created_at, updated_at) VALUES ('+998987654324', f799f1c5-ce5f-4fdd-ac23-f542247dcc04, toTimeStamp(now()), toTimeStamp(now()));
INSERT INTO user_phone(phone, user_id, created_at, updated_at) VALUES ('+998987654325', f799f1c5-ce5f-4fdd-ac23-f542247dcc05, toTimeStamp(now()), toTimeStamp(now()));
INSERT INTO user_phone(phone, user_id, created_at, updated_at) VALUES ('+998987654326', f799f1c5-ce5f-4fdd-ac23-f542247dcc06, toTimeStamp(now()), toTimeStamp(now()));
INSERT INTO user_phone(phone, user_id, created_at, updated_at) VALUES ('+998987654327', f799f1c5-ce5f-4fdd-ac23-f542247dcc07, toTimeStamp(now()), toTimeStamp(now()));
INSERT INTO user_phone(phone, user_id, created_at, updated_at) VALUES ('+998987654328', f799f1c5-ce5f-4fdd-ac23-f542247dcc08, toTimeStamp(now()), toTimeStamp(now()));

INSERT INTO user_email(email, user_id, created_at, updated_at) VALUES ('email1@example.com', f799f1c5-ce5f-4fdd-ac23-f542247dcc01, toTimeStamp(now()), toTimeStamp(now()));
INSERT INTO user_email(email, user_id, created_at, updated_at) VALUES ('email2@example.com', f799f1c5-ce5f-4fdd-ac23-f542247dcc02, toTimeStamp(now()), toTimeStamp(now()));

INSERT INTO user_login(login, user_id, created_at, updated_at) VALUES ('AD2100000001', f799f1c5-ce5f-4fdd-ac23-f542247dcc01, toTimeStamp(now()), toTimeStamp(now()));

INSERT INTO user_key(key, user_id, created_at, updated_at) VALUES ('SomeSecretKeY', f799f1c5-ce5f-4fdd-ac23-f542247dcc08, toTimeStamp(now()), toTimeStamp(now()));
