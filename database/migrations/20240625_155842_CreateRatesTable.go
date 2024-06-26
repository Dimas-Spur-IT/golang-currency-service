package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateRatesTable_20240625000000 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateRatesTable_20240625000000{}
	m.Created = "20240625000000"
	migration.Register("CreateRatesTable_20240625000000", m)
}

// Run the migrations
func (m *CreateRatesTable_20240625000000) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`
        CREATE TABLE rate (
            id INT AUTO_INCREMENT PRIMARY KEY,
            cur_id INT,
            date DATE,
            cur_abbreviation VARCHAR(10),
            cur_scale INT,
            cur_name VARCHAR(100),
            cur_official_rate FLOAT(12, 4)
        );
    `)
}

// Reverse the migrations
func (m *CreateRatesTable_20240625000000) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS rate")
}
