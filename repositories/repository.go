package repositories

type Repository struct {
    UsersDB      string 
    CartsDB      string 
    MinDB      string 
    AutocompleteDB string 
    TreeDB	string
    CodesDB      string 
    MaxDB      string 
    Years	[]string 
    FZs		[]string  
}

func NewRepository() *Repository{
	return &Repository{}
}
func (r *Repository)AddUsersDB(conf string){
	r.UsersDB = conf
 }
func (r *Repository)AddCartsDB(conf string){
	r.CartsDB = conf
}
func (r *Repository)AddTreeDB(conf string){
	r.TreeDB = conf
}
func (r *Repository)AddMinDB(conf string){
	r.MinDB = conf
}
func (r *Repository)AddAutocompleteDB(conf string){
	r.AutocompleteDB = conf
}
func (r *Repository)AddCodesDB(conf string){
	r.CodesDB = conf
} 
func (r *Repository)AddMaxDB(conf string){
	r.MaxDB = conf
}
func (r *Repository)AddYears(conf []string){
	r.Years = conf
} 
func (r *Repository)AddFZs(conf []string){
	r.FZs = conf
}
