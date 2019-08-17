/*
    build the sprint summary view
    name: vw_sprint_summary
    description: returns a sprint summary
    version: 0.0.2 - retrieve sprintID in the result set.
*/
CREATE OR REPLACE VIEW vw_sprint_summary AS
SELECT 
    map.workstream_id AS workstream_id,
    sn.id AS sprint_id,
    sn.name,
    SUM(sli.current_availability) AS working_days,
    SUM(sli.committed_points_this_sprint) AS committed_points,
    SUM(sli.completed_points_this_sprint) AS completed_points,
    SUM(sli.completed_points_last_sprint) AS completed_points_last_sprint
FROM sprint_line_item sli
INNER JOIN workstream_sprint_engineer_sprint_line_item_map map
ON map.sprint_line_item_id=sli.id
INNER JOIN sprint sn
ON sn.id=map.sprint_id

GROUP BY map.workstream_id, map.sprint_id
ORDER BY sn.id;

/*
    name: vw_sprint_detail_line_items
    description: engineer detail on a sprint detail line item.
    version: 0.0.1
*/
CREATE OR REPLACE VIEW vw_sprint_detail_line_items AS
SELECT 
    sn.id AS sprint_id,
    ws.id AS workstream_id,
    ed.first_name AS display_name,
    sli.current_availability AS current_availability,
    sli.previous_availability AS previous_availability,
    sli.capacity AS capacity,
    sli.target_points AS target_points,
    sli.committed_points_this_sprint AS committed_points
FROM engineer_details ed
INNER JOIN workstream_sprint_engineer_sprint_line_item_map map
ON map.engineer_id=ed.id
INNER JOIN sprint_line_item sli
ON sli.id=map.sprint_line_item_id
INNER JOIN workstream_name ws
ON ws.id=map.workstream_id

INNER JOIN sprint_name sn
ON sn.id=map.sprint_id;
