package seeddata

import (
	"hr-system/common/dao/models"
	"log"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Gen(db *gorm.DB) {
	// Create Departments
	departments := []models.Department{
		{Name: "Engineering"},
		{Name: "HR"},
		{Name: "Marketing"},
	}

	// Create or update departments (ignore duplicates)
	for _, dept := range departments {
		// Use OnConflict to handle duplicates by ignoring them
		if err := db.Clauses(clause.OnConflict{
			DoNothing: true, // Ignores the insert if there's a conflict
		}).Create(&dept).Error; err != nil {
			log.Fatalf("Error seeding department: %v", err)
		}
	}

	// Create Positions
	positions := []models.Position{
		{Title: "Junior Developer", Level: "Junior"},
		{Title: "Senior Developer", Level: "Senior"},
		{Title: "Manager", Level: "Manager"},
	}

	// Insert Positions into DB (handle duplicates)
	for _, pos := range positions {
		if err := db.Clauses(clause.OnConflict{
			DoNothing: true, // Ignores the insert if there's a conflict
		}).Create(&pos).Error; err != nil {
			log.Fatalf("Error seeding position: %v", err)
		}
	}

	// Create Users
	users := []models.User{
		{Username: "john_doe", Password: "hashedpassword1", Email: "john@example.com", Phone: "123456789", Role: "user", Status: "active"},
		{Username: "admin_user", Password: "hashedpassword2", Email: "admin@example.com", Phone: "987654321", Role: "admin", Status: "active"},
	}

	// Insert Users into DB
	for _, user := range users {
		if err := db.Clauses(clause.OnConflict{
			DoNothing: true, // Ignores the insert if there's a conflict
		}).Create(&user).Error; err != nil {
			log.Fatalf("Error seeding user: %v", err)
		}
	}

	// Create Employees
	employees := []models.Employee{
		{Name: "Alice", Email: "alice@example.com", Phone: "111222333", DepartmentID: 1, PositionID: 1, HireDate: time.Now(), Salary: 50000},
		{Name: "Bob", Email: "bob@example.com", Phone: "444555666", DepartmentID: 1, PositionID: 2, HireDate: time.Now(), Salary: 70000},
		{Name: "Charlie", Email: "charlie@example.com", Phone: "777888999", DepartmentID: 2, PositionID: 3, HireDate: time.Now(), Salary: 60000},
	}

	// Insert Employees into DB (handle duplicates)
	for _, emp := range employees {
		if err := db.Clauses(clause.OnConflict{
			DoNothing: true, // Ignores the insert if there's a conflict
		}).Create(&emp).Error; err != nil {
			log.Fatalf("Error seeding employee: %v", err)
		}
	}

	log.Println("Seed data successfully inserted")
}
