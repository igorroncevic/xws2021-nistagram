package controllers

import (
	"encoding/json"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	"net/http"
)

type PrivacyController struct {
	Service services.PrivacyService
}

func (controller *PrivacyController) CreatePrivacy(w http.ResponseWriter, r *http.Request) {
	var privacy *persistence.Privacy

	json.NewDecoder(r.Body).Decode(&privacy)
	//_, err := controller.Service.CreatePrivacy(privacy)
	//if err != nil {
	//	customerr.WriteErrToClient(w, err)
	//	return
	//}

	w.WriteHeader(http.StatusOK)

}
