/*

*/
USE goiban;
DROP TABLE IF EXISTS DATA_SOURCE;
CREATE TABLE DATA_SOURCE (
	id INT PRIMARY KEY,
	name VARCHAR(255) UNIQUE
) Engine=MyISAM, DEFAULT CHARSET=utf8;

INSERT INTO DATA_SOURCE (id, name) VALUES
(1, "German Bundesbank");

INSERT INTO DATA_SOURCE (id, name) VALUES
(2, "NBB");

INSERT INTO DATA_SOURCE (id, name) VALUES
(3, "NL");

DROP TABLE IF EXISTS BANK_DATA;
CREATE TABLE BANK_DATA (
	id BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
	source VARCHAR(255) NOT NULL,
	bankcode VARCHAR(20),
	name VARCHAR(255),
	zip VARCHAR(10),
	city VARCHAR(255),
	bic VARCHAR(12),
	country VARCHAR(2) NULL,	# ISO-3160
	algorithm VARCHAR(10),		# identifier for checksum algorithm
	created TIMESTAMP,
	last_update TIMESTAMP,
	INDEX bankcode_index (bankcode asc),
	INDEX bic_index (bic asc),
	INDEX source_id (source),
	UNIQUE INDEX (bankcode, country, bic),
	FOREIGN KEY (source) references DATA_SOURCE(id)
) Engine=MyISAM, DEFAULT CHARSET=utf8;
