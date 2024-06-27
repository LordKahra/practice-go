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

    IF (section1 IS NULL OR section1 < 2 OR section1 > 254) THEN SET section1 = FLOOR(2 + RAND() * (255 - 2)); END IF;
    IF (section2 IS NULL OR section2 < 2 OR section2 > 254) THEN SET section2 = FLOOR(2 + RAND() * (255 - 2)); END IF;
    IF (section3 IS NULL OR section3 < 2 OR section3 > 254) THEN SET section3 = FLOOR(2 + RAND() * (255 - 2)); END IF;
    IF (section4 IS NULL OR section4 < 2 OR section4 > 254) THEN SET section4 = FLOOR(2 + RAND() * (255 - 2)); END IF;

    SET generated_ip = CONCAT_WS('.', section1, section2, section3, section4);

    RETURN generated_ip;
END //

DELIMITER //

CREATE FUNCTION hav.getNewCharacterIPv4 ()
    RETURNS VARCHAR(16) DETERMINISTIC
BEGIN
    DECLARE section1 INT;
    SET section1 = (SELECT rule.argument FROM hav.hack_z_ip_rules rule WHERE rule.tag = 'character');

    RETURN hav.getNewIPv4(section1, 0,0,0);
END //

DELIMITER //

CREATE FUNCTION hav.getNewTaggedIPv4 (ip_tag varchar(255))
    RETURNS VARCHAR(16) DETERMINISTIC
BEGIN
    DECLARE section1 INT;
    DECLARE min INT;
    DECLARE max INT;

    IF (ip_tag IS NOT NULL AND ip_tag != 'all')
    THEN
        SET section1 = (SELECT rule.argument FROM hav.hack_z_ip_rules rule WHERE rule.tag = ip_tag);
    ELSE
        SET min = (SELECT rule.argument FROM hav.hack_z_ip_rules rule WHERE rule.operator = '>' AND rule.tag = 'all');
        SET max = (SELECT rule.argument FROM hav.hack_z_ip_rules rule WHERE rule.operator = '<' AND rule.tag = 'all');
        SET section1 = FLOOR(min + RAND() * (max - min));
    end if;

    IF (section1 IS NULL) THEN SET section1 = 0; end if;

    RETURN hav.getNewIPv4(section1, 0,0,0);
END //