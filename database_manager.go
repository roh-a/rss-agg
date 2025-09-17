package main

import (
		"fmt"
		"log"
		"os"
		"database/sql"
		_ "github.com/lib/pq"


) 

//DatabaseManager handles RSS Agg database with Vault
type DatabaseManager struct {
	vaultClient *VaultClient
	db          *sql.DB
	leaseID string
}
// NewDatabaseManager creates database manager for RSS aggregator
func NewDatabaseManager() (*DatabaseManager, error) {
	//Create Vault client using .env settings 

	vaultClient, err := NewVaultClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create vault client : %w", err)
	}

	dm := &DatabaseManager{
		vaultClient: vaultClient,
	}


	 // Get initial database connection with Vault credentials
    if err := dm.RefreshConnection(); err != nil {
        return nil, fmt.Errorf("failed to establish initial database connection: %w", err)
    }
    
    return dm, nil


}


// RefreshConnection gets new Vault credentials and reconnects
func (dm *DatabaseManager) RefreshConnection() error {

	creds, err := dm.vaultClient.GetDatabaseCredentials()

	if err != nil {
        return fmt.Errorf("failed to get database credentials: %w", err)
    }
    
    // Close existing connection
    dm.Close()


	    // Build connection string with Vault credentials
    // Use your existing .env values for host, port, database name
    dbHost := os.Getenv("DB_HOST")     // localhost
    dbPort := os.Getenv("DB_PORT")     // 5432
    dbName := os.Getenv("DB_NAME")     // rssagg


	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
        creds.Username, creds.Password, dbHost, dbPort, dbName)
    
    // Connect with Vault credentials
    dm.db, err = sql.Open("postgres", connStr)
    if err != nil {
        return fmt.Errorf("failed to open database connection: %w", err)
    }
    
    // Test connection
    if err := dm.db.Ping(); err != nil {
        return fmt.Errorf("failed to ping database: %w", err)
    }
    
    log.Printf("RSS Aggregator: Connected to database with Vault user: %s (TTL: %d seconds)",
        creds.Username, creds.TTL)
    
    return nil
}
// GetDB returns database connection
func (dm *DatabaseManager) GetDB() *sql.DB {
    return dm.db
}

// Close cleans up database connection
func (dm *DatabaseManager) Close() error {
    if dm.db != nil {
        return dm.db.Close()
    }
    return nil
}


//TODO: Add function for refresh creds before expiration