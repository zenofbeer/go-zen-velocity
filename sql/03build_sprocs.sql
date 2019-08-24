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
    version:        0.1.1 - extend to get current availability
*/
DROP PROCEDURE IF EXISTS spGetSprintLineItems;

CREATE PROCEDURE spGetSprintLineItems(
    workstreamID INT,
    sprintID INT)
BEGIN
    SELECT 
        display_name, 
        current_availability,
        previous_availability,
        capacity,
        target_points,
        committed_points
    FROM vw_sprint_detail_line_items
    WHERE workstream_id=workstreamID
    AND sprint_id=sprintID;
END;