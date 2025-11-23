package handlers

import (
	"net/http"
)

// DispositionStats represents disposition statistics
type DispositionStats struct {
	Status       string  `json:"status"`
	Count        int     `json:"count"`
	Percentage   float64 `json:"percentage"`
	AvgTalkTime  float64 `json:"avg_talk_time"`
	TotalTalkTime int    `json:"total_talk_time"`
}

// ListDispositionStats represents disposition stats for a specific list
type ListDispositionStats struct {
	ListID       int                `json:"list_id"`
	ListName     string             `json:"list_name"`
	TotalCalls   int                `json:"total_calls"`
	Dispositions []DispositionStats `json:"dispositions"`
}

// CampaignDispositionStats represents disposition stats for a specific campaign
type CampaignDispositionStats struct {
	CampaignID   string               `json:"campaign_id"`
	CampaignName string               `json:"campaign_name"`
	TotalCalls   int                  `json:"total_calls"`
	Dispositions []DispositionStats   `json:"dispositions"`
}

// GetKPIDispositions retrieves disposition statistics
// Supports filtering by list_id, campaign_id, or both
func (h *Handler) GetKPIDispositions(w http.ResponseWriter, r *http.Request) {
	listID := r.URL.Query().Get("list_id")
	campaignID := r.URL.Query().Get("campaign_id")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	groupBy := r.URL.Query().Get("group_by") // Options: "list", "campaign", "both"

	if groupBy == "" {
		groupBy = "campaign" // Default to campaign grouping
	}

	switch groupBy {
	case "list":
		h.getKPIByList(w, listID, campaignID, startDate, endDate)
	case "campaign":
		h.getKPIByCampaign(w, campaignID, startDate, endDate)
	case "both":
		h.getKPIBoth(w, campaignID, startDate, endDate)
	default:
		respondWithError(w, http.StatusBadRequest, "Invalid group_by parameter. Use 'list', 'campaign', or 'both'")
	}
}

// getKPIByList retrieves dispositions grouped by list
func (h *Handler) getKPIByList(w http.ResponseWriter, listID, campaignID, startDate, endDate string) {
	query := `
		SELECT
			vl.list_id,
			vl.list_name,
			log.status,
			COUNT(*) as count,
			AVG(log.length_in_sec) as avg_talk_time,
			SUM(log.length_in_sec) as total_talk_time
		FROM vicidial_log log
		INNER JOIN vicidial_lists vl ON log.list_id = vl.list_id
		WHERE 1=1
	`
	args := []interface{}{}

	if listID != "" {
		query += " AND log.list_id = ?"
		args = append(args, listID)
	}
	if campaignID != "" {
		query += " AND log.campaign_id = ?"
		args = append(args, campaignID)
	}
	if startDate != "" {
		query += " AND log.call_date >= ?"
		args = append(args, startDate)
	}
	if endDate != "" {
		query += " AND log.call_date <= ?"
		args = append(args, endDate)
	}

	query += `
		GROUP BY vl.list_id, vl.list_name, log.status
		ORDER BY vl.list_id, count DESC
	`

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve KPI data: "+err.Error())
		return
	}
	defer rows.Close()

	// Map to group by list
	listMap := make(map[int]*ListDispositionStats)

	for rows.Next() {
		var listID int
		var listName, status string
		var count int
		var avgTalkTime, totalTalkTime float64

		rows.Scan(&listID, &listName, &status, &count, &avgTalkTime, &totalTalkTime)

		if _, exists := listMap[listID]; !exists {
			listMap[listID] = &ListDispositionStats{
				ListID:       listID,
				ListName:     listName,
				TotalCalls:   0,
				Dispositions: []DispositionStats{},
			}
		}

		listMap[listID].TotalCalls += count
		listMap[listID].Dispositions = append(listMap[listID].Dispositions, DispositionStats{
			Status:        status,
			Count:         count,
			AvgTalkTime:   avgTalkTime,
			TotalTalkTime: int(totalTalkTime),
		})
	}

	// Calculate percentages
	lists := []ListDispositionStats{}
	for _, listStats := range listMap {
		for i := range listStats.Dispositions {
			listStats.Dispositions[i].Percentage = float64(listStats.Dispositions[i].Count) / float64(listStats.TotalCalls) * 100
		}
		lists = append(lists, *listStats)
	}

	respondWithSuccess(w, "KPI dispositions by list retrieved successfully", map[string]interface{}{
		"count": len(lists),
		"lists": lists,
	})
}

