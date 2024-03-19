--Membuat database
CREATE DATABASE e_commerce;

--Membuat tabel
CREATE TABLE tags(
    tag_id VARCHAR(10) PRIMARY KEY,
    tag ENUM('Kulit', 'Suede', 'Satin', 'Cotton', 'Small', 'Medium', 'Large') NOT NULL
);

CREATE TABLE tas (
	tas_id VARCHAR(10) PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    stok INT NOT NULL,
    harga INT NOT NULL,
    CHECK (stok >= 0 OR harga >= 0)
);

CREATE TABLE tas_tags (
	tas_id VARCHAR(10) NOT NULL,
    tag_id VARCHAR(10) NOT NULL,
    FOREIGN KEY (tas_id) REFERENCES tas(tas_id),
    FOREIGN KEY (tag_id) REFERENCES tags(tag_id)
);

CREATE TABLE pembeli (
    pembeli_id VARCHAR(10) PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    alamat VARCHAR(255) NOT NULL,
    kota VARCHAR(255) NOT NULL,
    negara VARCHAR(255) NOT NULL
);

CREATE TABLE `order` (
    order_id VARCHAR(50) PRIMARY KEY, 
    tanggal DATE NOT NULL,
    pembeli_id VARCHAR(10) NOT NULL,
    FOREIGN KEY (pembeli_id) REFERENCES pembeli(pembeli_id)
);

CREATE TABLE order_tas (
    order_id VARCHAR(50) NOT NULL,
    tas_id VARCHAR(10) NOT NULL,
    jumlah_pembelian INT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES `order`(order_id),
    FOREIGN KEY (tas_id) REFERENCES tas(tas_id)
);

--Tabel read only untuk customer
CREATE TABLE pembeli_read_only AS (
    SELECT tas.nama, tas.stok, tas.harga, tags.tag FROM tas
    LEFT JOIN tas_tags ON tas_tags.tas_id = tas.tas_id
    LEFT JOIN tags ON tas_tags.tag_id = tags.tag_id
);

--Membuat user admin dan customer
CREATE USER 'admin'@'localhost' IDENTIFIED BY 'password';
CREATE USER 'customer'@'localhost' IDENTIFIED BY 'password'; 
GRANT ALL PRIVILEGES ON e_commerce.* TO 'admin'@'localhost'; --Admin mempunyai semua akses
GRANT SELECT ON e_commerce.pembeli_read_only TO 'customer'@'localhost'; --customer read only
FLUSH PRIVILEGES;

--Menampilkan akses admin dan customer
SHOW GRANTS FOR 'admin'@'localhost';
SHOW GRANTS FOR 'customer'@'localhost';