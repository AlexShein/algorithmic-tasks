SELECT
    RANK() OVER (ORDER BY temp.total_points desc) AS rank,
    temp.*
FROM (
    SELECT
        COALESCE(NULLIF(clan,''), '[no clan specified]') AS clan,
        SUM(points) AS total_points,
        COUNT(*) AS total_people
    FROM people
    GROUP BY clan
) as temp
ORDER BY total_points DESC;