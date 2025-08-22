# User Mapping Example in Go

This Go program demonstrates how to map between a detailed domain model (`UserEntity`) and a simplified data transfer object (`UserDTO`). 

## Overview

- **UserEntity:** Represents the full user model including sensitive fields like password and timestamps.
- **UserDTO:** A safe, trimmed-down version suitable for network communication, excluding sensitive data.
  
## Features

- Convert a `UserEntity` to a `UserDTO` for client communication, hiding sensitive fields.
- Update a `UserEntity` from a `UserDTO`, preserving sensitive/internal fields like password and creation timestamp.
- Split full names in the DTO into first and last names during conversion.
  
## Usage

The main function demonstrates:

1. Creating an example `UserEntity`.
2. Mapping it to a `UserDTO`.
3. Simulating an updated DTO from a client.
4. Mapping the updated DTO back to a `UserEntity` while preserving sensitive data.

This pattern is commonly used in applications to separate internal domain models from external API representations, improving security and maintainability.
