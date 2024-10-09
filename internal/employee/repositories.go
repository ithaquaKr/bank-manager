package employee

type Repository interface{}

// employee.MongoRepository implements the employee.Repository interface using MongoDB
type MongoRepository struct{}
