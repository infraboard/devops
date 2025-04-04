package impl

import (
	"context"

	"github.com/infraboard/devops/cmdb/resource"
	"github.com/infraboard/mcube/v2/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Save 更新与创建同时
// in
func (i *impl) Save(ctx context.Context, in *resource.Resource) (*resource.Resource, error) {
	_, err := i.col.UpdateOne(ctx, bson.M{"_id": in.Id}, bson.M{"$set": in}, options.Update().SetUpsert(true))
	if err != nil {
		return nil, err
	}
	return in, nil
}

// 资源搜索
func (i *impl) Search(ctx context.Context, in *resource.SearchRequest) (*types.Set[*resource.Resource], error) {
	set := resource.NewResourceSet()

	filter := bson.M{}
	// 1. 关键字搜索 { name: { $regex: 'acme.*corp', $options: "si" } }
	if in.Keywords != "" {
		filter["name"] = bson.M{"$regex": in.Keywords, "$options": "im"}
	}

	// 2. 类型过滤
	if in.Type != nil {
		filter["resource_type"] = *in.Type
	}
	// 3.标签过滤
	// "tags": {
	//         "app": "app01"
	//       },
	for k, v := range in.Tags {
		filter["tags."+k] = v
	}

	opt := options.Find()
	opt.SetLimit(int64(in.PageSize))
	opt.SetSkip(in.ComputeOffset())
	cursor, err := i.col.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		e := resource.NewResource()
		if err := cursor.Decode(e); err != nil {
			return nil, err
		}
		set.Add(e)
	}

	return set, nil
}

// 删除
// {"_id": {"$in": []}}
func (i *impl) DeleteResource(ctx context.Context, in *resource.DeleteResourceRequest) error {
	_, err := i.col.DeleteOne(ctx, bson.M{"_id": bson.M{"$in": in.ResourceIds}})
	if err != nil {
		return err
	}
	return nil
}
