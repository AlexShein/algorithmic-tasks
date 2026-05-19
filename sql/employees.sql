 -- First version
WITH employees_formatted_names AS
    (SELECT id,
         manager_id,
         CONCAT(name,
         ' (', CAST(id AS varchar), ')') AS f_name
    FROM employees)
SELECT manager.id AS manager_id,
         ARRAY_AGG(other.f_name
ORDER BY  other.id) AS employee_names
FROM employees manager
INNER JOIN employees_formatted_names other
    ON manager.id = other.manager_id
GROUP BY  manager.id;

-- Second version
SELECT manager_id,
         ARRAY_AGG(CONCAT(name,
         ' (', CAST(id AS varchar), ')') ) as employee_names
FROM employees
WHERE manager_id is NOT null
ORDER BY  manager_id
GROUP BY  manager_id;