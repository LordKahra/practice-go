CREATE VIEW hack_server_details AS

SELECT server.id AS id, server.name as `name`, ip.ipv4 as `ipv4`,
       server.address as `address`, server.character_id as `character_id`,
       server.tags as `tags`, ip.effective_date as `ip_effective_date`
FROM hack_servers server
         LEFT JOIN hack_current_ips ip ON ip.server_id = server.id;