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

	// 1️⃣ get all active players
	var players []models.Player
	if err := db.
		Preload("MainRole").
		Preload("Roles").
		Where("active = true").
		Find(&players).Error; err != nil {
		return err
	}

	if len(players) == 0 {
		return errors.New("no active players found")
	}

	// 2️⃣ unavailability
	var unavs []models.Unavailability
	db.Where("month = ? AND year = ?", month, year).Find(&unavs)

	unavailable := map[uint]bool{}
	for _, u := range unavs {
		unavailable[u.PlayerID] = true
	}

	// 3️⃣ available players
	var available []models.Player
	for _, p := range players {
		if !unavailable[p.ID] {
			available = append(available, p)
		}
	}

	if len(available) == 0 {
		return errors.New("no available players")
	}

	// 4️⃣ get saturdays
	serviceDates := getAllSaturdays(month, year)
	if len(serviceDates) == 0 {
		return errors.New("no saturday found")
	}

	// 5️⃣ generate
	for _, date := range serviceDates {

		// clean old schedule
		db.Where("service_date = ?", date).Delete(&models.ServiceSchedule{})

		usedPlayer := map[uint]bool{}

		for roleCode, qty := range teamComposition {

			selected := pickPlayersForRole(roleCode, qty, available, usedPlayer)

			if len(selected) < qty {
				return errors.New("not enough players for role: " + roleCode)
			}

			for _, p := range selected {
				usedPlayer[p.ID] = true

				db.Create(&models.ServiceSchedule{
					ServiceDate: date,
					Month:       month,
					Year:        year,
					Role:        roleCode,
					PlayerID:    p.ID,
				})
			}
		}
	}

	return nil
}

// =========================
// ROLE PICKER ENGINE
// =========================
func pickPlayersForRole(roleCode string, qty int, players []models.Player, used map[uint]bool) []models.Player {

	main := []models.Player{}
	additional := []models.Player{}

	for _, p := range players {

		if used[p.ID] {
			continue
		}

		// main role match
		if p.MainRole.Code == roleCode {
			main = append(main, p)
			continue
		}

		// additional role match
		for _, r := range p.Roles {
			if r.Role.Code == roleCode {
				additional = append(additional, p)
				break
			}
		}
	}

	// shuffle fairness
	rand.Shuffle(len(main), func(i, j int) { main[i], main[j] = main[j], main[i] })
	rand.Shuffle(len(additional), func(i, j int) { additional[i], additional[j] = additional[j], additional[i] })

	selected := []models.Player{}

	// priority main role
	for _, p := range main {
		if len(selected) < qty {
			selected = append(selected, p)
		}
	}

	// fallback additional role
	if len(selected) < qty {
		for _, p := range additional {
			if len(selected) < qty {
				selected = append(selected, p)
			}
		}
	}

	return selected
}

// =========================
// UTIL
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
