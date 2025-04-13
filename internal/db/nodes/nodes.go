package nodes

import (
	"database/sql"
	"fmt"

	"github.com/ByteForge-Systems/vpn-api/internal/models"
)

type NodeStore struct {
	db *sql.DB
}

func NewNodeStore(db *sql.DB) *NodeStore {
	return &NodeStore{db: db}
}

// Create inserts a new node into the database.
func (s *NodeStore) Create(node *models.Node) error {
	query := `
		INSERT INTO nodes (ip, port, ssh_port, country, comment, is_online, username, password)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`
	return s.db.QueryRow(query, node.IP, node.Port, node.SSHPort, node.Country, node.Comment, node.IsOnline, node.Username, node.Password).
		Scan(&node.ID)
}

// Get retrieves a node by ID.
func (s *NodeStore) Get(id int) (*models.Node, error) {
	node := &models.Node{}
	query := `
		SELECT id, ip, port, ssh_port, country, comment, is_online, username, password
		FROM nodes
		WHERE id = $1`
	err := s.db.QueryRow(query, id).Scan(
		&node.ID, &node.IP, &node.Port, &node.SSHPort, &node.Country,
		&node.Comment, &node.IsOnline, &node.Username, &node.Password,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("node not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get node: %w", err)
	}
	return node, nil
}

// List retrieves all nodes.
func (s *NodeStore) List() ([]*models.Node, error) {
	query := `
		SELECT id, ip, port, ssh_port, country, comment, is_online, username, password
		FROM nodes`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to list nodes: %w", err)
	}
	defer rows.Close()

	var nodes []*models.Node
	for rows.Next() {
		node := &models.Node{}
		if err := rows.Scan(
			&node.ID, &node.IP, &node.Port, &node.SSHPort, &node.Country,
			&node.Comment, &node.IsOnline, &node.Username, &node.Password,
		); err != nil {
			return nil, fmt.Errorf("failed to scan node: %w", err)
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

// Update modifies an existing node.
func (s *NodeStore) Update(node *models.Node) error {
	query := `
		UPDATE nodes
		SET ip = $1, port = $2, ssh_port = $3, country = $4, comment = $5, is_online = $6, username = $7, password = $8
		WHERE id = $9`
	result, err := s.db.Exec(query, node.IP, node.Port, node.SSHPort, node.Country, node.Comment, node.IsOnline, node.Username, node.Password, node.ID)
	if err != nil {
		return fmt.Errorf("failed to update node: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("node not found")
	}
	return nil
}

// Delete removes a node by ID.
func (s *NodeStore) Delete(id int) error {
	query := `DELETE FROM nodes WHERE id = $1`
	result, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete node: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("node not found")
	}
	return nil
}
