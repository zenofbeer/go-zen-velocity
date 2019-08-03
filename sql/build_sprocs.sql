DROP PROCEDURE IF EXISTS spGetSprintSummary;

CREATE PROCEDURE spGetSprintSummary(
    workstreamID INT)
BEGIN
    SELECT * FROM vw_sprint_summary
    WHERE workstream_id=workstreamID;
END;