// getKPIByCampaign retrieves dispositions grouped by campaign
func (h *Handler) getKPIByCampaign(w http.ResponseWriter, campaignID, startDate, endDate string) {
	query := `
		SELECT
			log.campaign_id,
			vc.campaign_name,
			log.status,
			COUNT(*) as count,
			AVG(log.length_in_sec) as avg_talk_time,
			SUM(log.length_in_sec) as total_talk_time
		FROM vicidial_log log
		INNER JOIN vicidial_campaigns vc ON log.campaign_id = vc.campaign_id
		WHERE 1=1
	`
	args := []interface{}{}

	if campaignID != "" {
		query += " AND log.campaign_id = ?"
		args = append(args, campaignID)
	}
	if startDate != "" {
		query += " AND log.call_date >= ?"
		args = append(args, startDate)
	}
	if endDate != "" {
		query += " AND log.call_date <= ?"
		args = append(args, endDate)
	}

	query += `
		GROUP BY log.campaign_id, vc.campaign_name, log.status
		ORDER BY log.campaign_id, count DESC
	`

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve KPI data: "+err.Error())
		return
	}
	defer rows.Close()

	// Map to group by campaign
	campaignMap := make(map[string]*CampaignDispositionStats)

	for rows.Next() {
		var campaignID, campaignName, status string
		var count int
		var avgTalkTime, totalTalkTime float64

		rows.Scan(&campaignID, &campaignName, &status, &count, &avgTalkTime, &totalTalkTime)

		if _, exists := campaignMap[campaignID]; !exists {
			campaignMap[campaignID] = &CampaignDispositionStats{
				CampaignID:   campaignID,
				CampaignName: campaignName,
				TotalCalls:   0,
				Dispositions: []DispositionStats{},
			}
		}

		campaignMap[campaignID].TotalCalls += count
		campaignMap[campaignID].Dispositions = append(campaignMap[campaignID].Dispositions, DispositionStats{
			Status:        status,
			Count:         count,
			AvgTalkTime:   avgTalkTime,
			TotalTalkTime: int(totalTalkTime),
		})
	}

	// Calculate percentages
	campaigns := []CampaignDispositionStats{}
	for _, campStats := range campaignMap {
		for i := range campStats.Dispositions {
			campStats.Dispositions[i].Percentage = float64(campStats.Dispositions[i].Count) / float64(campStats.TotalCalls) * 100
		}
		campaigns = append(campaigns, *campStats)
	}

	respondWithSuccess(w, "KPI dispositions by campaign retrieved successfully", map[string]interface{}{
		"count":     len(campaigns),
		"campaigns": campaigns,
	})
}

// getKPIBoth retrieves dispositions grouped by both campaign and list
func (h *Handler) getKPIBoth(w http.ResponseWriter, campaignID, startDate, endDate string) {
	query := `
		SELECT
			log.campaign_id,
			vc.campaign_name,
			log.list_id,
			vl.list_name,
			log.status,
			COUNT(*) as count,
			AVG(log.length_in_sec) as avg_talk_time,
			SUM(log.length_in_sec) as total_talk_time
		FROM vicidial_log log
		INNER JOIN vicidial_campaigns vc ON log.campaign_id = vc.campaign_id
		INNER JOIN vicidial_lists vl ON log.list_id = vl.list_id
		WHERE 1=1
	`
	args := []interface{}{}

	if campaignID != "" {
		query += " AND log.campaign_id = ?"
		args = append(args, campaignID)
	}
	if startDate != "" {
		query += " AND log.call_date >= ?"
		args = append(args, startDate)
	}
	if endDate != "" {
		query += " AND log.call_date <= ?"
		args = append(args, endDate)
	}

	query += `
		GROUP BY log.campaign_id, vc.campaign_name, log.list_id, vl.list_name, log.status
		ORDER BY log.campaign_id, log.list_id, count DESC
	`

	rows, err := h.DB.Query(query, args...)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to retrieve KPI data: "+err.Error())
		return
	}
	defer rows.Close()

	type CampaignListKey struct {
		CampaignID string
		ListID     int
	}

	type CampaignListStats struct {
		CampaignID   string             `json:"campaign_id"`
		CampaignName string             `json:"campaign_name"`
		ListID       int                `json:"list_id"`
		ListName     string             `json:"list_name"`
		TotalCalls   int                `json:"total_calls"`
		Dispositions []DispositionStats `json:"dispositions"`
	}

	// Map to group by campaign and list
	dataMap := make(map[CampaignListKey]*CampaignListStats)

	for rows.Next() {
		var campaignID, campaignName, listName, status string
		var listID, count int
		var avgTalkTime, totalTalkTime float64

		rows.Scan(&campaignID, &campaignName, &listID, &listName, &status, &count, &avgTalkTime, &totalTalkTime)

		key := CampaignListKey{CampaignID: campaignID, ListID: listID}

		if _, exists := dataMap[key]; !exists {
			dataMap[key] = &CampaignListStats{
				CampaignID:   campaignID,
				CampaignName: campaignName,
				ListID:       listID,
				ListName:     listName,
				TotalCalls:   0,
				Dispositions: []DispositionStats{},
			}
		}

		dataMap[key].TotalCalls += count
		dataMap[key].Dispositions = append(dataMap[key].Dispositions, DispositionStats{
			Status:        status,
			Count:         count,
			AvgTalkTime:   avgTalkTime,
			TotalTalkTime: int(totalTalkTime),
		})
	}

	// Calculate percentages
	results := []CampaignListStats{}
	for _, stats := range dataMap {
		for i := range stats.Dispositions {
			stats.Dispositions[i].Percentage = float64(stats.Dispositions[i].Count) / float64(stats.TotalCalls) * 100
		}
		results = append(results, *stats)
	}

	respondWithSuccess(w, "KPI dispositions by campaign and list retrieved successfully", map[string]interface{}{
		"count":   len(results),
		"results": results,
	})
}
