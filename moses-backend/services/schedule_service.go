package services

import "moses/models"

type ScheduleView struct {
	ServiceDate string              `json:"service_date"`
	Teams       map[string][]string `json:"teams"`
}

func BuildScheduleView(schedules []models.ServiceSchedule) []ScheduleView {

	grouped := make(map[string]map[string][]string)
	// date -> role -> []player

	for _, s := range schedules {
		date := s.ServiceDate.Format("2006-01-02")

		if _, ok := grouped[date]; !ok {
			grouped[date] = make(map[string][]string)
		}

		grouped[date][s.Role] = append(
			grouped[date][s.Role],
			s.Player.Name,
		)
	}

	var result []ScheduleView

	for date, roles := range grouped {
		result = append(result, ScheduleView{
			ServiceDate: date,
			Teams:       roles,
		})
	}

	return result
}
