CREATE TABLE signins (
	id         CHAR(36),
	identity   CHAR(42),
	msg        VARCHAR(128),
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL,
	PRIMARY KEY (identity)
);

CREATE TABLE org_users (
	id           CHAR(36),
	identity     CHAR(42),
	roles        VARCHAR(512),
	provider_arn VARCHAR(64),
	role_arn     VARCHAR(64),
	aws_id       VARCHAR(16),
	is_active    TINYINT(1),
	created_at   TIMESTAMP NULL,
	updated_at   TIMESTAMP NULL,
	PRIMARY KEY (identity, roles)
);

CREATE TABLE users (
	id           CHAR(36),
	org_user     CHAR(42),
	roles        VARCHAR(512),
	identity     CHAR(42),
	is_active    TINYINT(1),
	created_at   TIMESTAMP NULL,
	updated_at   TIMESTAMP NULL,
	PRIMARY KEY (org_user, identity, roles)
);