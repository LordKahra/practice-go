CREATE VIEW hack_ip_details AS

SELECT ip.ipv4 as ipv4, server.id as server_id, server.name as server_name, ip.effective_date as effective_date,
       IF (server_details.id IS NULL, 'offline', 'online') AS status
FROM hack_ips ip
         LEFT JOIN hack_server_details server ON ip.server_id = server.id
         LEFT JOIN hack_server_details server_details ON ip.ipv4 = server_details.ipv4;