package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewPost(t *testing.T) {
	post, err := NewPost("Testes de Unidade Go", "Para fazer um teste de unidade Go, crie um arquivo nomeEntidade_test.go")

	assert.Nil(t, err)
	assert.Equal(t, "Testes de Unidade Go", post.Title)
	assert.Equal(t, "Para fazer um teste de unidade Go, crie um arquivo nomeEntidade_test.go", post.Body)
}

func TestCreateNewPostWithInvalidArgs(t *testing.T) {
	post, err := NewPost("", "Para fazer um teste de unidade Go, crie um arquivo nomeEntidade_test.go")
	post2, err2 := NewPost("Testes de Unidade Go", "")

	assert.Nil(t, post)
	assert.Nil(t, post2)

	assert.EqualError(t, err, "title e um campo obrigatorio")
	assert.EqualError(t, err2, "body e um campo obrigatorio")
}

func TestUpdatePost(t *testing.T) {
	post, err := NewPost("1", "2")

	post.UpdatePost("Testes de Unidade Go", "Para fazer um teste de unidade Go, crie um arquivo nomeEntidade_test.go")

	assert.Nil(t, err)
	assert.Equal(t, "Testes de Unidade Go", post.Title)
	assert.Equal(t, "Para fazer um teste de unidade Go, crie um arquivo nomeEntidade_test.go", post.Body)
}

func TestUpdatePostWithInvalidArgs(t *testing.T) {
	post, _ := NewPost("1", "1")
	post2, _ := NewPost("1", "1")

	err := post.UpdatePost("", "2")
	err2 := post2.UpdatePost("2", "")

	assert.EqualError(t, err, "title e um campo obrigatorio")
	assert.EqualError(t, err2, "body e um campo obrigatorio")
}
