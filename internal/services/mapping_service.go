package mappingservice

import (
	"dottest/internal/db"
	"dottest/internal/models"
)

func InteractiveUpsertDomain(domain string, target string) error {
	existingMapping := FindByDomain(domain)
	if existingMapping == nil {
		return Add(domain, target)
	}
	// TODO: Ask the user if they wish to update the existing mapping
	existingMapping.Domain = domain
	existingMapping.Target = target
	db := db.InitDB()
	return db.Save(existingMapping).Error
}

func Add(domain string, target string) error {
	db := db.InitDB()
	err := db.Create(&models.Mapping{
		Domain: domain,
		Target: target,
	})
	return err.Error
}

func FindByDomain(domain string) *models.Mapping {
	db := db.InitDB()
	mapping := &models.Mapping{}
	result := db.First(mapping, "domain = ?", domain)
	if result.Error != nil {
		return nil
	}
	return mapping
}
