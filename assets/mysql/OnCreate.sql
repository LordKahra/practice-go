DELIMITER //

CREATE TRIGGER OnCreateCharacter AFTER INSERT ON characters
    FOR EACH ROW
BEGIN
    INSERT INTO hack_servers (name, character_id, tags) VALUES (CONCAT(NEW.name, '\'s PC'), NEW.id, 'character');
END;
//

DELIMITER ;




DELIMITER //

CREATE TRIGGER OnCreateServer AFTER INSERT ON hack_servers
    FOR EACH ROW
BEGIN
    DECLARE generated_ip VARCHAR(16);

    IF (NEW.character_id IS NOT NULL)
        THEN
            SET generated_ip = getNewCharacterIPv4();
        ELSE
            SET generated_ip = getNewTaggedIPv4('all');
    end if;

    INSERT INTO hack_ips (ipv4, server_id) VALUES (generated_ip, NEW.id);
END;
//

DELIMITER ;




DELIMITER //

CREATE TRIGGER OnCreateFile AFTER INSERT ON hack_server_files
    FOR EACH ROW
BEGIN
    DECLARE chara_id INT;

    IF (NEW.intel_id IS NOT NULL) THEN
        # Check if there's a character_id.
        SET chara_id = (SELECT hack_servers.character_id
                                FROM hack_servers
                                LEFT JOIN hack_server_files ON hack_servers.id = hack_server_files.server_id
                                WHERE hack_servers.character_id IS NOT NULL AND hack_server_files.id = NEW.id);

        IF (chara_id IS NOT NULL) THEN
            # Save the intel.
            INSERT INTO hack_r_character_intel (character_id, intel_id) VALUES (chara_id, NEW.intel_id);
        end if;
    end if;
END;
//

DELIMITER ;