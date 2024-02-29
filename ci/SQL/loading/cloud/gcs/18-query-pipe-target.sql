SELECT ItemID, CategoryID
FROM mydatabase.user_behavior_from_pipe
WHERE UserID = 1 AND DATE(Timestamp) = "2017-11-25";
