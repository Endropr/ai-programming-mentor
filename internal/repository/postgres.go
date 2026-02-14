package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/Endropr/ai-programming-mentor/internal/domain"
)

// PostgresRepo будет отвечать за связь Go с твоей базой ai_mentor_db
type PostgresRepo struct {
	conn *pgx.Conn
}

func NewPostgresRepo(conn *pgx.Conn) *PostgresRepo {
	return &PostgresRepo{conn: conn}
}

// SaveMessage автоматически выполнит INSERT, который ты до этого делал вручную
func (r *PostgresRepo) SaveMessage(ctx context.Context, msg domain.Message) error {
	query := `INSERT INTO messages (user_id, role, content, selected_language) VALUES ($1, $2, $3, $4)`
	
	// Передаем значение из структуры domain.Message в запрос
	_, err := r.conn.Exec(ctx, query, msg.UserID, msg.Role, msg.Content, msg.SelectedLanguage)
	return err
}