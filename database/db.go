package database

import
(
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "repsheets-go/models"
)

var db *gorm.DB

func Connect() error {
    var err error
    dsn := "test.db"  // SQLite bestandsnaam
    
    db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
    if err != nil {
        return err  // Geef fout terug als verbinding mislukt
    }
    
    // Automatisch tabellen aanmaken/updaten
    err = db.AutoMigrate(&models.User{})
    if err != nil {
        return err
    }
    
    return nil  // Geen fout = succes
}

func GetDB() *gorm.DB {
    return db
}