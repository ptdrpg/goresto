package repository

import "github.com/ptdrpg/resto/entity"

func (r *Repository) FindEntrepriseById(id int) (entity.Entreprise, error) {
	var entreprise entity.Entreprise
	result := r.DB.Find(&entreprise, id)
	if result != nil {
		return entreprise, result.Error
	}

	return entreprise, nil
}

func (r *Repository) CreateEntreprise(entreprise *entity.Entreprise) error {
	err := r.DB.Create(entreprise)
	if err != nil {
		return err.Error
	}

	return nil
}

func (r *Repository) UpdateEntreprise(entreprise *entity.Entreprise) error {
	err := r.DB.Model(entreprise).Updates(entreprise)
	if err != nil {
		return err.Error
	}

	return nil
}

func (r *Repository) DeleteEntreprise(id int) error {
	var entreprise entity.Entreprise
	if err := r.DB.Where("id = ?", id).Delete(&entreprise).Error; err != nil {
		return err
	}

	return nil
}
