package main

import (
	"context"
	"fmt"
	"log"

	"github.com/smallstep/certificates/authority/admin"
	adminnosql "github.com/smallstep/certificates/authority/admin/db/nosql"
	"github.com/smallstep/nosql"
)

func main() {
	// Open the existing database
	dbPath := "authority/testdata/db/db.4admin"
	db, err := nosql.New("badgerv2", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Create admin database
	adminDB, err := adminnosql.New(db, admin.DefaultAuthorityID)
	if err != nil {
		log.Fatalf("Failed to create admin DB: %v", err)
	}

	// Get all admins
	ctx := context.Background()
	admins, err := adminDB.GetAdmins(ctx)
	if err != nil {
		log.Fatalf("Failed to get admins: %v", err)
	}

	fmt.Printf("Found %d admins:\n", len(admins))
	for i, adm := range admins {
		fmt.Printf("Admin %d:\n", i+1)
		fmt.Printf("  ID: %s\n", adm.Id)
		fmt.Printf("  AuthorityId: %s\n", adm.AuthorityId)
		fmt.Printf("  ProvisionerId: %s\n", adm.ProvisionerId)
		fmt.Printf("  Subject: %s\n", adm.Subject)
		fmt.Printf("  Type: %s\n", adm.Type)
		fmt.Printf("  CreatedAt: %v\n", adm.CreatedAt)
		fmt.Printf("  DeletedAt: %v\n", adm.DeletedAt)
		fmt.Println()
	}
}
