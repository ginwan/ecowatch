package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ginwan/ecowatch/apps/api/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FacilityHandler struct {
	Pool *pgxpool.Pool
}

func (h *FacilityHandler) GetFacilities(w http.ResponseWriter, r *http.Request) {
	rows, err := h.Pool.Query(context.Background(), "SELECT id, name, type, location, status FROM facilities")
	if err != nil {
		http.Error(w, "Failed to get facilities data", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var facilities []models.Facility

	for rows.Next() {
		var facility models.Facility
		err := rows.Scan(
			&facility.ID,
			&facility.Name,
			&facility.Type,
			&facility.Location,
			&facility.Status,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to access facility data %v", err), http.StatusInternalServerError)
			return
		}
		facilities = append(facilities, facility)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(facilities)
}

func (h *FacilityHandler) CreateFacilities(w http.ResponseWriter, r *http.Request) {
	var facility models.Facility

	err := json.NewDecoder(r.Body).Decode(&facility)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to parse facility data %v", err), http.StatusInternalServerError)
		return
	}

	sqlString := `INSERT INTO facilities(name, type, location, status)
		VALUES ($1, $2, $3, $4) 
		RETURNING id`

	err = h.Pool.QueryRow(context.Background(), sqlString, facility.Name, facility.Type, facility.Location, facility.Status).Scan(&facility.ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to create facility %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(facility)
}

func (h *FacilityHandler) GetFacilityByID(w http.ResponseWriter, r *http.Request) {
	facilityId := r.PathValue("id")

	id, err := strconv.Atoi(facilityId)
	if err != nil {
		http.Error(w, "Invalid facility id", http.StatusBadRequest)
		return
	}

	var facility models.Facility
	sqlString := `SELECT id, name, type, location, status FROM facilities WHERE id = $1`

	err = h.Pool.QueryRow(context.Background(), sqlString, id).Scan(&facility.ID,
		&facility.Name,
		&facility.Type,
		&facility.Location,
		&facility.Status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Facilities not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Failed to get facility data %v", err), http.StatusInternalServerError)
		return
	}

	facility.ID = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(facility)
}

func (h *FacilityHandler) UpdateFacility(w http.ResponseWriter, r *http.Request) {
	facilityID := r.PathValue("id")

	id, err := strconv.Atoi(facilityID)
	if err != nil {
		http.Error(w, "Invalid Facility ID", http.StatusBadRequest)
		return
	}

	var facility models.Facility
	// your fixed decode line here
	err = json.NewDecoder(r.Body).Decode(&facility)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Unable to parse facility data: %v", err),
			http.StatusInternalServerError,
		)
		return
	}

	sqlString := `UPDATE facilities SET 
			name = $1, 
			type = $2, 
			location = $3, 
			status = $4 
		WHERE 
			id = $5`

	row, err := h.Pool.Exec(context.Background(), sqlString,
		facility.Name, facility.Type, facility.Location, facility.Status, id,
	)

	if err != nil {
		http.Error(w, "Update facility data failed", http.StatusInternalServerError)
		return
	}

	if row.RowsAffected() == 0 {
		http.Error(w, "Facility not found", http.StatusNotFound)
		return
	}

	facility.ID = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(facility)
}

func (h *FacilityHandler) DeleteFacility(w http.ResponseWriter, r *http.Request) {
	facilityID := r.PathValue("id")

	id, err := strconv.Atoi(facilityID)
	if err != nil {
		http.Error(w, "Invalid Facility ID", http.StatusBadRequest)
		return
	}

	sqlString := `DELETE FROM facilities WHERE id = $1`

	row, err := h.Pool.Exec(context.Background(), sqlString, id)

	if err != nil {
		http.Error(w, "Delete facility failed", http.StatusInternalServerError)
		return
	}

	if row.RowsAffected() == 0 {
		http.Error(w, "Facility not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
