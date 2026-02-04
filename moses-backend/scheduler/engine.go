package scheduler

import (
	"errors"
	"math/rand"
	"time"

	"moses/config"
	"moses/database"
	"moses/models"
)

// =========================
// TEAM COMPOSITION
// =========================
var teamComposition = map[string]int{
	config.RoleWL:     2,
	config.RoleSinger: 2,
	config.RoleBass:   1,
	config.RoleKeys:   1,
	config.RoleDrum:   1,
	config.RoleGuitar: 1,
}

// =========================
// MAIN ENGINE
// =========================
func GenerateSchedule(month int, year int) error {
	db := database.DB

	// 1️⃣ Ambil semua player aktif
	var players []models.Player
	if err := db.Where("active = true").Find(&players).Error; err != nil {
		return err
	}

	if len(players) == 0 {
		return errors.New("no active players found")
	}

	// 2️⃣ Ambil semua unavailability bulan tsb
	var unavs []models.Unavailability
	if err := db.
		Where("month = ? AND year = ?", month, year).
		Find(&unavs).Error; err != nil {
		return err
	}

	// 3️⃣ Build map unavailable
	unavailableMap := make(map[uint]bool)
	for _, u := range unavs {
		unavailableMap[u.PlayerID] = true
	}

	// 4️⃣ Filter available players + group by role
	roleMap := make(map[string][]models.Player)

	for _, p := range players {
		if !unavailableMap[p.ID] {
			roleMap[p.Role] = append(roleMap[p.Role], p)
		}
	}

	// 5️⃣ Ambil semua hari sabtu dalam bulan
	serviceDates := getAllSaturdays(month, year)

	if len(serviceDates) == 0 {
		return errors.New("no saturday found in this month")
	}

	// 6️⃣ Generate per sabtu
	for _, date := range serviceDates {

		// hapus jadwal lama (safe regenerate)
		db.Where("service_date = ?", date).Delete(&models.ServiceSchedule{})

		for role, qty := range teamComposition {

			playersByRole := roleMap[role]

			if len(playersByRole) < qty {
				return errors.New("not enough players for role: " + role)
			}

			// shuffle (fair + random)
			rand.Shuffle(len(playersByRole), func(i, j int) {
				playersByRole[i], playersByRole[j] = playersByRole[j], playersByRole[i]
			})

			selected := playersByRole[:qty]

			for _, p := range selected {
				schedule := models.ServiceSchedule{
					ServiceDate: date,
					Month:       month,
					Year:        year,
					Role:        role,
					PlayerID:    p.ID,
				}
				db.Create(&schedule)
			}
		}
	}

	return nil
}

// =========================
// UTIL FUNCTION
// =========================
func getAllSaturdays(month int, year int) []time.Time {
	var saturdays []time.Time

	loc := time.Now().Location()
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)

	for d := start; d.Month() == start.Month(); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == time.Saturday {
			saturdays = append(saturdays, d)
		}
	}

	return saturdays
}
