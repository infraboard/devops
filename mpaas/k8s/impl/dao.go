package impl

import (
	"context"

	"github.com/infraboard/devops/mpaas/k8s"
	"github.com/infraboard/mcube/v2/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *K8sServiceImpl) save(ctx context.Context, ins *k8s.Cluster) error {
	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted cluster(%s) document error, %s",
			ins.Name, err)
	}
	return nil
}

func (s *K8sServiceImpl) get(ctx context.Context, id string) (*k8s.Cluster, error) {
	filter := bson.M{"_id": id}

	ins := &k8s.Cluster{}
	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("cluster %s not found", id)
		}

		return nil, exception.NewInternalServerError("find cluster %s error, %s", id, err)
	}

	ins.SetIsEncrypted(true)
	return ins, nil
}
