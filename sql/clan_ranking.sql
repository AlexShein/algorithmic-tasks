WITH people_formatted_name AS
    (SELECT name,
         points,
         coalesce(clan,
         '[no clan specified]') AS clan_name
    FROM people )
SELECT
    clan_name,
    SUM(points) OVER (PARTITION BY clan_name) AS total_points,
    COUNT(*) OVER (PARTITION BY clan_name) AS total_people,
    RANK() OVER (PARTITION BY clan_name ORDER BY total_points desc) AS rank
FROM people_formatted_name
ORDER BY total_points;



SELECT temp.*,
    RANK() OVER (PARTITION BY temp.clan_name ORDER BY temp.total_points desc) AS rank
FROM (
    WITH people_formatted_name AS
    (SELECT name,
         points,
         coalesce(clan,
         '[no clan specified]') AS clan_name
    FROM people )
    SELECT
        clan_name,
        SUM(points) AS total_points,
        COUNT(*) AS total_people
    FROM people_formatted_name
    GROUP BY clan_name
) as temp
ORDER BY total_points DESC;


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