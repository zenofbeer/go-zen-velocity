/*
*   name:           spGetSprintSummary
*   description:    returns sprint summary data by workstreamID
*   version:        0.0.1
*/
DROP PROCEDURE IF EXISTS spGetSprintSummary;

CREATE PROCEDURE spGetSprintSummary(
    workstreamID INT)
BEGIN
    SELECT * FROM vw_sprint_summary
    WHERE workstream_id=workstreamID;
END;

/*
    name:           spGetSprintLineItems
    description:    returns sprint line items by workstreamID and sprintID
    version:        0.0.1
*/
DROP PROCEDURE IF EXISTS spGetSprintLineItems;

CREATE PROCEDURE spGetSprintLineItems(
    workstreamID INT,
    sprintID INT)
BEGIN
    SELECT ed.first_name as name
    FROM engineer_details ed
    INNER JOIN workstream_sprint_engineer_sprint_line_item_map map
    ON map.workstream_id=workstreamID
    AND map.sprint_id=sprintID
    AND map.engineer_id=ed.id;
END;