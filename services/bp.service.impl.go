package services

import (
	"ProjectFirst/config"
	"ProjectFirst/dao"
	"ProjectFirst/models"
	"errors"
	"time"
)

var bpDao dao.BpDao = dao.BpDaoImpl{}

type BpServiceImp struct{}

// CreateBusinessPartner ...
func (BpServiceImp) CreateBusinessPartner(bp *models.BusinessPartner) (*models.BusinessPartner, error) {
	_, err := userDao.GetUserById(bp.UserId)
	// _, err := UserServiceImp.GetUserById(bp.UserId)
	if err != nil {
		return nil, errors.New("User id not exist")
	}

	var createDt = models.Timestamp(time.Now())
	bp.CreatedDate = &createDt

	return bpDao.CreateBp(bp)

}

func (BpServiceImp) GetBpById(id string) (bp models.BusinessPartner, e error) {
	defer config.CatchError(&e)

	bp, err := bpDao.GetBpById(id)

	return bp, err

}

func (BpServiceImp) GetBpByUserId(id string) (bp models.BusinessPartner, e error) {
	defer config.CatchError(&e)

	bp, err := bpDao.GetBpByUserId(id)

	return bp, err
}
