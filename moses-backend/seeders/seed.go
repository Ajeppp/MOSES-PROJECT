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
		{Code: "SINGER", Name: "Singer"},
		{Code: "GUITAR", Name: "Guitar"},
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
		{"Andi", "SINGER", []string{"WL"}},
		{"Budi", "WL", []string{"SINGER"}},
		{"Rina", "KEYS", []string{"SINGER"}},
		{"Sinta", "SINGER", []string{}},
		{"Rama", "GUITAR", []string{"SINGER"}},
		{"Kevin", "DRUM", []string{}},
		{"Riko", "BASS", []string{}},

		{"Dina", "SINGER", []string{"WL"}},
		{"Tina", "WL", []string{"SINGER"}},
		{"Lina", "KEYS", []string{"SINGER"}},
		{"Mira", "SINGER", []string{}},
		{"Nina", "GUITAR", []string{"SINGER"}},
		{"Vina", "DRUM", []string{}},
		{"Wina", "BASS", []string{}},

		{"Joko", "SINGER", []string{"WL"}},
		{"Dodo", "WL", []string{"SINGER"}},
		{"Eko", "KEYS", []string{"SINGER"}},
		{"Fajar", "SINGER", []string{}},
		{"Gilang", "GUITAR", []string{"SINGER"}},
		{"Hadi", "DRUM", []string{}},
		{"Iwan", "BASS", []string{}},

		{"Lukas", "SINGER", []string{"WL"}},
		{"Marcel", "WL", []string{"SINGER"}},
		{"Nico", "KEYS", []string{"SINGER"}},
		{"Owen", "SINGER", []string{}},
		{"Petrus", "GUITAR", []string{"SINGER"}},
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
