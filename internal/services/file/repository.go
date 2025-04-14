package file

import (
	"context"
	"io"

	"github.com/CP-RektMart/schat-g28-backend/internal/model"
	"github.com/CP-RektMart/schat-g28-backend/pkg/apperror"
	"github.com/CP-RektMart/schat-g28-backend/pkg/storage"
	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	object *storage.Client
}

func NewRepository(db *gorm.DB, object *storage.Client) *Repository {
	return &Repository{
		db:     db,
		object: object,
	}
}

func (r *Repository) Create(
	ctx context.Context, path string, contentType string, data io.Reader, ownerID uint,
	fileFn func(URL string, path *string, ownerID uint) (model.File, error),
) (model.File, error) {
	URL, err := r.object.UploadFile(ctx, path, contentType, data, true)
	if err != nil {
		return model.File{}, errors.Wrap(err, "failed upload file")
	}

	f, err := fileFn(URL, &path, ownerID)
	if err != nil {
		return model.File{}, err
	}

	if err := r.db.Create(&f).Error; err != nil {
		return model.File{}, err
	}

	return f, nil
}

func (r *Repository) GetByID(id uint) (model.File, error) {
	var f model.File
	if err := r.db.First(&f, id).Error; err != nil {
		return model.File{}, apperror.NotFound("file not found", err)
	}
	return f, nil
}

func (r *Repository) Delete(ctx context.Context, id uint, permissionFn func(model.File) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		f, err := r.GetByID(id)
		if err != nil {
			return err
		}

		if err := permissionFn(f); err != nil {
			return errors.Wrap(err, "user don't have permission to delete")
		}

		if err := r.db.Delete(&model.File{}, id).Error; err != nil {
			return err
		}

		if f.Path == nil {
			return errors.New("there is no path provide")
		}

		if err := r.object.DeleteFile(ctx, *f.Path); err != nil {
			return errors.Wrap(err, "failed delete file")
		}

		return nil
	})
}
