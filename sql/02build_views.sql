/*
    build the sprint summary view
*/
CREATE OR REPLACE VIEW vw_sprint_summary AS
SELECT 
    map.workstream_id AS workstream_id,
    sn.name,
    SUM(sli.current_availability) AS working_days,
    SUM(sli.committed_points_this_sprint) AS committed_points,
    SUM(sli.completed_points_this_sprint) AS completed_points,
    SUM(sli.completed_points_last_sprint) AS completed_points_last_sprint
FROM sprint_line_item sli
INNER JOIN workstream_sprint_engineer_sprint_line_item_map map
ON map.sprint_line_item_id=sli.id
INNER JOIN sprint_name sn
ON sn.id=map.sprint_id

GROUP BY map.workstream_id, map.sprint_id
ORDER BY sn.id;


select * from sprint_name
where id = (select max(id)
    from sprint_name
    where id < 1)