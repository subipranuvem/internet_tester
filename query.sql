SELECT
    r.id,
    r.request_address,
    r.status_code,
    r.method,
    r.response_time_ms,
    r.timeout_time_ms,
    r.created_at,
    s.ping_ms,
    s.download_speed_mbps,
    s.upload_speed_mbps,
    s.isp,
    s.ip_address,
    s.packet_loss_dup,
    s.packet_loss_max,
    s.packet_loss_sent,
    s.packet_loss_percentage
FROM
    requests r
    JOIN speed_test_results s ON r.speed_test_result_id = s.id
ORDER BY
    r.created_at DESC;