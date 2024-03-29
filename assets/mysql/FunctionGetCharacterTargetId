DELIMITER //

CREATE FUNCTION hav.getTargetTypeId (target_name VARCHAR(255))
    RETURNS INT DETERMINISTIC
BEGIN
    RETURN (SELECT id FROM hav.hack_z_target_types WHERE name = target_name);
END //