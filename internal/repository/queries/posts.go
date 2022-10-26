package queries

import (
	"board/internal/handlers"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (q *Queries) AddPost(ctx context.Context, p *handlers.Post) (int, error) {
	if err := validPost(p); err != nil {
		return 0, err
	}
	var id int
	rows, err := q.conn.Query(ctx, "INSERT INTO board (title,price,descr,date)"+
		" VALUES ($1,$2,$3,CURRENT_TIMESTAMP(2)) "+
		"RETURNING post_id", p.Title, p.Price, p.Descr)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
	}

	var main bool
	for i, v := range p.Links {
		if i == 0 {
			main = true
		} else {
			main = false
		}
		_, err := q.conn.Exec(ctx, "INSERT INTO photo (post_id,main,link) "+
			"VALUES ($1,$2,$3)", id, main, v)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func validPost(p *handlers.Post) error {
	if p.Title != "" && p.Price != 0 && len(p.Links) != 0 {
		return nil
	}
	return errors.New("Invalid input")
}

func (q *Queries) GetPost(ctx context.Context, r *http.Request) (*handlers.Post, error) {
	idx, err := strconv.ParseInt(r.URL.Query()["id"][0], 10, 64)
	if err != nil {
		return nil, err
	}
	sql, full := getSql(r)
	rows, err := q.conn.Query(context.Background(), sql, idx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	p := handlers.NewPost()
	var link string
	for rows.Next() {
		if full {
			err = rows.Scan(&p.Title, &link, &p.Price, &p.Descr)
		} else {
			err = rows.Scan(&p.Title, &link, &p.Price)
		}
		if err != nil {
			return nil, err
		}
		p.Links = append(p.Links, link)
	}
	return p, nil
}

func getSql(r *http.Request) (string, bool) {
	var sql string
	if v, ok := r.URL.Query()["fields"]; ok {
		for _, fv := range v {
			switch fv {
			case "full":
				sql = fmt.Sprintf("SELECT title, link, price, descr " +
					"FROM board b JOIN photo p ON b.post_id = p.post_id " +
					"WHERE b.post_id = $1")
				return sql, true
			}
		}
		sql = "SELECT title, link, price FROM board b JOIN photo p " +
			"ON b.post_id = p.post_id WHERE b.post_id = $1 AND p.main = TRUE"
	} else {
		sql = "SELECT title, link, price FROM board b JOIN photo p " +
			"ON b.post_id = p.post_id WHERE b.post_id = $1 AND p.main = TRUE"
	}
	return sql, false
}

func (q *Queries) GetAllPosts(ctx context.Context, r *http.Request) ([]handlers.Post, error) {
	page, err := getPage(r)
	if err != nil {
		return nil, err
	}
	sortby := getSortby(r)
	rows, err := q.conn.Query(ctx, "SELECT title, link, price FROM board b JOIN photo p "+
		"ON b.post_id = p.post_id WHERE p.main = TRUE ORDER BY $1 "+
		"LIMIT 10 OFFSET $2*10", sortby, page-1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []handlers.Post{}
	for rows.Next() {
		p := handlers.NewPost()
		var link string
		err := rows.Scan(&p.Title, &link, &p.Price)
		if err != nil {
			return nil, err
		}
		p.Links = append(p.Links, link)
		posts = append(posts, *p)
	}
	return posts, nil
}

func getPage(r *http.Request) (int, error) {
	if v, ok := r.URL.Query()["page"]; ok {
		page, err := strconv.ParseInt(v[0], 10, 64)
		if err != nil {
			return 0, err
		}
		return int(page), nil
	}
	return 1, nil
}

func getSortby(r *http.Request) string {
	if v, ok := r.URL.Query()["sortby"]; ok {
		switch v[0] {
		case "plow":
			return "price ASC"
		case "phigh":
			return "price DESC"
		case "dlow":
			return "date ASC"
		case "dhigh":
			return "date DESC"
		default:
			return "post_id ASC"
		}
	}
	return "post_id ASC"
}
