package campaign

type Repository interface {
	Save(campaign *Campaign) error
	All() ([]Campaign, error)
	Get(id string) (Campaign, error)
	Update(id string, values *Campaign) error
	Delete(id string) error
}
