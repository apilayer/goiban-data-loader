/*

*/
USE goiban;
DROP TABLE IF EXISTS BANK_DATA;
CREATE TABLE BANK_DATA (
	id BIGINT(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
	bankcode VARCHAR(20),
	name VARCHAR(255),
	zip VARCHAR(10),
	city VARCHAR(255),
	bic VARCHAR(12),
	country VARCHAR(2) NULL,	# ISO-3160
	algorithm VARCHAR(10),		# identified for checksum algorithm
	created TIMESTAMP,
	last_update TIMESTAMP,
	INDEX bankcode_index (bankcode asc),
	INDEX bic_index (bic asc),
	UNIQUE INDEX (bankcode, country, bic)
) Engine=InnoDB, DEFAULT CHARSET=utf8;