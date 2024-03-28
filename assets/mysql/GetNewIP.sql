DELIMITER //

CREATE FUNCTION hav.getNewIPv4 (section1 INT, section2 INT, section3 INT, section4 INT)

    RETURNS VARCHAR(16) DETERMINISTIC
BEGIN
    DECLARE generated_ip VARCHAR(15);
    DECLARE num_found INT;
    SET num_found = 1;

    find_new_ip:
    WHILE (num_found > 0) DO
            SET generated_ip = getRandomIPv4(section1, section2, section3, section4);
            SET num_found = (SELECT COUNT(ipv4) FROM hack_ips WHERE ipv4 = generated_ip);
        END WHILE find_new_ip;

    RETURN generated_ip;
END //

DELIMITER //

CREATE FUNCTION hav.getRandomIPv4 (section1 INT, section2 INT, section3 INT, section4 INT)

    RETURNS VARCHAR(16) DETERMINISTIC
BEGIN
    DECLARE generated_ip VARCHAR(15);

    IF (section1 < 2 OR section1 > 254) THEN SET section1 = FLOOR(2 + RAND() * (255 - 2)); END IF;
    IF (section2 < 2 OR section2 > 254) THEN SET section2 = FLOOR(2 + RAND() * (255 - 2)); END IF;
    IF (section3 < 2 OR section3 > 254) THEN SET section3 = FLOOR(2 + RAND() * (255 - 2)); END IF;
    IF (section4 < 2 OR section4 > 254) THEN SET section4 = FLOOR(2 + RAND() * (255 - 2)); END IF;

    SET generated_ip = CONCAT_WS('.', section1, section2, section3, section4);

    RETURN generated_ip;
END //