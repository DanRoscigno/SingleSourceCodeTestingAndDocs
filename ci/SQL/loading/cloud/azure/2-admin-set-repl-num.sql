-- Note: set this to the number of replicas you have. In a dev system
-- you might have only one BE, this sets the replication number to 1:
ADMIN SET FRONTEND CONFIG ('default_replication_num' = "1");
