package seeders

import (
	"moses/database"
	"moses/models"
)

// =========================
// ROLE SEED
// =========================
func SeedRoles() {
	db := database.DB

	roles := []models.Role{
		{Code: "WL", Name: "Worship Leader"},
		{Code: "SGR", Name: "Singer"},
		{Code: "GTR", Name: "Guitar"},
		{Code: "BASS", Name: "Bass"},
		{Code: "KEYS", Name: "Keyboard"},
		{Code: "DRUM", Name: "Drum"},
	}

	for _, r := range roles {
		var role models.Role
		if err := db.Where("code = ?", r.Code).First(&role).Error; err != nil {
			db.Create(&r)
		}
	}
}

// =========================
// PLAYER SEED
// =========================
func SeedPlayers() {
	db := database.DB

	data := []struct {
		Name string
		Main string
		Add  []string
	}{
		{"Andi", "SGR", []string{"WL"}},
		{"Budi", "WL", []string{"SGR"}},
		{"Rina", "KEYS", []string{"SGR"}},
		{"Sinta", "SGR", []string{}},
		{"Rama", "GTR", []string{"SGR"}},
		{"Kevin", "DRUM", []string{}},
		{"Riko", "BASS", []string{}},

		{"Dina", "SGR", []string{"WL"}},
		{"Tina", "WL", []string{"SGR"}},
		{"Lina", "KEYS", []string{"SGR"}},
		{"Mira", "SGR", []string{}},
		{"Nina", "GTR", []string{"SGR"}},
		{"Vina", "DRUM", []string{}},
		{"Wina", "BASS", []string{}},

		{"Joko", "SGR", []string{"WL"}},
		{"Dodo", "WL", []string{"SGR"}},
		{"Eko", "KEYS", []string{"SGR"}},
		{"Fajar", "SGR", []string{}},
		{"Gilang", "GTR", []string{"SGR"}},
		{"Hadi", "DRUM", []string{}},
		{"Iwan", "BASS", []string{}},

		{"Lukas", "SGR", []string{"WL"}},
		{"Marcel", "WL", []string{"SGR"}},
		{"Nico", "KEYS", []string{"SGR"}},
		{"Owen", "SGR", []string{}},
		{"Petrus", "GTR", []string{"SGR"}},
		{"Quentin", "DRUM", []string{}},
		{"Rafael", "BASS", []string{}},
	}

	for _, d := range data {

		var mainRole models.Role
		if err := db.Where("code = ?", d.Main).First(&mainRole).Error; err != nil {
			continue
		}

		player := models.Player{
			Name:       d.Name,
			MainRoleID: mainRole.ID,
			Active:     true,
		}

		db.Create(&player)

		var addRoles []models.Role
		for _, rc := range d.Add {
			var r models.Role
			if err := db.Where("code = ?", rc).First(&r).Error; err == nil {
				addRoles = append(addRoles, r)
			}
		}

		if len(addRoles) > 0 {
			db.Model(&player).Association("Roles").Append(&addRoles)
		}
	}
}

// =========================
// MASTER SEED
// =========================
func SeedAll() {
	SeedRoles()
	SeedPlayers()
}
