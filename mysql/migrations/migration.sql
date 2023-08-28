
-- create table
CREATE TABLE dbo.users (
  `name` longtext,
  `email` varchar(191) UNIQUE NOT NULL,
  `password` longtext,
  PRIMARY KEY (`email`)
) 
ENGINE=InnoDB 
DEFAULT CHARSET=utf8mb4 
COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE dbo.logins (
  `id` bigint NOT NULL AUTO_INCREMENT,
 `last_login` TIMESTAMP NOT NULL DEFAULT NOW(),
  `email` varchar(191) NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`email`) REFERENCES users(`email`) ON UPDATE CASCADE ON DELETE CASCADE

) 
ENGINE=InnoDB 
DEFAULT CHARSET=utf8mb4 
COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE dbo.customers (
	name varchar(255) NOT NULL,
	email varchar(255) UNIQUE NOT NULL,
	phone varchar(255) UNIQUE NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
	PRIMARY KEY (`email`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE dbo.orders (
	id BIGINT NOT NULL auto_increment,
	name varchar(255) NOT NULL,
	user_email varchar(255) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
	PRIMARY KEY (`id`),
	FOREIGN KEY (`user_email`) REFERENCES customers(`email`) ON UPDATE CASCADE ON DELETE CASCADE
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;


-- insert
INSERT INTO dbo.customers
(name, email, phone, created_at, updated_at)
VALUES
('budi', 'budi@email', '0812345678', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('anto', 'anto@email', '0812345671', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('tono', 'tono@email', '0812345672', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('susi', 'susi@email', '0812345673', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('tuti', 'tuti@email', '0812345674', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('ari', 'ari@email', '0812345675', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
;

INSERT INTO dbo.orders
(name, user_email, created_at, updated_at)
VALUES
('budi order', 'budi@email', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('anto order', 'anto@email', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('tono order', 'tono@email', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('susi order', 'susi@email', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('tuti order', 'tuti@email', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('ari order', 'ari@email', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
;
