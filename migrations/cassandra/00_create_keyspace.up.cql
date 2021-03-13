CREATE KEYSPACE IF NOT EXISTS goauth_service 
    WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3} 
    AND durable_writes = true;