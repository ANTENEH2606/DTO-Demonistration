// user_mapping.go
package main

import (
	"fmt"
	"strings"
	"time"
)

// UserEntity represents the full domain model with more fields and business logic.
// It contains all data related to a user, including sensitive information like the password.
type UserEntity struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string // sensitive data not to be exposed in DTO
	CreatedAt time.Time
}

// UserDTO is a simplified data transfer object used for communication,
// such as sending data over a network. It only contains the fields
// that should be exposed to the client.
type UserDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// EntityToDTO maps a UserEntity to a UserDTO.
// This function transforms the internal domain model into a network-safe representation.
func EntityToDTO(entity UserEntity) UserDTO {
	return UserDTO{
		ID:    entity.ID,
		Name:  entity.FirstName + " " + entity.LastName,
		Email: entity.Email,
	}
}

// DTOToEntity maps a UserDTO back to an existing UserEntity.
// It is used to update an existing user from incoming data, ensuring that
// sensitive or unchangeable fields (like Password and CreatedAt) are preserved.
func DTOToEntity(dto UserDTO, original UserEntity) UserEntity {
	// Split the full name from the DTO into first and last names.
	names := strings.Split(dto.Name, " ")
	firstName := ""
	lastName := ""
	if len(names) > 0 {
		firstName = names[0]
		if len(names) > 1 {
			lastName = names[1]
		}
	}

	// Create a new UserEntity.
	// We use values from the DTO for updatable fields (ID, FirstName, LastName, Email).
	// We retain the original entity's values for non-updatable fields (Password, CreatedAt).
	return UserEntity{
		ID:        dto.ID,
		FirstName: firstName,
		LastName:  lastName,
		Email:     dto.Email,
		Password:  original.Password,
		CreatedAt: original.CreatedAt,
	}
}

// main function to demonstrate the mapping process.
func main() {
	// 1. Define an initial UserEntity, simulating a user from a database.
	userEntity := UserEntity{
		ID:        "123",
		FirstName: "Tsegaye",
		LastName:  "G/Medehen",
		Email:     "Tsegaye@gmail.com",
		Password:  "123456",
		CreatedAt: time.Now(),
	}

	fmt.Println("Original UserEntity:")
	fmt.Printf("%+v\n\n", userEntity)

	// 2. Convert the entity to a DTO to be sent to a client.
	// Notice that the sensitive 'Password' and 'CreatedAt' fields are not included.
	userDTO := EntityToDTO(userEntity)

	fmt.Println("Converted UserDTO (for a client/network):")
	fmt.Printf("%+v\n\n", userDTO)

	// 3. Simulate an update by modifying the DTO.
	// For example, the user changes their name and email.
	updatedDTO := UserDTO{
		ID:    "123",
		Name:  "Mulatu Astatkie",    // New name
		Email: "Mulatu@example.com", // New email
	}

	fmt.Println("Simulated updated UserDTO from client:")
	fmt.Printf("%+v\n\n", updatedDTO)

	// 4. Convert the updated DTO back to a new entity, preserving the original entity's data.
	// This ensures that the password and creation date are not accidentally overwritten.
	updatedEntity := DTOToEntity(updatedDTO, userEntity)

	fmt.Println("Re-mapped UserEntity (after update):")
	fmt.Printf("%+v\n", updatedEntity)
}
