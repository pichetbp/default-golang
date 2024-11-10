package usecases

import "default-repo/libs/httpEntry/models"

type HttpEntryInterface interface {
	Post(args models.HTTPPostModel)
	Get(args models.HTTPGetModel)
	Put(args models.HTTPPutModel)
}
