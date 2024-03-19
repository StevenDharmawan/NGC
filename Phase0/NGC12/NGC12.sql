CREATE DATABASE tourism;

USE tourism;
CREATE TABLE destinations(
    destination_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    rating INT(2)
);

CREATE TABLE hotels(
    hotel_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    destination_id INT NOT NULL,
    rating INT(2),
    address VARCHAR(255) NOT NULL,
    FOREIGN KEY (destination_id) REFERENCES destinations(destination_id)
);

CREATE TABLE bookings(
    booking_id INT AUTO_INCREMENT PRIMARY KEY,
    guest_details VARCHAR(255) NOT NULL,
    hotel_references INT NOT NULL,
    check_in_dates DATE NOT NULL,
    check_out_dates DATE NOT NULL
);

INSERT INTO destinations (name, country, description, rating)
VALUES 	('Destination A', 'Country A', 'Description A', 7),
		('Destination B', 'Country B', 'Description B', 8),
        ('Destination C', 'Country C', 'Description C', 9);

INSERT INTO hotels (name, destination_id, rating, address)
VALUES 	('Hotel A', 1, 5, 'Adress A'),
		('Hotel B', 2, 6, 'Adress B'),
        ('Hotel C', 3, 7, 'Adress C'),
        ('Hotel D', 1, 8, 'Adress D'),
		('Hotel E', 2, 8, 'Adress E'),
        ('Hotel F', 3, 8, 'Adress F');

INSERT INTO bookings (guest_details, hotel_references, check_in_dates, check_out_dates)
VALUES 	('Guest Details A', 1, '2023-01-01', '2024-01-02'),
		('Guest Details B', 2, '2023-02-02', '2024-02-03'),
        ('Guest Details C', 3, '2023-03-03', '2024-03-04');

SELECT * FROM destinations;
SELECT * FROM hotels;
SELECT * FROM bookings;

SELECT * FROM destinations
WHERE name = "Destination A"

SELECT AVG(hotels.rating), destinations.name FROM hotels
LEFT JOIN destinations ON hotels.destination_id = destinations.destination_id
GROUP BY hotels.destination_id;