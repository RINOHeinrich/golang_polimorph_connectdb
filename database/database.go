package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConnector interface {
	Connect() (*gorm.DB, error)
}

func Connect(connector DatabaseConnector) (*gorm.DB, error) {
	db, err := connector.Connect()
	if err != nil {
		return nil, err
	}
	return db, nil
}

type MySQLConnector struct {
	// MySQL specific configuration fields
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func (m *MySQLConnector) Connect() (*gorm.DB, error) {
	// Implement the connection logic for MySQL using gorm
	// Example:
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", m.Username, m.Password, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	return db, err
}

type PostgreSQLConnector struct {
	// PostgreSQL specific configuration fields
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func (p *PostgreSQLConnector) Connect() (*gorm.DB, error) {
	// Implement the connection logic for PostgreSQL using gorm
	// Example:
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", p.Host, p.Port, p.Username, p.Password, p.Database)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	return db, err
}

type MongoDBConnector struct {
	// MongoDB specific configuration fields
	Host     string
	Port     int
	Username string
	Password string
	Document string
}

func (m *MongoDBConnector) Connect() (*mongo.Client, error) {
	// Implement the connection logic for MongoDB using mongo-driver
	// Example:
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d", m.Username, m.Password, m.Host, m.Port)
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	return client, err
}
