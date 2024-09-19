package repositories

import (
	"database/sql"
	"devbook_api/src/models"
)

type Post struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *Post {
	return &Post{db}
}

func (repository Post) Create(post models.Post) (models.Post, error) {
	statement, err := repository.db.Prepare("insert into post (title, description, author_id) values (?, ?, ?)")
	if err != nil {
		return models.Post{}, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Description, post.AuthorID)
	if err != nil {
		return models.Post{}, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return models.Post{}, err
	}

	post.ID = uint64(lastInsertedID)
	return post, nil
}

func (repository Post) FindPost(postId uint64) (models.Post, error) {
	statement, err := repository.db.Prepare(
		`select p.id,
		 	p.title,
		  	p.description,
		   	p.author_id,
		    u.nick,
			COALESCE(l.like_count, 0) as likes,
			p.created_at
		   	from post p 
		   	inner join user u on p.author_id = u.id
			left join (select post_id, count(1) as like_count from devbook.like group by post_id) l on l.post_id = p.id
	   		where p.id = ?`)
	if err != nil {
		return models.Post{}, err
	}
	defer statement.Close()

	result := statement.QueryRow(postId)

	var post models.Post
	if err = result.Scan(&post.ID, &post.Title, &post.Description, &post.AuthorID, &post.AuthorNick, &post.Likes, &post.CreatedAt); err != nil {
		return models.Post{}, err
	}

	return post, nil
}

func (repository Post) FindPosts(userId uint64) ([]models.Post, error) {
	statement, err := repository.db.Prepare(
		`select p.id, p.title, p.description, p.author_id, COALESCE(l.like_count, 0) as likes, p.created_at, p.updated_at, u2.nick,
		coalesce(liked_by_me.is_liked, 0) as i_liked
		from post p
		join user u2 on u2.id = p.author_id
		left join follower f on p.author_id = f.user_id
		left join (select post_id, count(1) as like_count from devbook.like group by post_id) l on l.post_id = p.id
		left join (
            select post_id, 1 as is_liked
            from devbook.like
            where user_id = ?
        ) liked_by_me on p.id = liked_by_me.post_id
		where f.follower_id = ? or p.author_id = ?
		group by p.id, p.title, p.description, p.author_id, p.created_at, p.updated_at, u2.nick;`)

	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query(userId, userId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.UpdateAt,
			&post.AuthorNick,
			&post.ILiked,
		); err != nil {
			return []models.Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Post) UpdatePost(postID int64, post models.Post) error {
	statement, err := repository.db.Prepare("update post set title = ?, description = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(post.Title, post.Description, postID); err != nil {
		return err
	}

	return nil
}

func (repository Post) DeletePost(postID int64) error {
	statement, err := repository.db.Prepare("delete from post where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (repository Post) FindPostsByUser(userId, myId uint64) ([]models.Post, error) {
	statement, err := repository.db.Prepare(
		`select p.id, p.title, p.description, p.author_id, COALESCE(l.like_count, 0) as likes, p.created_at, p.updated_at, u2.nick,
		coalesce(liked_by_me.is_liked, 0) as i_liked
		from post p
		join user u2 on u2.id = p.author_id
		left join (select post_id, count(1) as like_count from devbook.like group by post_id) l on l.post_id = p.id
		left join (
            select post_id, 1 as is_liked
            from devbook.like
            where user_id = ?
        ) liked_by_me on p.id = liked_by_me.post_id
		where p.author_id = ?`)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query(myId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Description,
			&post.AuthorID,
			&post.Likes,
			&post.CreatedAt,
			&post.UpdateAt,
			&post.AuthorNick,
			&post.ILiked,
		); err != nil {
			return []models.Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (repository Post) LikePost(postID, userID uint64) error {
	statement, err := repository.db.Prepare("insert ignore into devbook.like (post_id, user_id) values (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID, userID); err != nil {
		return err
	}

	return nil
}

func (repository Post) UnlikePost(postID, userID uint64) error {
	statement, err := repository.db.Prepare("delete from devbook.like where post_id = ? and user_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID, userID); err != nil {
		return err
	}

	return nil
}
