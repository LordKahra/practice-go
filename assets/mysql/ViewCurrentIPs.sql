CREATE VIEW hack_current_ips AS

    SELECT ip.ipv4 as ipv4, ip.server_id as server_id, ip.effective_date
        FROM hack_ips ip
        INNER JOIN (SELECT server_id, MAX(effective_date) AS effective_date FROM hack_ips GROUP BY server_id) AS ip2
            ON ip.server_id = ip2.server_id AND ip.effective_date = ip2.effective_date;