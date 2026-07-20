package domain

type UserRole string

const (
	RoleWorker           UserRole = "Worker"
	RoleSafetySupervisor UserRole = "SafetySupervisor"
	RolePlantManager     UserRole = "PlantManager"
)

// User represents the platform's core identity domain entity.
type User struct {
	ID        string   `json:"id"`
	Email     string   `json:"email"`
	Role      UserRole `json:"role"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
}
