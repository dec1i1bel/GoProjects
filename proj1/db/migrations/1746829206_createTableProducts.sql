DROP TABLE IF EXISTS product;

CREATE TABLE product (
  id INT AUTO_INCREMENT NOT NULL,
  active TINYINT(1) DEFAULT 0,
  name VARCHAR(255) NOT NULL,
  description VARCHAR(1000) DEFAULT NULL,
  xml_id varchar (50) DEFAULT NULL,
  width_mm FLOAT(7,2) DEFAULT NULL,
  height_mm FLOAT(7,2) DEFAULT NULL,
  weight_gr FLOAT(7,2) DEFAULT NULL,
  include_ts TIMESTAMP /* mariadb-5.3 */ NOT NULL
    DEFAULT current_timestamp()
    ON UPDATE current_timestamp(),
  update_ts TIMESTAMP /* mariadb-5.3 */ NOT NULL
    DEFAULT current_timestamp()
    ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
);