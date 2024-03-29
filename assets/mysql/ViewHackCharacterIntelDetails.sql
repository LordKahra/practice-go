CREATE VIEW hack_character_intel_details AS

SELECT intel.id as id, char_intel.character_id AS character_id, intel.name AS name, intel.type_id as type_id, hzit.name as type_name, intel.target as target
FROM hav.hack_r_character_intel char_intel
         LEFT JOIN hack_intel intel on char_intel.intel_id = intel.id
         LEFT JOIN hav.hack_z_intel_types hzit on intel.type_id = hzit.id;