# Go Pgx Adapter for Casbin

This is a simple casbin adapter implementation that uses go's pgx driver. It implement filter, context and batch interfaces of casbin adapters. It does not handle any database or table creation, it connects to your *pgxpool.Pool instance with a given name.